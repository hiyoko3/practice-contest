package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello, World")
}

func sandbox() {
	fmt.Println("Welcome to the playground!")

	fmt.Println("The time is", time.Now())
}

func main() {
	fmt.Println("===== tour 1 =====")
	hello()
	sandbox()
}
