package main

import (
	"flag"
	"log"

	"github.com/Shuba-Buba/Trading-propper-backtest/postman"
	"github.com/Shuba-Buba/Trading-propper-backtest/saver"
)

func main() {
	var mode string
	var configPath string
	var postmanPort uint
	flag.StringVar(&mode, "mode", "client", "mode to run the program")
	flag.StringVar(&configPath, "config", "", "config path")
	flag.UintVar(&postmanPort, "postmanPort", 7777, "postman port to listen connection requests")
	flag.Parse()

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	switch mode {
	case "postman":
		postman.Init(configPath, postmanPort)
	case "saver":
		saver.Init(configPath, postmanPort)
	default:
		log.Fatal("Invalid mode")
	}
}
