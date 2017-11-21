package main

import (
	//"rabbit-worker/rpc_server"
	"os"
	"log"
	"github.com/BurntSushi/toml"
)

type Config struct {
	Rabbit RabbitConfig
}

func main() {
	_, err := os.Stat(configfile)
	if err != nil {
		log.Fatal("Config file is missing: ", configfile)
	}

	var config Config
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		log.Fatal(err)
	}
	//rpcServer := new(rpc_server.RPCServer)
	//rpcServer.Init()

}