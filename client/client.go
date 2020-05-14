package client

import (
	"context"
	"google.golang.org/grpc"
	"grpc-test/proto"
	"log"
)

func SendGrpcClientRequest() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewHelloWorldClient(conn)

	response, err := c.SayMyName(context.Background(), &proto.HelloRequest{Name: "Rezoan"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Message)
	o := proto.NewRestaurantServiceClient(conn)

	response2, err := o.OrderFood(context.Background(), &proto.OrderList{FoodItem1: "Garlic Nun"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: \n")
	log.Println(response2.FoodItem1)
	log.Println(response2.FoodItem2)

}
