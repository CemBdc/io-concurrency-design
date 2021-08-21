package main

import (
	"fmt"
	"time"
)

type User struct {
	Data []UserData `json:"data"`
}

type Todo struct {
	Data []UserTodoData `json:"data"`
}

type UserData struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserTodoData struct {
	Id     int    `json:"id"`
	UserId int    `json:"user_id"`
	Title  string `json:"title"`
}

//badWay took 1.353974541s
func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

func main() {

	doItBadway()

}
