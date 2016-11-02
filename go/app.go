package main

import (
  "io"
  "log"
  "net/http"
  "io/ioutil"
  "net/http/httputil"
  "net/url"
)


var target, _ = url.Parse("http://188.166.211.189");
var proxy = httputil.NewSingleHostReverseProxy(target);

const PORT = ":8000"

func handle(w http.ResponseWriter, r *http.Request) {
  content, err := ioutil.ReadFile("cache.html")
  if err != nil {
    log.Println(err)
  }
  // log.Println("I have a request: " + r.URL.Path)
  io.WriteString(w, string(content))

  // w.Header().Set("X-GoProxy", "GoProxy")
  // proxy.ServeHTTP(w, r)
}

func main() {
  log.Println("Go server started " + PORT)
  http.HandleFunc("/", handle)
  http.ListenAndServe(PORT, nil)
}
