package main

import (
	"flag"
	"fmt"
	"os"
	"seminarioGo/trabajoSeminario/internal/config"
)

func main() {

	archConfig := flag.String("config", "./config/config.yaml", "...")
	flag.Parse()

	cfg, err := config.LoadConfig(*archConfig)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(cfg.DB.Driver)
	fmt.Println(cfg.Version)

}
