
all_protoc: cq_decode_protoc

cq_decode_protoc:
	protoc --go_out=plugins=grpc:./gosdk/cq_decode  ./proto/cq_decode/cq_decode.proto ;\
    protoc --grpc-gateway_out=logtostderr=true,grpc_api_configuration=./proto/cq_decode/cq_decode.http.yaml:./gosdk/cq_decode  ./proto/cq_decode/cq_decode.proto




