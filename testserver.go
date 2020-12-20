package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/xsephiroth/grpc-websocket-proxy/wsproxy"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	pb "sa_system/sa_system_interface/proto/test"
	"sync"
	"time"
)

type Test struct {

}

func (this *Test)SayHello(ctx context.Context, in *pb.TestRequest) (*pb.TestResponse, error) {
	res:=pb.TestResponse{}
	res.Res="hello"
	return &res,nil
}

func (this *Test)GetStream(req *pb.StreamReqData, res pb.Test_GetStreamServer) error {
	i:=0
	for{
		i++
		res.Send(&pb.StreamResData{Data: fmt.Sprintf("%v",time.Now().Unix())})
		time.Sleep(time.Second)
		if i>1{
			break
		}
	}
	return nil
}

func (this *Test) PutStream(clistr pb.Test_PutStreamServer) error{
	for{
		if tem,err:=clistr.Recv();err==nil{
			log.Println(tem)
		}else {
			log.Println("break,err:",err)
			break
		}
	}
	return  nil
}

func (this *Test)AllStream(allStr pb.Test_AllStreamServer) error{
	wg:= sync.WaitGroup{}
	wg.Add(2)

	go func() {
		for{
			data,_:=allStr.Recv()
			log.Println(data)
		}
		wg.Done()
	}()

	go func() {
		for{
			allStr.Send(&pb.StreamResData{Data:"ssss"})
			time.Sleep(time.Second)
		}
		wg.Done()
	}()
	wg.Wait()
	return nil
}

func main() {
	listen,_ := net.Listen("tcp",":8080")
	grcServer:=grpc.NewServer()
	srv:=Test{}
	pb.RegisterTestServer(grcServer,&srv)

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

	wsMux := wsproxy.WebsocketProxy(mux)

	httpServer:=&http.Server{
		Addr: ":8081",
		Handler: wsMux,
	}
	httpServer.ListenAndServe()
}
