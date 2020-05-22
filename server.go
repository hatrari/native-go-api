package main

import "net/http"

type User struct {
	Id		string
	Name	string
}

type usersHandler struct {
	store map[string]User
}

func (h *usersHandler) get(w http.ResponseWriter, r *http.Request) {
}

func newUsersHandler() *usersHandler {
	return &usersHandler {
		store: map[string]User{},
	}
}

func main() {
	usersHandler := newUsersHandler()
	http.HandleFunc("/users", usersHandler.get)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
