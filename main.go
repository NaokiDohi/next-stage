package main

import (
	"fmt"
	"log"
)

func init() {
	log.SetPrefix("[Test]")
}

func main() {
	log.Println("Hello world")
	fmt.Println("Hello world")
}
