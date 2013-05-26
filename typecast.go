package main

import (
  "fmt"
)

type mystring string

func (m *mystring) modify() {
  *m = mystring("NEW")
}
func (m mystring) dont_modify() {
  m = mystring("NEW");
}

func main() {
  var x mystring
  x = mystring("OLD")
  fmt.Print(x + "\nCalling dont_modify(): x=")
  x.dont_modify()
  fmt.Print(x + "\nCalling modify(): x=")
  x.modify()
  fmt.Println(x)
}
