package main

import (
	"devmentor-BE103-golang/infrastructure"
	"log"
)

func main() {

	dbErr := infrastructure.InitMySQL()
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	_, err := infrastructure.InitGinServer()
	if err != nil {
		log.Fatal(err)
	}
}
