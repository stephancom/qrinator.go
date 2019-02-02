package main

import (
  "fmt"
  "log"
  "bytes"
  "strconv"
  "net/http"
  "github.com/disintegration/imaging"
  qrcode "github.com/skip2/go-qrcode"
)

func main() {
  http.HandleFunc("/", handler)
  log.Println("Listening on :8080")
  http.ListenAndServe(":8080", nil)
}

func 

func handler(w http.ResponseWriter, r *http.Request) {
  if r.Method == "DELETE" {
    fmt.Fprintf(w, "clear cache")  
  } else {
    url := "http://stephan.com" + r.URL.Path
    var png[]byte
    png, err := qrcode.Encode(url, qrcode.Medium, 256)
    if err != nil {
      log.FatalF("unable to encode")
    }
    w.Header().Set("Content-Type", "image/png")
    w.Header().Set("Content-Length", strconv.Itoa(len(png)))

    src, err := imaging.Decode(bytes.NewReader(png))
    if err != nil {
      log.Fatalf("failed to decode image: %v", err)
    }

    imaging.Encode(w, src, imaging.PNG)
    // w.Write(png)
    // if _, err := w.Write(png); err != nil {
    //     log.Println("unable to write image.")
    // }

  }
}