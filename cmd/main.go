package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"bitbucket.org/udevs/example_api_gateway/config"
	"bitbucket.org/udevs/example_api_gateway/genproto/admin_api_gateway"
	"bitbucket.org/udevs/example_api_gateway/pkg/logger"
	"bitbucket.org/udevs/example_api_gateway/pkg/middleware"
	"bitbucket.org/udevs/example_api_gateway/service"
	"bitbucket.org/udevs/example_api_gateway/service/grpc_client"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.Environment, "example_api_gateway")
	defer logger.Cleanup(log)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		err := RunHttpServer(ctx, log, cfg)
		if err != nil {
			panic(err)
		}
	}()

	if err := RunGrpcServer(ctx, log, cfg); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func RunHttpServer(ctx context.Context, log logger.Logger, cfg config.Config) error {

	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames: true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		}),
	)
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// oa := getOpenAPIHandler()
	err := admin_api_gateway.RegisterProfessionApiGatewayHandlerFromEndpoint(ctx, mux, "localhost"+cfg.RPCPort, opts)
	if err != nil {
		log.Fatal("failed to start HTTP gateway", zap.String("reason", err.Error()))
		return err
	}

	httpMux := http.NewServeMux()
	httpMux.Handle("/", mux)
	prefix := "/swagger/"
	fileServer := http.FileServer(http.Dir(string("third_party/")))
	httpMux.Handle(prefix, http.StripPrefix(prefix, fileServer))

	server := &http.Server{
		Addr: cfg.HttpPort,
		// add handler with middleware
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, "/api") {
				// custom function
				mux.ServeHTTP(w, r)
				return
			}
			httpMux.ServeHTTP(w, r)
			// oa.ServeHTTP(w, r)
		}),
	}

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for range c {
			// sig is a ^C, handle it
		}

		_, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		_ = server.Shutdown(ctx)
	}()

	log.Warn("starting HTTP/REST gateway...")
	return server.ListenAndServe()
}

func RunGrpcServer(ctx context.Context, log logger.Logger, cfg config.Config) error {
	client, err := grpc_client.NewGrpcClients(&cfg)
	if err != nil {
		log.Error("error while connecting to clients", logger.Error(err))
		return err
	}

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Error("error while listening: %v", logger.Error(err))
		return err
	}
	// gRPC server statup options
	opts := []grpc.ServerOption{}

	// add middleware
	opts = middleware.AddLogging(logger.GetZapLogger(log), opts)

	professionService := service.NewProfessionService(client, log)

	s := grpc.NewServer(opts...)

	admin_api_gateway.RegisterProfessionApiGatewayServer(s, professionService)

	reflection.Register(s)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Warn("shutting down gRPC server...")

			s.GracefulStop()

			<-ctx.Done()
		}
	}()
	// start gRPC server
	log.Info("starting gRPC server...")
	return s.Serve(lis)
}
