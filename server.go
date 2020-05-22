package main

import "net/http"
import "encoding/json"
import "sync"
import "io/ioutil"

type User struct {
	Id		string
	Name	string
}

type usersHandler struct {
	sync.Mutex
	store map[string]User
}

func (h *usersHandler) methods(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "GET":
			h.get(w, r)
			return
		case "POST":
			h.post(w, r)
			return
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method not Allowed."))
			return
	}
}

func (h *usersHandler) get(w http.ResponseWriter, r *http.Request) {
	users := make([]User, len(h.store))

	h.Lock()
	i := 0
	for _, user := range h.store {
		users[i] = user
		i++
	}
	h.Unlock()

	jsonBytes, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *usersHandler) post(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	var user User
	err = json.Unmarshal(bodyBytes, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	h.Lock()
	h.store[user.Id] = user
	defer h.Unlock()
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
	http.HandleFunc("/users", usersHandler.methods)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
