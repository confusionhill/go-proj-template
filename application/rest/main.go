package main

import (
	"log"
	"project-template/internal/config"
)

func main() {
	conf, err := config.Init()
	if err != nil {
		log.Fatal(err)
		log.Println()
	}
	if err = application(conf); err != nil {
		log.Fatal(err)
	}
}
