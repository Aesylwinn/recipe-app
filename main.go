package main

import (
  "fmt"
  "net/http"

  "github.com/gorilla/mux"
)

func main() {
  fmt.Println("Starting server")
  r := mux.NewRouter()
  r.HandleFunc("/", TestHandler)
  http.ListenAndServe(":8080", r)
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Hello World!\n"))
}

