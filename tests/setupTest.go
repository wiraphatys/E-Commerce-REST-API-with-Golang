package tests

import (
	"encoding/json"

	"github.com/wiraphatys/E-Commerce-REST-API-with-Golang/config"
	"github.com/wiraphatys/E-Commerce-REST-API-with-Golang/modules/servers"
	"github.com/wiraphatys/E-Commerce-REST-API-with-Golang/pkg/databases"
)

func SetupTest() servers.IModuleFactory {
	cfg := config.LoadConfig("../.env.test")

	db := databases.DbConnect(cfg.Db())

	s := servers.NewServer(cfg, db)
	return servers.InitModule(nil, s.GetServer(), nil)
}

func CompressToJSON(obj any) string {
	result, _ := json.Marshal(&obj)
	return string(result)
}
