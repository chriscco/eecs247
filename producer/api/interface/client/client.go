package client

import (
	"context"
	"io"
	"log"
	"sync"
	pb "wordCountServer/api/interface/server/proto"
	"wordCountServer/global"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	SenderChannel chan string
)

// streamClient 
var streamClient pb.WordCountClient 

// wordCount 
//	@return error 
func wordCount() error {
	var senderWg sync.WaitGroup 
	var receivWg sync.WaitGroup
	senderWg.Add(1)
	receivWg.Add(1)

	stream, err := streamClient.WordCount(context.Background())
	defer stream.CloseSend() 

	if err != nil {
		log.Println("unable to create stream")
		return err 
	}
	go func(stream pb.WordCount_WordCountClient) {
		defer senderWg.Done() 
		sender(stream)
	}(stream)
	
	go func (stream pb.WordCount_WordCountClient)  {
		defer receivWg.Done() 
		receiver(stream)
	}(stream)

	senderWg.Wait() 
	receivWg.Wait() 

	return nil 
}

// sender 
//	@param stream 
//	@return error 
func sender(stream pb.WordCount_WordCountClient) error {
	for {
		select {
		case text := <- SenderChannel:
			err := stream.Send(&pb.WordCountRequest {
				Message: text,
			})
			if err != nil {
				log.Fatal("unable to send grpc ", err)
				return err 
			}
			log.Println("successfully send one message")
		}
	}
}

// receiver 
//	@param stream 
//	@return error 
func receiver(stream pb.WordCount_WordCountClient) error {
	for {
		req, err := stream.Recv() 
		if err == io.EOF {
			log.Println("EOF in receiver")
			return nil 
		}
		if err != nil {
			log.Fatalf("error in receiver: %s, %s\n", err, req) 
			return err
		} 
		
		key := req.GetKey() 
		for _, cnt := range req.GetCt() {
			global.Redis.HSet(key, cnt.GetWord(), cnt.GetCount())
		}
	}
}

// RunClient 
//	@return error 
func RunClient() error {
	SenderChannel = make(chan string, 10)

	conn, err := grpc.NewClient(":9091", 
		grpc.WithTransportCredentials(insecure.NewCredentials().Clone()), 
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(1024*1024*20)), 
		grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(1024*1024*20)))
		
	if err != nil {
		log.Println("unable to dial grpc")
		return err 
	}
	defer conn.Close() 
	streamClient = pb.NewWordCountClient(conn)
	wordCount()
	return nil 
}