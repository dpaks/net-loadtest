package main

import (
	"fmt"
	"net/http"
)

func Respond(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hit!")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Processed"))
}

func main() {
	http.HandleFunc("/", Respond)
	fmt.Println("serving")
	panic(http.ListenAndServe(":3001", nil))
}
