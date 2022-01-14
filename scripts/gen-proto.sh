#!/bin/bash
CURRENT_DIR=$(pwd)
for x in $(find ${CURRENT_DIR}/protos/* -type d); do
    if [ $x == ${CURRENT_DIR}/protos/*_gateway ]
    then
        sudo protoc -I=${x} -I=${CURRENT_DIR}/protos -I /usr/local/include \
        -I=${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
        -I=${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway \
         --go_out=${CURRENT_DIR} --go-grpc_out=${CURRENT_DIR} \
         --go_opt=paths=import \
         --swagger_out=allow_merge=true,merge_file_name=api:./third_party/swagger/ \
		 --grpc-gateway_opt=generate_unbound_methods=true \
         --grpc-gateway_out=logtostderr=true,repeated_path_param_separator=ssv:. \
         --openapiv2_out=logtostderr=true,repeated_path_param_separator=ssv:./third_party/swagger/ \
         --grpc-gateway_out=logtostderr=true:${CURRENT_DIR} ${x}/*.proto
    else
        sudo protoc -I=${x} -I=${CURRENT_DIR}/protos -I /usr/local/include \
         --go_out=${CURRENT_DIR} \
         --go-grpc_out=${CURRENT_DIR} ${x}/*.proto
    fi
done
