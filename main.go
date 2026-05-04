package main

import "fmt"
var double = func(a,b int)int{
	return a+b
}
func main(){
	fmt.Println(double(7,3))
	fmt.Println("Hello world how is the going today")
}
func greetMember(name string)string{
	return "Hello, "+ name
}