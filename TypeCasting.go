package main

import (
  "fmt"
)

type mystring string

func (m *mystring) modify() {
  *m = mystring("NEW")
}
func (m mystring) doesnt_modify() {
  m = mystring("NEW");
}

func main() {
  var x mystring
  x = mystring("OLD")
  fmt.Println("x = ", x)
  x.doesnt_modify()
  fmt.Println("x = ", x, "#doesn't modify x")
  x.modify()
  fmt.Println("x = ", x, "#modifies x")
}
