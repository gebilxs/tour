package main

import (
	"flag"
	"log"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "Hello, world!", "Help")
	flag.StringVar(&name, "n", "Hello,world!", "Help")
	flag.Parse()

	log.Printf("name: %s", name)
}
