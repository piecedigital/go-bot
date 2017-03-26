package main

import (
  "bytes"
  "os"
  "os/exec"
  "fmt"
  "log"
)

func main() {
  if len(os.Args) >= 2 {
    switch os.Args[1] {
      case "get":
        quote(os.Args[1], "")
        break;
      case "set":
        quote(os.Args[1], os.Args[2]);
        break;
    }
  }
}

func quote(op string, val string)  {
  if len(os.Args) >= 2 {
    // dir, dirErr := filepath.Abs(filepath.Dir(os.Args[0]))
    // if dirErr != nil {
    //   log.Fatal(dirErr)
    //   return
    // }
    fmt.Println(op, val)
    fmt.Println("Getting quote: ")
    cmd := exec.Command("node", "quote.js", op, val)
    // cmd := exec.CommandContext(ctx, "ls")
    stdin, err := cmd.StdinPipe()
    if err != nil {
        fmt.Println(err)
    }
    defer stdin.Close()
    // set outerr buffer
    var outerr bytes.Buffer
    cmd.Stderr = &outerr

    cmdErr := cmd.Run()
    if cmdErr != nil {
      fmt.Println(fmt.Sprint(cmdErr) + ": " + outerr.String())
      log.Fatal(cmdErr)
      return
    }

    // set output buffer
    out := new(bytes.Buffer)
    cmd.Stdout = out

    cmd.Wait()
    fmt.Printf("output: %q", out.String())
  }
}
