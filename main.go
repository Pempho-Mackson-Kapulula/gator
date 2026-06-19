package main

import (
	"fmt"
	"log"

	"github.com/Pempho-Mackson-Kapulula/gator/internal/config"
)

func main() {
	//read config file
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	//set name to cfg struct
	err = cfg.SetUser("Pempho")
	if err != nil {
		log.Fatal(err)
	}

	//read and print config
	cfgData, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfgData)
 
}
