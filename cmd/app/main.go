package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/pyankovzhe/auth/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if os.Getenv("SECRET_JWT_KEY") == "" {
		log.Fatal("Need to specify SECRET_JWT_KEY env")
	}

	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(config, ctx); err != nil {
		log.Fatal(err)
	}
}
