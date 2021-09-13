package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

var userChan = make(chan UserData, 20)
var todoChan = make(chan UserTodoData, 20)

func getUserList(wg *sync.WaitGroup) {

	url := "https://gorest.co.in/public/v1/users"

	res := getRequest(url)

	if res == nil {
		return
	}

	var users User

	err := json.Unmarshal(res, &users)

	failOnError(err, "Error on parsing response")

	for _, v := range users.Data {
		userChan <- v
	}

	wg.Done()

	close(userChan)
}

func getUserTodoList(wg *sync.WaitGroup) {

	url := "https://gorest.co.in/public/v1/todos"

	res := getRequest(url)

	if res == nil {
		return
	}

	var todo Todo
	err := json.Unmarshal(res, &todo)

	failOnError(err, "Error on parsing response")

	for _, v := range todo.Data {
		todoChan <- v
	}

	wg.Done()

	close(todoChan)
}

func userToConsole(u UserData, wg *sync.WaitGroup) {

	fmt.Printf("Id: %d, Name: %s, Email: %s\n", u.Id, u.Name, u.Email)

	wg.Done()

}

func userToFile(u UserData, file *os.File, wg *sync.WaitGroup) {

	fmt.Fprintf(file, "Id: %d, Name: %s, Email: %s\n", u.Id, u.Name, u.Email)

	wg.Done()

}

func todoConsole(t UserTodoData, wg *sync.WaitGroup) {

	fmt.Printf("Id: %d, UserId: %d, Title: %s\n", t.Id, t.UserId, t.Title)

	wg.Done()

}

func todoToFile(t UserTodoData, file *os.File, wg *sync.WaitGroup) {

	fmt.Fprintf(file, "Id: %d, UserId: %d, Title: %s\n", t.Id, t.UserId, t.Title)

	wg.Done()

}

func doItGoodway() {
	defer elapsed("goodWay")()

	var wg sync.WaitGroup

	userFile, err := os.Create("users.txt")
	failOnError(err, "Error on creating users.txt")
	defer userFile.Close()

	todoFile, err := os.Create("todos.txt")
	failOnError(err, "Error on creating users.txt")
	defer todoFile.Close()

	userOpen := true
	todoOpen := true

	wg.Add(2)
	go getUserList(&wg)
	go getUserTodoList(&wg)

	for userOpen || todoOpen {
		select {
		case user, open := <-userChan:

			if open {
				wg.Add(2)
				go userToFile(user, userFile, &wg)
				go userToConsole(user, &wg)
			} else {
				userOpen = false
			}

		case todo, open := <-todoChan:

			if open {
				wg.Add(2)
				go todoToFile(todo, todoFile, &wg)
				go todoConsole(todo, &wg)
			} else {
				todoOpen = false
			}
		}
	}

	fmt.Println("All spawned now waiting....")
	wg.Wait()
	fmt.Println("All done....")

}
