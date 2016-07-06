package main

import (
  "net/http"
  "log"
  "./routes"
)

func main()  {
  http.HandleFunc("/", routes.PipeRequests)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
