package main

import (
  "fmt"
  "os"
  "os/signal"
  "net"
)

func main() {

  c := make(chan os.Signal, 1)
  signal.Notify(c, os.Interrupt, os.Kill)

  next := make(chan bool, 1)
  next <- true
  tcp := make(chan *net.Conn, 1)

  go tcpServer(tcp, next)
    select {
      case s := <-c:
        switch s.String() {
        case "interrupt":
          fmt.Println("Interruption. Restarting")
        case "kill":
          fmt.Println("Stopping Program")
        default:
          fmt.Println(s.String())
        }
      case t := <-tcp:
        fmt.Println(t)
        (*t).Close()
    }

}

func tcpServer (c chan *net.Conn, next chan bool) {
  addr, err := net.ResolveTCPAddr("tcp4", ":8888")
  if err != nil {
    fmt.Fprintf(os.Stderr, err.Error())
  }
  listener, err := net.ListenTCP("tcp", addr)
  if err != nil {
    fmt.Fprintf(os.Stderr, err.Error())
  }
  select {
  case <-next:
    conn,err := listener.Accept()
    if err != nil { fmt.Println("asdf") }
    c <- &conn
  }
}


