package main

import (
  "fmt"
  "time"
)

type Ball struct {
  hits int
}

func main() {
  table := make(chan *Ball)
  go player("ping", table)
  go player("pong", table)

  table <- new(Ball) //toss the ball
  time.Sleep(time.Second * 3)
  <-table
  panic("Printing the Stack trace: ")
}

func player(name string, table chan *Ball) {
  for {
    ball := <-table // receive the ball hit
    ball.hits++
    fmt.Println(name, ball.hits)
    time.Sleep(500 * time.Millisecond) //take some time
    table <- ball //hit back
  }
}
