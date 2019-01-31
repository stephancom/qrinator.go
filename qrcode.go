package main

import (
  "fmt"
  "log"
  "strconv"
  "net/http"
  qrcode "github.com/skip2/go-qrcode"
)


func main() {
  http.HandleFunc("/", handler)
  log.Println("Listening on :8080")
  http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
  if r.Method == "DELETE" {
    fmt.Fprintf(w, "clear cache")  
  } else {
    var url string
    url = "http://stephan.com" + r.URL.Path
    var png[]byte
    png, _ = qrcode.Encode(url, qrcode.Medium, 256)
    w.Header().Set("Content-Type", "image/png")
    w.Header().Set("Content-Length", strconv.Itoa(len(png)))
    w.Write(png)
  }
}