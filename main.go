package main

import (
	"fmt"
	"time"

	"github.com/ewanvalentine/mgo-proto-test/drivers"
	proto "github.com/ewanvalentine/mgo-proto-test/proto/greeter"

	grpc "github.com/micro/go-grpc"
	micro "github.com/micro/go-micro"
)

func main() {
	session := drivers.Init()

	defer session.Close()

	service := grpc.NewService(
		micro.Name("go.micro.srv.greeter"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// Init will parse the command line flags.
	service.Init()

	greeter := &Greeter{session}

	// Register handler
	proto.RegisterGreeterHandler(service.Server(), greeter)

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
