package main

import (
	"os"

	"github.com/wiraphatys/E-Commerce-REST-API-with-Golang/config"
	"github.com/wiraphatys/E-Commerce-REST-API-with-Golang/modules/servers"
	"github.com/wiraphatys/E-Commerce-REST-API-with-Golang/pkg/databases"
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

	db := databases.DbConnect(cfg.Db())
	defer db.Close()

	servers.NewServer(cfg, db).Start()
}
