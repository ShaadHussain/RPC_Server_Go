package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Item struct {
	Title string
	Body  string
}

func main() {

	var reply Item
	var db []Item

	client, err := rpc.DialHTTP("tcp", "localhost:4040")

	if err != nil {
		log.Fatal("Cnxn error: ", err)
	}

	a := Item{"First", "First item"}
	b := Item{"Second", "2nd item"}
	c := Item{"Third", "3rd item"}

	client.Call("API.AddItem", a, &reply)
	client.Call("API.AddItem", b, &reply)
	client.Call("API.AddItem", c, &reply)
	client.Call("API.GetDB", "", &db)

	fmt.Println("Database 1: ", db)

	client.Call("API.EditItem", Item{"Second", "New second item"}, &reply)

	client.Call("API.DeleteItem", a, &reply)
	client.Call("API.GetDB", "", &db)
	fmt.Println("Database 2: ", db)

	client.Call("API.GetByName", "First", &reply)
	fmt.Println("First item", reply)

}
