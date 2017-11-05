package main

import (
	"fmt"
	"log"
	"time"

	proto "github.com/ewanvalentine/mgo-proto-test/proto/greeter"
	mgo "gopkg.in/mgo.v2"

	grpc "github.com/micro/go-grpc"
	micro "github.com/micro/go-micro"
	"golang.org/x/net/context"
)

type Repository struct {
	session *mgo.Session
}

func (repo *Repository) Get() *proto.Greeting {
	var greeting *proto.Greeting

	return greeting
}

func (repo *Repository) Close() {
	repo.session.Close()
}

type Greeter struct {
}

func (g *Greeter) repo() *Repository {
	return &Repository{g.session.Copy()}
}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	log.Println(req.Name)
	repo := g.repo()
	defer repo.Close()
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	service := grpc.NewService(
		micro.Name("go.micro.srv.greeter"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// Init will parse the command line flags.
	service.Init()

	// Register handler
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
