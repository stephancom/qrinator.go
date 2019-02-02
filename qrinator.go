package main

import (
  "fmt"
  "log"
  "bytes"
  "image"
  "strconv"
  "net/http"
  "github.com/go-redis/redis"
  "github.com/disintegration/imaging"
  qrcode "github.com/skip2/go-qrcode"
)

const BaseUrl = "http://github.com"
const LogoUrl = "https://github.githubassets.com/images/modules/logos_page/Octocat.png"
const Size    = 384
const Offset  = Size / 3
const Inset   = (Size - Offset) / 2

var logo image.Image
var redisdb *redis.Client

func init() {
  log.Println("init redis")
  redisdb = redis.NewClient(&redis.Options{
    Addr:     "localhost:6379",
    Password: "", // no password set
    DB:       0,  // use default DB
  })

  log.Println("read logo")
  response, _ := http.Get(LogoUrl)
  defer response.Body.Close()
  logo, _, _ = image.Decode(response.Body)
  logo = imaging.Resize(logo, Inset, Inset, imaging.Lanczos)
}

func main() {
  http.HandleFunc("/", handler)
  log.Println("Listening on :8080")
  http.ListenAndServe(":8080", nil)
}

func url(payload string) string {
  return BaseUrl + payload
}

func buildQr(payload string) image.Image {
    img, _ := qrcode.New(url(payload), qrcode.Medium)
    image := img.Image(Size)
    image = imaging.OverlayCenter(image, logo, 1.0)
    return image
}

func cachePng(payload string) *bytes.Buffer {
  cached, err := redisdb.Get(payload).Result()
  png := new(bytes.Buffer)

  if err == redis.Nil {
    log.Println("cache miss, generating "+payload)
    imaging.Encode(png, buildQr(payload), imaging.PNG)
    redisdb.Set(payload, png.String(), 0)
  } else {
    png = bytes.NewBufferString(cached)
  }

  return png
}

func handler(w http.ResponseWriter, r *http.Request) {
  if r.Method == "DELETE" {
    log.Println("clear cache")
    fmt.Fprintf(w, "clear cache")
    redisdb.FlushDB()
  } else {
    buffer := cachePng(r.URL.Path)

    w.Header().Set("Content-Type", "image/png")
    w.Header().Set("Content-Length", strconv.Itoa(buffer.Len()))
    w.Write(buffer.Bytes())
  }
}
