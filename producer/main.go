package main

import (
	"log"
	"wordCountServer/api/interface/client"
	"wordCountServer/global/initialize"
)

// main
func main() {
	router := initialize.GlobalInit() 

	go func() {
		err := client.RunClient()
		if err != nil {
			log.Fatal("unable to run grpc: ", err) 
		}
	}()

	log.Println("running grpc...")
	go func() {
		err := router.Run(":8080") 
		if err != nil {
			log.Fatal("unable to run Gin: ", err) 
		}
	}()
	
	select {}
}   