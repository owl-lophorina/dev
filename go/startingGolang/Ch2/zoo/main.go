package main

import (
	"fmt"

	"github.com/owl-lophorina/dev/go/startingGolang/Ch2/zoo/animals"
)

func main() {
	fmt.Println(AppName())

	fmt.Println(animals.ElephantFeed())
	fmt.Println(animals.MonkeyFeed())
	fmt.Println(animals.RabbitFeed())
}
