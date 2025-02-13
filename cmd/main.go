package main

import (
    "encoding/json"
    "net/http"
    "sync"
    "time"

    "github.com/gorilla/mux"
)

type Todo struct {
    ID        int       `json:"id"`
    Text      string    `json:"text"`
    Completed bool      `json:"completed"`
    Deadline  time.Time `json:"deadline"`
}

type TodoStore struct {
    sync.Mutex
    todos map[int]Todo
    nextID int
}

func (store *TodoStore) GetTodos(w http.ResponseWriter, r *http.Request) {
    store.Lock()
    defer store.Unlock()
    todos := make([]Todo, 0, len(store.todos))
    for _, todo := range store.todos {
        todos = append(todos, todo)
    }
    json.NewEncoder(w).Encode(todos)
}

func (store *TodoStore) AddTodo(w http.ResponseWriter, r *http.Request) {
    store.Lock()
    defer store.Unlock()
    var todo Todo
    json.NewDecoder(r.Body).Decode(&todo)
    todo.ID = store.nextID
    store.nextID++
    store.todos[todo.ID] = todo
    json.NewEncoder(w).Encode(todo)
}

func main() {
    store := &TodoStore{
        todos:  make(map[int]Todo),
        nextID: 1,
    }

    r := mux.NewRouter()
    r.HandleFunc("/todos", store.GetTodos).Methods("GET")
    r.HandleFunc("/todos", store.AddTodo).Methods("POST")

    http.ListenAndServe(":51008", r)
}