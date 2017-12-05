package main

import (
	"rabbit-worker/rpc_server"
	"os"
	"log"
	"github.com/BurntSushi/toml"
)

var (
	cfgFile  = "Config.ini"
	Config  MainConfig
)

type MainConfig struct {
	Token string
	EndPoint EndPointConfig
	Rabbit rpc_server.RabbitConfig
}

type EndPointConfig struct {
	PAPI string
}

func main() {
	_, err := os.Stat(cfgFile)
	if err != nil {
		log.Fatal("Config file is missing: ", cfgFile)
	}


	if _, err := toml.DecodeFile(cfgFile, &Config); err != nil {
		log.Fatal(err)
	}

	rpcServer := new(rpc_server.RPCServer)
	rpcServer.Init(Config.Rabbit)
	rpcServer.Start(ProcedureCallManager{})
}