package main

import (
	"devmentor-BE103-golang/infrastructure"
	"log"
)

func main() {
	_, err := infrastructure.InitGinServer()
	if err != nil {
		log.Fatal(err)
	}
}
