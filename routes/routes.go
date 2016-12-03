package routes

import (
  "fmt"
  "io"
  "path"
  "runtime"
  "net/http"
  "regexp"
  "github.com/piecedigital/go-bot/socket"
)

// a channel to tell it to stop
var stopChan = make(chan int)

var methodMap = map[string]map[string](func(res http.ResponseWriter, req *http.Request) int) {
  "GET": pathMapGET,
}

var pathMapGET = map[string](func(res http.ResponseWriter, req *http.Request) int) {
  "/": indexGet,
  "/bot-page": botGet,
  "/bot": startSockets,
  "/stop-bot": stopSockets,
}

func PipeRequests(res http.ResponseWriter, req *http.Request) {
  req.Header.Add("Content-Type", "text/html; charset=utf-8")
  method, path := req.Method, req.URL.Path
  cb, okay := methodMap[method][path]
  // match the path to see if it's trying to get a static file or not
  match, _ := regexp.MatchString("/public/.*", path)

  // serve static files if the path matches the public URL
  if match {
    http.ServeFile(res, req, getPath("/.." + path))
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
  // fmt.Println("404 page")
  io.WriteString(res, "404: Not Found\n")
}

func indexGet(res http.ResponseWriter, req *http.Request) int {
  // fmt.Println("index page")
  http.ServeFile(res, req, getPath("/../views/index.html"))
  // io.WriteString(res, "<h1>Hello World!</h1>")
  return 200
}

func botGet(res http.ResponseWriter, req *http.Request) int {
  // fmt.Println("bot page")
  http.ServeFile(res, req, getPath("/../views/bot-page.html"))
  return 200
}

func startSockets(res http.ResponseWriter, req *http.Request) int {
  req.ParseForm()
  channelName := "piecedigital"
  form := req.Form
  // fmt.Println(form)
  if form != nil {
    queryArray := form["channel"]
    // fmt.Println(form)
    if queryArray != nil {
      name := queryArray[0]
      // fmt.Println(form)
      if name != "" {
        channelName = name
      }
    }
  }

  fmt.Printf("channel: %v, type: %T\n", channelName, channelName)

  var statusChan = make(chan int)

  go socket.StartSockets(res, req, channelName, statusChan, stopChan)
  // return 200
  statusCode := <- statusChan
  // return 200
  return statusCode
}

func stopSockets(res http.ResponseWriter, req *http.Request) int {
  close(stopChan)
  return 200
}

func getPath(pathString string) string {
  _, filename, _, _ := runtime.Caller(0)
  // fmt.Printf("filename: %v\n", filename)
  finalPath := path.Dir(filename) + pathString
  // fmt.Printf("Final path: %v\n", finalPath)
  return finalPath
}
