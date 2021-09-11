package main

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
