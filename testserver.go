package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/tmc/grpc-websocket-proxy/wsproxy"
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
	log.Println(in)
	res:=pb.TestResponse{Res: "hello  "+ time.Now().String()}
	return &res,nil
}

func (this *Test)GetStream(req *pb.StreamReqData, res pb.Test_GetStreamServer) error {
	for{
		err:=res.Send(&pb.StreamResData{Data: "getstream  "+ time.Now().String()})
		time.Sleep(time.Second)
		if err!= nil{
			break
		}
	}
	log.Println("get stream退出")
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
			data,err:=allStr.Recv()
			if err != nil{
				log.Println("all stearm rec err:",err)
				break
			}
			log.Println("all stream rec:",data)
		}
		wg.Done()
	}()

	go func() {
		for{
			err:=allStr.Send(&pb.StreamResData{Data:"ssss"})
			if err != nil{
				log.Println("allstream send error:",err)
				break
			}
			log.Println("allsteam sending")
			time.Sleep(time.Second)
		}
		wg.Done()
	}()
	wg.Wait()
	return nil
}

func main() {
	listen,err:= net.Listen("tcp",":8000")
	if err != nil{
		log.Fatal("tcp listen err:",err)
	}
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

	err=pb.RegisterTestHandlerFromEndpoint(ctx,mux,"localhost:8000",opt)
	if err!= nil{
		log.Fatal("RegisterTestHandlerFromEndpoint err:",err)
	}

	wsMux := wsproxy.WebsocketProxy(mux)

	httpServer:=&http.Server{
		Addr: ":9000",
		Handler: wsMux,
	}
	httpServer.ListenAndServe()
}
