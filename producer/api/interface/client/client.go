package client

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "wordCountServer/api/interface/server/proto"
)

func Client() {
	conn, err := grpc.Dial("127.0.0.1:9090", 
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("unable to dial") 
	} 
	defer conn.Close() 

	client := pb.NewWordCountClient(conn) 
	resp, _ := client.WordCount(context.Background(), &pb.WordCountRequest{RequestMessage: ""}) 
	fmt.Println(resp.GetCt()) 
}