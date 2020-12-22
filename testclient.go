package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "sa_system/sa_system_interface/proto/test"
	"time"
)

func main() {
	conn,err:=grpc.Dial("localhost:8000",grpc.WithInsecure())
	if err!=nil{
		log.Fatal("conn err:",err)
	}

	defer conn.Close()

	c:=pb.NewTestClient(conn)
	//双方普通请求
	sayHelloRes,err:=c.SayHello(context.Background(),&pb.TestRequest{Req: "123"})
	log.Println(sayHelloRes)

	//服务器流
	//getStreamRes,_:= c.GetStream(context.Background(),&pb.StreamReqData{Data: "123"})
	//for{
	//	singleRec,err:=getStreamRes.Recv()
	//	if err!=nil{
	//		log.Println(err)
	//		break
	//	}
	//	log.Println(singleRec)
	//}
	//
	////客户端流
	//putRes,_:=c.PutStream(context.Background())
	//i:=1
	//for{
	//	i++
	//	putRes.Send(&pb.StreamReqData{Data: "ss"})
	//	time.Sleep(time.Second)
	//	if i>2{
	//		break
	//	}
	//}

	//双向流
	allStr,_:=c.AllStream(context.Background())
	go func() {
		for{
			data,_:=allStr.Recv()
			log.Println(data)
		}
	}()

	go func() {
		for  {
			allStr.Send(&pb.StreamReqData{Data: "ssss"})
			time.Sleep(time.Second)
		}
	}()

	select {

	}
}
