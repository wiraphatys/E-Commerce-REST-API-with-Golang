package main

import (
	"fmt"
	"os"

	"github.com/wiraphatys/E-Commerce-REST-API-with-Golang/config"
)

func envPath() string {
	if len(os.Args) == 1 {
		return ".env"
	} else {
		return os.Args[1]
	}
}

func main() {
	cfg := config.LoadConfig(envPath())
	fmt.Println(cfg.App())
}
