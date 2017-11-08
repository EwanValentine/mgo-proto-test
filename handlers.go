package main

import (
	"context"
	"encoding/hex"

	proto "github.com/ewanvalentine/mgo-proto-test/proto/greeter"

	mgo "gopkg.in/mgo.v2"
)

type Greeter struct {
	session *mgo.Session
}

func (g *Greeter) repo() *Repository {
	return &Repository{g.session.Copy()}
}

// Hex - convert ID to string
// This is an annoying thing to have to do manually actually.
// But an ID shows up as garbage when rendered to JSON without this.
func Hex(id string) string {
	return hex.EncodeToString([]byte(id))
}

// convertIDs - converts all Id's from hex to string.
func convertIDs(objects []*proto.Greeting) []*proto.Greeting {
	var updated []*proto.Greeting
	for _, v := range objects {
		v.Id = Hex(v.Id)
		updated = append(updated, v)
	}
	return updated
}

func (g *Greeter) GetGreeting(ctx context.Context, req *proto.HelloRequest, res *proto.Response) error {
	repo := g.repo()
	defer repo.Close()
	greeting, err := repo.Get(req.Id)
	if err != nil {
		return err
	}
	res.Body = greeting
	res.Body.Id = Hex(greeting.Id)
	return nil
}

func (g *Greeter) GetGreetings(ctc context.Context, req *proto.Request, res *proto.Response) error {
	repo := g.repo()
	defer repo.Close()
	greetings, err := repo.GetAll()
	if err != nil {
		return err
	}
	res.Collection = convertIDs(greetings)
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
