#gRPC:
 gRPC works by using protocol buffers tThe services can be written is any supported languages, as long as they share the protocol, they cn setup communication channels between themselves.

##What you need to do:
Since we're using golang, we need to set up `protoc` and `protoc-gen-go`.
####Set up `protoc` : 
```shell script
$ curl -LO $PB_REL/download/v3.11.4/protoc-3.11.4-linux-x86_64.zip
4 sudo unzip protoc-3.11.4-linux-x86_64.zip -d /usr
```
####Set up `protoc-gen-go` : 

```shell script
$ go get -u -v github.com/golang/protobuf/protoc-gen-go
```

####Write a proto file:

You need to have a protocol file (.proto) that will be shared between client and server. We will generate language specific (in our case golang) protocol functions using this file.
A proto file contains services and messages. Services contain the rpc definitions, the remote procedure calls you want to facilitate. The messages represent the data, that will be exchanged between server and client.

Here's an example of a protocol file:
```proto
syntax = "proto3";
package proto;
service HelloWorld {
  rpc SayHello (HelloRequest) returns (HelloResponse) {}
  rpc SayMyName (HelloRequest) returns (HelloResponse) {}
}

service RestaurantService {
  rpc OrderFood (OrderList) returns (Servings) {}
}

message HelloRequest {
  string name = 1;
}

message HelloResponse {
  string message = 1;
}


message OrderList {
  string foodItem1 = 1;
  string foodItem2 = 2;
}

message Servings {
  string foodItem1 = 1;
  string foodItem2 = 2;
}
```

note: package is the destination directory where you are going to generate the language specific protocol file.

####Generate language specific protocol file:
```shell script
$ protoc --go_out=plugins=grpc:<destinatin_directory> <source.proto>
```

Now you are ready to write gRPC servers and clients using the generated protocol file.

##Writing gRPC servers:
Writing grPC ervers are simple, but we must register defined services in our grpc server to serve messages to client. For that, we need to define a structure that implements functions (or procedures) corresponding to a specific service defined in the protocol.
For example, for this service defined in the protocol file:

```proto
service RestaurantService {
  rpc OrderFood (OrderList) returns (Servings) {}
}
```
we need to write a service structure like this:

```go
type FoodOrderService struct{}

func (f *FoodOrderService) OrderFood(context context.Context, in *proto.OrderList) (*proto.Servings, error) {
	return &proto.Servings{
		FoodItem1: "Your food item " + in.FoodItem1 + " available",
		FoodItem2: "Your food item " + in.FoodItem2 + " available",
	}, nil
}
```

Then we will use this structure when we register this service in our server declaration:
```go
func RunServer() {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}
	grpcServer := grpc.NewServer()

	proto.RegisterRestaurantServiceServer(grpcServer, &FoodOrderService{})

	reflection.Register(grpcServer)
	grpcServer.Serve(listen)
}

```

Now we have registered our service in gRPC server.

##Writing gRPC client:

First, we need to set up a connection on the appropriate server and port.

```go
conn, err := grpc.Dial(":9000", grpc.WithInsecure())
```

we will use this to connect to desired services, and make call to remote procedures associated with that service.

```go
	o := proto.NewRestaurantServiceClient(conn)

	response2, err := o.OrderFood(context.Background(), &proto.OrderList{FoodItem1: "Garlic Nun"})
```

Viola! We have successfully set up functional remote procedure calls.