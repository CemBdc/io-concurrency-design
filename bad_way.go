package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func getUsers() []UserData {
	url := "https://gorest.co.in/public/v1/users"

	res := getRequest(url)

	if res == nil {
		return nil
	}

	var users User

	err := json.Unmarshal(res, &users)

	failOnError(err, "Error on parsing response")

	return users.Data
}

func getUserTodos() []UserTodoData {
	url := "https://gorest.co.in/public/v1/todos"

	res := getRequest(url)

	if res == nil {
		return nil
	}

	var todo Todo

	err := json.Unmarshal(res, &todo)

	failOnError(err, "Error on parsing response")

	return todo.Data
}

func doItBadway() {

	defer elapsed("badWay")()

	users := getUsers()
	todos := getUserTodos()

	if users != nil {

		userFile, err := os.Create("users.txt")

		failOnError(err, "Error on creating users.txt")

		defer userFile.Close()

		for _, u := range users {
			fmt.Printf("Id: %d, Name: %s, Email: %s\n", u.Id, u.Name, u.Email)
			fmt.Fprintf(userFile, "Id: %d, Name: %s, Email: %s\n", u.Id, u.Name, u.Email)
		}
	}

	if todos != nil {

		todoFile, err := os.Create("todos.txt")

		failOnError(err, "Error on creating todos.txt")

		defer todoFile.Close()

		for _, u := range todos {
			fmt.Printf("Id: %d, UserId: %d, Title: %s\n", u.Id, u.UserId, u.Title)
			fmt.Fprintf(todoFile, "Id: %d, UserId: %d, Title: %s\n", u.Id, u.UserId, u.Title)
		}
	}

}
