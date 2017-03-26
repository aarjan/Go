package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type database struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	By          string `json:"by"`
	Likes       int    `json:"likes"`
}

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("test").C("myCollection")
	// var mapArray []map[string]interface{}

	i := []database{}
	err = c.Find(bson.M{}).All(&i)

	if err != nil {
		panic(err)
	}
	// fmt.Println(i)

	for _, m := range i {
		fmt.Println("title : ", m.Title)
		fmt.Println("description : ", m.Description)
		fmt.Println("by : ", m.By)
		fmt.Println("likes : ", m.Likes)
		fmt.Println("--------------")

	}

	// var m1 = map[string]interface{}{
	// 	"title":       "Cassandra",
	// 	"description": "apache relational db",
	// 	"by":          "anuj",
	// 	"likes":       189,
	// }

	// err = c.Insert(i[1])
	// if err != nil {
	// 	panic(err)
	// }

}
