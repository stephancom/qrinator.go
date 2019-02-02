package main

import (
  "fmt"
  "log"
  "bytes"
  "image"
  "image/color"
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

func url(payload string) string {
  return "http://stephan.com" + payload
}

func buildQr(payload string) image.Image {
  img, err := qrcode.New(url(payload), qrcode.Medium)
  if err != nil { log.Fatalf("unable to encode") }
  image := img.Image(256)
  // placeholder for processing to come
  image = imaging.Rotate(image, 10.0, color.Black)
  return image
}

func handler(w http.ResponseWriter, r *http.Request) {
  if r.Method == "DELETE" {
    fmt.Fprintf(w, "clear cache")  
  } else {
    buffer := new(bytes.Buffer)
    imaging.Encode(buffer, buildQr(r.URL.Path), imaging.PNG)

    w.Header().Set("Content-Type", "image/png")
    w.Header().Set("Content-Length", strconv.Itoa(buffer.Len()))
    w.Write(buffer.Bytes())
  }
}
