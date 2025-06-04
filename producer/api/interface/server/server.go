package server

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "wordCountServer/api/interface/server/proto"
)

// server 
type server struct {
	pb.UnimplementedWordCountServer
}

// WordCount 
//	@param ctx 
//	@param req 
//	@return *pb.WordCountResponse 
//	@return error 
func (s *server) WordCount(ctx context.Context, req *pb.WordCountRequest) (*pb.WordCountResponse, error) {;
	// ask Spark for word count result 
	results := s.dialSpark(req.RequestMessage) 

	return &pb.WordCountResponse{
		Ct: results, 
	}, nil 
}

// dialSpark 
//	@param text 
//	@return []*pb.WordCountResponse_WordCountResult 
func (s *server) dialSpark(text string) []*pb.WordCountResponse_WordCountResult {
	conn, err := grpc.NewClient("127.0.0.1:9091", 
		grpc.WithTransportCredentials(insecure.NewCredentials())) 
	if err != nil {
		log.Fatal("unable to dial spark ", err) 
	} 
	defer conn.Close() 

	client := pb.NewWordCountClient(conn) 
	resp, _ := client.WordCount(context.Background(), &pb.WordCountRequest{
		RequestMessage: text,
	})
	
	return resp.GetCt()
}

// Server 
func Server() {
	listen, err := net.Listen("tcp", ":9090") 
	if err != nil {
		log.Fatal("unable to listen: ", err)
	}
	grpcServer := grpc.NewServer() 
	pb.RegisterWordCountServer(grpcServer, &server{}) 

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatal("unable to run grpc") 
	}
}

