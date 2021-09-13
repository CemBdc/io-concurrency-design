package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

func getUserList(c chan<- UserData, wg *sync.WaitGroup) {
	defer close(c)
	defer wg.Done()

	url := "https://gorest.co.in/public/v1/users"

	res := getRequest(url)

	if res == nil {
		return
	}

	var users User

	err := json.Unmarshal(res, &users)

	failOnError(err, "Error on parsing response")

	for _, v := range users.Data {
		c <- v
	}
}

func getUserTodoList(c chan<- UserTodoData, wg *sync.WaitGroup) {
	defer close(c)
	defer wg.Done()

	url := "https://gorest.co.in/public/v1/todos"

	res := getRequest(url)

	if res == nil {
		return
	}

	var todo Todo
	err := json.Unmarshal(res, &todo)

	failOnError(err, "Error on parsing response")

	for _, v := range todo.Data {
		c <- v
	}
}

func userToConsole(user <-chan UserData) {

	for {
		u := <-user
		fmt.Printf("Id: %d, Name: %s, Email: %s\n", u.Id, u.Name, u.Email)
	}
}

func userToFile(user <-chan UserData, file *os.File) {

	for {
		u := <-user
		fmt.Fprintf(file, "Id: %d, Name: %s, Email: %s\n", u.Id, u.Name, u.Email)
	}
}

func todoConsole(todo <-chan UserTodoData) {

	for {
		t := <-todo
		fmt.Printf("Id: %d, UserId: %d, Title: %s\n", t.Id, t.UserId, t.Title)
	}
}

func todoToFile(todo <-chan UserTodoData, file *os.File) {

	for {
		t := <-todo
		fmt.Fprintf(file, "Id: %d, UserId: %d, Title: %s\n", t.Id, t.UserId, t.Title)
	}
}

func doItGoodway() {
	defer elapsed("goodWay")()

	var wg sync.WaitGroup
	userDataChan := make(chan UserData, 20)
	userFileChan := make(chan UserData, 20)
	userConsoleChan := make(chan UserData, 20)

	userTodoDataChan := make(chan UserTodoData, 20)
	userTodoDataFileChan := make(chan UserTodoData, 20)
	userTodoDataConsoleChan := make(chan UserTodoData, 20)

	wg.Add(1)
	go getUserList(userDataChan, &wg)

	wg.Add(1)
	go getUserTodoList(userTodoDataChan, &wg)

	userFile, err := os.Create("users.txt")
	failOnError(err, "Error on creating users.txt")
	defer userFile.Close()

	todoFile, err := os.Create("todos.txt")
	failOnError(err, "Error on creating users.txt")
	defer todoFile.Close()

	go userToFile(userFileChan, userFile)
	go userToConsole(userConsoleChan)

	go todoToFile(userTodoDataFileChan, todoFile)
	go todoConsole(userTodoDataConsoleChan)

	wg.Wait()

	userOpen := true
	todoOpen := true

	for userOpen || todoOpen {
		select {
		case user, open := <-userDataChan:

			if open {
				userFileChan <- user
				userConsoleChan <- user
			} else {
				userOpen = false
			}

		case todo, open := <-userTodoDataChan:

			if open {
				userTodoDataFileChan <- todo
				userTodoDataConsoleChan <- todo
			} else {
				todoOpen = false
			}
		}
	}

}
