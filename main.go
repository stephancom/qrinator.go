package main

import (
  "fmt"
  "net/http"
)

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
  if r.Method == "DELETE" {
    fmt.Fprintf(w, "clear cache")  
  } else {
    fmt.Fprintf(w, "build QR")  
  }
}