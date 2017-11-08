package main

import (
	proto "github.com/EwanValentine/mgo-proto-test/proto/greeter"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

func (repo *Repository) GetAll() ([]*proto.Greeting, error) {
	var greetings []*proto.Greeting
	err := repo.collection().Find(nil).All(&greetings)
	if err != nil {
		return nil, err
	}
	return greetings, nil
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
