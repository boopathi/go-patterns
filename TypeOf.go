package main

import (
  "os"
  "fmt"
  "reflect"
  "net"
)

func main() {
  a := make([]interface{},10)
  a[0] = make(chan int, 4)
  a[1] = 45
  a[2] = 23.902
  a[3] = "asdf"
  a[4] = 'e'
  a[5] = make([]int, 3)
  a[6] = make(map[string] string)
  a[7] = new(os.Signal)
  a[8] = new(net.Conn)
  fmt.Println("Type of a =", reflect.TypeOf(a))
  for i:=0;i<len(a);i++ {
    fmt.Println("Type of a[", i ,"] =", reflect.TypeOf(a[i]))
  }
}
