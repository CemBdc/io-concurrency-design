package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func getUsers() []UserData {
	url := "https://gorest.co.in/public/v1/users"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {

		fmt.Printf("Something went wrong on while wrapping request to %s", url)
		return nil
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)

	if err != nil {

		fmt.Printf("Something went wrong on while requesting to %s", url)
		return nil
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {

		fmt.Printf("Something went wrong on while response from %s", url)
		return nil
	}

	var users User
	if err := json.Unmarshal(body, &users); err != nil {
		fmt.Printf("Something went wrong on while parsing response from %s", url)
	}

	return users.Data
}

func getUserTodos() []UserTodoData {
	url := "https://gorest.co.in/public/v1/todos"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {

		fmt.Printf("Something went wrong on while wrapping request to %s", url)
		return nil
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)

	if err != nil {

		fmt.Printf("Something went wrong on while requesting to %s", url)
		return nil
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {

		fmt.Printf("Something went wrong on while response from %s", url)
		return nil
	}

	var todo Todo
	if err := json.Unmarshal(body, &todo); err != nil {
		fmt.Printf("Something went wrong on while parsing response from %s", url)
	}

	return todo.Data
}

func doItBadway() {

	defer elapsed("badWay")()

	users := getUsers()
	todos := getUserTodos()

	if users != nil {

		userFile, err := os.Create("users.txt")

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		defer userFile.Close()

		for _, u := range users {
			fmt.Printf("Id: %d, Name: %s, Email: %s\n", u.Id, u.Name, u.Email)
			fmt.Fprintf(userFile, "Id: %d, Name: %s, Email: %s\n", u.Id, u.Name, u.Email)
		}
	}

	if todos != nil {

		todoFile, err := os.Create("todos.txt")

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		defer todoFile.Close()

		for _, u := range todos {
			fmt.Printf("Id: %d, UserId: %d, Title: %s\n", u.Id, u.UserId, u.Title)
			fmt.Fprintf(todoFile, "Id: %d, UserId: %d, Title: %s\n", u.Id, u.UserId, u.Title)
		}
	}

}
