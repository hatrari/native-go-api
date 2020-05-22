package main

import "net/http"

type User struct {
	Id		string
	Name	string
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
}

func main() {
	http.HandleFunc("/users", usersHandler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
