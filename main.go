package main

import (
	"grpc-test/client"
	"grpc-test/server"
	"log"
	"time"
)

func main() {
	log.Println("running server")
	go server.RunServer()
	time.Sleep(time.Second * 1)
	log.Println("sending message to server")
	client.SendGrpcClientRequest()
}
