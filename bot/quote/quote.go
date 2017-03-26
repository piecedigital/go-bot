package quote

import (
  "net/http"
  "log"
  "fmt"
  "io/ioutil"
  // "net/url"
  "encoding/json"
  "bytes"
)

func Quote(op string, val string, user string) string {
  // fmt.Println("quote info", op, val)
  // fmt.Println(op, val)
  body := makeRequest(op, val, user)
  return body
}

func makeRequest(op string, val string, user string) string {
  object := map[string]string{"op": op, "val": val, "quoter": user}
  jsonData, _ := json.Marshal(object)
  resp, err := http.Post("http://localhost:8090",
                "application/json",
                bytes.NewBuffer(jsonData))
  if err != nil {
    fmt.Println(err)
    return "problem making request"
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(string(body))
  return string(body)
}
