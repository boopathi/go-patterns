package main

import (
  "fmt"
  "math/rand"
)

func main() {
  a,b := make(chan string), make(chan string)
  go func(){ a<- "a"}()
  go func(){ b<- "b"}()
  if rand.Intn(2) == 0 {
    a = nil
    fmt.Println("setting a as nil")
  } else {
    b = nil
    fmt.Println("setting b as nil")
  }
  select {
  case s := <-a:
    fmt.Println("got",s)
  case s := <-b:
    fmt.Println("got",s)
  }
}
