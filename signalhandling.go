package main

import (
  "fmt"
  "os"
  "os/signal"
  "time"
)

func main() {
  c := make(chan os.Signal, 1)
  signal.Notify(c, os.Interrupt, os.Kill)

  time.Sleep(time.Second * 5)

  select {
  case s := <-c:
    fmt.Println(s.String())
  default:
    fmt.Println("No signal received till here.")
  }
}
