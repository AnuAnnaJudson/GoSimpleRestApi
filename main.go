package main

import (
	"17-assignment/handlers"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", Logger(handlers.GetAllStudentsData))
	http.HandleFunc("/students", Logger(handlers.GetStudentFromId))

	http.ListenAndServe(":8080", nil)
}

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("request url: %s request method: %s\n", r.URL, r.Method)
		next(w, r)
	}

}
