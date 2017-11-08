package drivers

import (
	"os"

	mgo "gopkg.in/mgo.v2"
)

func Init() *mgo.Session {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}

	session, err := mgo.Dial(host)
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	return session
}
