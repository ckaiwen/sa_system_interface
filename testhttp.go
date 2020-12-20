package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	pb "sa_system/sa_system_interface/proto/test"
)

type Test struct {

}

func (this *Test)SayHello(ctx context.Context, in *pb.TestRequest) (*pb.TestResponse, error) {
	res:=pb.TestResponse{}
	res.Res="hello"
	return &res,nil
}


func main() {
	grcServer:=grpc.NewServer()
	srv:=Test{}
	pb.RegisterTestServer(grcServer,&srv)
	listen,_ := net.Listen("tcp",":8080")
    go func() {
		err:= grcServer.Serve(listen)
		if err!=nil{
			log.Fatal("grpc server err:",err)
		}
	}()

	ctx:= context.Background()
	ctx,cancel:= context.WithCancel(ctx)
	defer cancel()

	mux:= runtime.NewServeMux()
	opt:= []grpc.DialOption{grpc.WithInsecure()}

	err:=pb.RegisterTestHandlerFromEndpoint(ctx,mux,"localhost:8080",opt)
	if err!= nil{
		log.Fatal("RegisterTestHandlerFromEndpoint err:",err)
	}

	httpServer:=&http.Server{
		Addr: ":8081",
		Handler: mux,
	}
	httpServer.ListenAndServe()
}
