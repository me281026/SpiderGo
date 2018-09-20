package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

func main() {

}

//获取Mongo连接
func MongoConnection() *mgo.Collection {
	//host,port,dbname
	host := "127.0.0.1"
	port := "27017"
	dbname := "test"
	//session
	session, err := mgo.Dial(host + ":" + port)
	if err != nil {
		fmt.Println(err.Error())
	}
	//defer
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	//获取DB连接
	collection := session.DB(dbname).C("User")
	return collection
}
