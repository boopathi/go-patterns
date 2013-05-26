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
  //handle the signal
  go func() {
    s := <-c
    fmt.Println(s.String())
    os.Exit()
  }
  //Some task that takes 5 seconds to complete
  time.Sleep(time.Second * 5)
  fmt.Println("No signal received till here.")
}
