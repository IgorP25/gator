package main

import (
	"fmt"

	"github.com/IgorP25/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	cfg.SetUser("igor")
	cfg, err = config.Read()
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println(cfg)
}
