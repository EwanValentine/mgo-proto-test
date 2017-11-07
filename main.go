package main

import (
	"fmt"
	"time"

	proto "github.com/ewanvalentine/mgo-proto-test/proto/greeter"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	grpc "github.com/micro/go-grpc"
	micro "github.com/micro/go-micro"
	"golang.org/x/net/context"
)

type Repository struct {
	session *mgo.Session
}

func (repo *Repository) Get(id string) (*proto.Greeting, error) {
	var greeting *proto.Greeting
	err := repo.collection().FindId(bson.ObjectIdHex(id)).One(&greeting)
	if err != nil {
		return nil, err
	}
	return greeting, nil
}

func (repo *Repository) Post(greeting *proto.Greeting) (*proto.Greeting, error) {
	err := repo.collection().Insert(greeting)
	if err != nil {
		return greeting, err
	}
	return greeting, nil
}

func (repo *Repository) collection() *mgo.Collection {
	return repo.session.DB("greeter").C("greetings")
}

func (repo *Repository) Close() {
	repo.session.Close()
}

type Greeter struct {
	session *mgo.Session
}

func (g *Greeter) repo() *Repository {
	return &Repository{g.session.Copy()}
}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, res *proto.Response) error {
	repo := g.repo()
	defer repo.Close()
	greeting, err := repo.Get(req.Id)
	if err != nil {
		return err
	}
	res.Body = greeting
	return nil
}

func (g *Greeter) PostGreeting(ctx context.Context, req *proto.Request, res *proto.Response) error {
	repo := g.repo()
	defer repo.Close()
	greeting, err := repo.Post(req.Body)
	if err != nil {
		return err
	}
	res.Body = greeting
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

	greeter := &Greeter{session}

	// Register handler
	proto.RegisterGreeterHandler(service.Server(), greeter)

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
