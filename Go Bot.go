package main

import (
  "fmt"
  "net/http"
  "log"
  "github.com/piecedigital/go-bot/routes"
)

func main()  {
  http.HandleFunc("/", routes.PipeRequests)
  go log.Fatal(http.ListenAndServe(":8080", nil))
  fmt.Println("Now listening to localhost at port 8080")
}
