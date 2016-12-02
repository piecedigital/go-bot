package main

import (
  "fmt"
  "net/http"
  "log"
  "./routes"
)

func main()  {
  http.HandleFunc("/", routes.PipeRequests)
  log.Fatal(http.ListenAndServe(":8080", nil))
  fmt.Println("Now listening to localhost at port 8080")
}
