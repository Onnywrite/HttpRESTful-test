package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/Onnywrite/HttpRESTful-test/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config")
}

func main() {
	flag.Parse()

	conf := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, conf)
	if err != nil {
		log.Fatal(err)
	}

	serv := apiserver.New(conf)
	if err = serv.Start(); err != nil {
		log.Fatal(err)
	}
}
