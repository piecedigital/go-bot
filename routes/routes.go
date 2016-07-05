package routes

import (
  "fmt"
  "io"
  "net/http"
  "regexp"
)

var methodMap = map[string]map[string](func(res http.ResponseWriter, req *http.Request) int) {
  "GET": pathMapGET,
}

var pathMapGET = map[string](func(res http.ResponseWriter, req *http.Request) int) {
  "/": indexGet,
  "/bot": botGet,
}

func PipeRequests(res http.ResponseWriter, req *http.Request) {
  req.Header.Add("Content-Type", "text/html; charset=utf-8")
  method, path := req.Method, req.URL.Path
  cb, okay := methodMap[method][path]
  // match the path to see if it's trying to get a static file or not
  match, _ := regexp.MatchString("/public/.*", path)

  // serve static files if the path matches the public URL
  if match {
    http.ServeFile(res, req, "." + path)
    fmt.Printf("%v %v - %v\n", method, path, 200)
  } else
  // else, call the appropriate function
  if okay {
    status := cb(res, req)
    fmt.Printf("%v %v - %v\n", method, path, status)
  } else {
    // else, failure, serve 404 text
    fourOhFourGet(res, req)
    fmt.Printf("%v %v - RESPONSE: '404: Not found'\n", method, path)
  }
}

func fourOhFourGet(res http.ResponseWriter, req *http.Request) {
  io.WriteString(res, "404: Not Found\n")
}

func indexGet(res http.ResponseWriter, req *http.Request) int {
  http.ServeFile(res, req, "./views/index.html")
  // io.WriteString(res, "<h1>Hello World!</h1>")
  return 200
}

func botGet(res http.ResponseWriter, req *http.Request) int {
  io.WriteString(res, "<h1>This will be the bot page.</h1>")
  return 200
}
