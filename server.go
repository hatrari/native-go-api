package main

import "net/http"
import "encoding/json"

type User struct {
	Id		string
	Name	string
}

type usersHandler struct {
	store map[string]User
}

func (h *usersHandler) get(w http.ResponseWriter, r *http.Request) {
	users := make([]User, len(h.store))

	i := 0
	for _, user := range h.store {
		users[i] = user
		i++
	}

	jsonBytes, err := json.Marshal(users)
	if err != nil {
		// todo
	}

	w.Write(jsonBytes)
}

func newUsersHandler() *usersHandler {
	return &usersHandler {
		store: map[string]User{
			"id1": User{"aaaa","name"},
		},
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
