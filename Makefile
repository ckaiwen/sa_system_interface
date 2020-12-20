APP=modbustest

protoc:
	protoc --go_out=plugins=grpc:.  ./proto/test/test.proto

http:
	protoc --grpc-gateway_out=logtostderr=true,grpc_api_configuration=./proto/test/test.yaml:. ./proto/test/test.proto
run:
	go run cmd/cmd.go
test:
	go run cmd/test.go
clean:
	rm -f ${APP}


