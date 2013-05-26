package main

import (
  "flag"
  "fmt"
  "net"
  "os"
  "os/Signal"
  "time"
)

func check(e Error) {
  if e != nil {
    fmt.Fprintf(os.Stderr, err.Error())
  }
}

func TCPServer(port int) {
  addr, err := net.ResolveTCPAddr("tcp4",":"+port.String())
  check(err)
  listener, err := net.ListenTCP("tcp", addr)
  check(err)
}

func main() {
  c := make(chan os.Signal,1)
  os.Notify(c, os.Interrupt, os.Kill)
  var port int
  flag.Intvar(&port, "p", 8080, "Port Number")
  flag.Parse()

}
