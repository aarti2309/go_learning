package main

import "fmt"
const ok = true
var a = 10
var decalare = 10
func main(){
  var hello = "Hello"
  fmt.Println("Hello this is my first program")
  hey()
  fmt.Println(hello,ok,a)
  fmt.Println("inside main:", decalare)
  nested()
  fmt.Println("inside main again:", decalare)
  
}

func nested(){
  var decalare = 5
  fmt.Println("inside nested function:", decalare)
}
