package main

import (
	"fmt"

	"github.com/ashraddhansh/gator/internal/config"
)


func main() {
	config, err := config.Read()
	if err != nil {
		fmt.Printf("Error Reading File: %v", err)
	}
	err = config.SetUser("ashraddhansh")
	if err != nil {
		fmt.Printf("Error Writing File: %v", err)
	}
	fmt.Printf("%v", config)
}
