// package main

// import (
// 	"fmt"
// 	"net/http"
//   "encoding/json"
// 	"github.com/gorilla/mux"
// )

// func HelloHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodGet {
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 		return
// 	}
// 	fmt.Fprintf(w, "Hello, World")
// }

// func TaskHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodPost {
//    http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 		return
// 	}
// 	type requestBody struct {
// 	Task string `json:"task"`
// }

// err := json.NewDecoder(r.Body).Decode(&requestBody)
// 	if err != nil {
// 		http.Error(w, "Invalid request body", http.StatusBadRequest)
// 		return

// 	}
// task = requestBody.Task

// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprintf(w, "Task updated successfully: %s", task)

// }

// var task string

// func main() {

// 	router := mux.NewRouter()
// 	router.HandleFunc("/task", TaskHandler).Methods("POST")
// 	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
// 	http.ListenAndServe(":8080", router)
// }

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var task string

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Hello, World")
}

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var requestData struct {
		Task string `json:"task"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	task = requestData.Task

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Task updated successfully: %s", task)
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/task", TaskHandler).Methods("POST")
	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
