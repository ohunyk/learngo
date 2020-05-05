package main

import (
	"fmt"

	"github.com/ohunyk/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}

	err := dictionary.Add("hello", "Greeting")
	if err != nil {
		fmt.Println(err)
	}
	hello, err := dictionary.Search("hello")
	fmt.Println(hello)

	err2 := dictionary.Add("hello", "Greeting")
	if err2 != nil {
		fmt.Println(err2)
	}
}
