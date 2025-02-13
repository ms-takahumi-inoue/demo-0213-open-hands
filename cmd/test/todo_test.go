package main

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestGetTodos(t *testing.T) {
    store := &TodoStore{
        todos:  make(map[int]Todo),
        nextID: 1,
    }
    store.todos[1] = Todo{ID: 1, Text: "Test Todo", Completed: false}

    req, err := http.NewRequest("GET", "/todos", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(store.GetTodos)
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    var todos []Todo
    json.NewDecoder(rr.Body).Decode(&todos)
    if len(todos) != 1 || todos[0].Text != "Test Todo" {
        t.Errorf("handler returned unexpected body: got %v", rr.Body.String())
    }
}

func TestAddTodo(t *testing.T) {
    store := &TodoStore{
        todos:  make(map[int]Todo),
        nextID: 1,
    }

    todo := Todo{Text: "New Todo"}
    body, _ := json.Marshal(todo)
    req, err := http.NewRequest("POST", "/todos", bytes.NewBuffer(body))
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(store.AddTodo)
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    var addedTodo Todo
    json.NewDecoder(rr.Body).Decode(&addedTodo)
    if addedTodo.Text != "New Todo" {
        t.Errorf("handler returned unexpected body: got %v", rr.Body.String())
    }
}