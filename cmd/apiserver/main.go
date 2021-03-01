package main

import (
	"flag"
	"github.com/BohdanShmalko/mesGoBack/internal/app/apiserver"
	"github.com/BurntSushi/toml"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config", "./configs/apiserver.toml", "path to config.toml file")
}

func main()  {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err);
	}
}