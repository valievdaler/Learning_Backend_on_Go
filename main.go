package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var tasks = []Task{

	{ID: 1, Title: "Learning fundamentals of Go", Completed: true},
	{ID: 2, Title: "Writing REST API", Completed: false},
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, my first Backend on Go")
}

func getTaskHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func main() {
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/tasks", getTaskHandler)

	fmt.Println("Server is run on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Printf("Error %v", err)
	}
}
