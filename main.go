package main

import (
	"net/http"

	"github.com/aungshanbo/go-mongo-rest/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {

	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:9000", r)
	// if err := http.ListenAndServe(":8080", r); err != nil {
	// 	log.Fatalf("Server failed to start: %v", err)
	// }

}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://127.0.0.1:27017")
	if err != nil {
		panic(err)
	}

	return s
	// info := &mgo.DialInfo{
	// 	Addrs:    []string{"127.0.0.1:27017"},
	// 	Username: "users",
	// 	Password: "users",
	// 	Database: "mongo-golang",
	// }

	// session, err := mgo.DialWithInfo(info)
	// if err != nil {
	// 	log.Fatalf("MongoDB Connection failed: %v", err)
	// }
	// return session
}
