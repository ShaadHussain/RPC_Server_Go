package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Item struct {
	Title string
	Body  string
}

type API int

var database []Item

func (a *API) GetDB(title string, reply *[]Item) error {
	*reply = database
	return nil
}

func (a *API) GetByName(Title string, reply *Item) error {

	var getItem Item

	for _, val := range database {
		if val.Title == Title {
			getItem = val

		}
	}

	*reply = getItem

	return nil
}

func (a *API) AddItem(item Item, reply *Item) error {
	database = append(database, item)

	*reply = item

	return nil
}

func (a *API) EditItem(edit Item, reply *Item) error {
	var changed Item

	// fmt.Println("In edit item")

	for idx, val := range database {
		// fmt.Println("In for ")

		// fmt.Println("Val Title: ", val.title)
		// fmt.Println("Edit Title: ", edit.title)
		// fmt.Println("In for ")

		if val.Title == edit.Title {

			// fmt.Println("Idx: ", idx)
			database[idx] = Item{edit.Title, edit.Body}
			changed = database[idx]

		}

	}

	*reply = changed

	return nil
}

func (a *API) DeleteItem(item Item, reply *Item) error {
	var del Item

	for idx, val := range database {

		if val.Title == item.Title && val.Body == item.Body {

			database = append(database[:idx], database[idx+1:]...)

			del = item
			break

		}
	}

	*reply = del

	return nil
}

func main() {

	var api = new(API)

	err := rpc.Register(api)

	if err != nil {
		log.Fatal("Error registering api", err)

	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4040")

	if err != nil {

		log.Fatal("Listener error", err)
	}

	log.Printf("Serving rpc on port %d", 4040)

	err = http.Serve(listener, nil)

	if err != nil {
		log.Fatal("Error serving", err)
	}

	// fmt.Println("initial db: ", database)
	// a := Item{"first", "test item 1"}
	// b := Item{"2", "test item 2"}
	// c := Item{"third", "test item 3"}

	// AddItem(a)
	// AddItem(b)
	// AddItem(c)

	// fmt.Println("2nd db", database)

	// DeleteItem(b)

	// fmt.Println("db 3", database)

	// EditItem("third", Item{"fourth", "3's rplced"})

	// fmt.Println("db 4", database)

	// x := GetByName("3 new")
	// y := GetByName("first")

	// fmt.Println(x, y)

}
