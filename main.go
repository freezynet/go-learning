package main

import (
	"awesomeProject1/hello"
	"fmt"
	"github.com/freezynet/go-learning/say"
	"github.com/freezynet/go-learning/hello"
)

func main()  {

	fmt.Println("Starting gogogo")
	fmt.Println(say.CallFromSay() + hello.CallFromHello())
}