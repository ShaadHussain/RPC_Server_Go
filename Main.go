package main

import "fmt"

type Item struct {
	title string
	body  string
}

type API int

var database []Item

func (a *API) GetByName(title string, reply *Item) error {

	var getItem Item

	for _, val := range database {
		if val.title == title {
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

		if val.title == edit.title {

			// fmt.Println("Idx: ", idx)
			database[idx] = Item{edit.title, edit.body}
			changed = database[idx]

		}

	}

	*reply = changed

	return nil
}

func DeleteItem(item Item) Item {
	var del Item

	for idx, val := range database {

		if val.title == item.title && val.body == item.body {

			database = append(database[:idx], database[idx+1:]...)

			del = item
			break

		}
	}

	return del
}

func main() {

	fmt.Println("initial db: ", database)
	a := Item{"first", "test item 1"}
	b := Item{"2", "test item 2"}
	c := Item{"third", "test item 3"}

	AddItem(a)
	AddItem(b)
	AddItem(c)

	fmt.Println("2nd db", database)

	DeleteItem(b)

	fmt.Println("db 3", database)

	EditItem("third", Item{"fourth", "3's rplced"})

	fmt.Println("db 4", database)

	x := GetByName("3 new")
	y := GetByName("first")

	fmt.Println(x, y)

}
