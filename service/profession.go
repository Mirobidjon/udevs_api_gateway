package service

import (
	"context"
	"fmt"

	"bitbucket.org/udevs/example_api_gateway/genproto/admin_api_gateway"
	"bitbucket.org/udevs/example_api_gateway/pkg/logger"
	"bitbucket.org/udevs/example_api_gateway/service/grpc_client"
)

type professionService struct {
	client grpc_client.ServiceManager
	logger logger.Logger
	admin_api_gateway.UnimplementedProfessionApiGatewayServer
}

func NewProfessionService(client grpc_client.ServiceManager, log logger.Logger) *professionService {
	return &professionService{
		client: client,
		logger: log,
	}
}

func (h *professionService) Create(ctx context.Context, req *admin_api_gateway.GatewayCreateProfessionReq) (*admin_api_gateway.GatewayProfessionId, error) {
	fmt.Printf("\n\nCreate Profession request!!  %+v\n\n", req)

	return &admin_api_gateway.GatewayProfessionId{
		ProfessionId: "Successfully created",
	}, nil
}

func (h *professionService) GetAll(ctx context.Context, req *admin_api_gateway.GatewayGetAllProfessionRequest) (*admin_api_gateway.GatewayGetAllProfessionResponse, error) {
	fmt.Printf("\n\n Get All req %+v\n\n", req)

	return &admin_api_gateway.GatewayGetAllProfessionResponse{
		Professions: []*admin_api_gateway.GatewayProfession{},
		Count:       15,
	}, nil
}

func (h *professionService) Get(ctx context.Context, req *admin_api_gateway.GatewayProfessionId) (*admin_api_gateway.GatewayProfession, error) {
	fmt.Printf("\n\n Get All req %+v\n\n", req)

	return &admin_api_gateway.GatewayProfession{
		Id:   "1",
		Name: "Worker",
	}, nil
}
