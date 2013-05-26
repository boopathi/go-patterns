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
    fmt.Println("\nReceived ", s.String())
    os.Exit(0)
  }()

  //Some code
  fmt.Print("Press CTRL+C to Interrupt.")
  for i:=0;i<25;i++ {
    fmt.Print(".")
    time.Sleep(time.Millisecond * 100)
  }

  //End of the Program
  fmt.Println("\nNo signal received till here.")
}
