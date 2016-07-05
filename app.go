package main

import (
  // "fmt"
  // "net/http"
  // "log"
  // "./routes"
  "./bot"
  // "./private"
)

func main()  {
  // http.HandleFunc("/", routes.PipeRequests)
  // log.Fatal(http.ListenAndServe(":8080", nil))
  bot.Connect()
}
