package main

import (
	"discord-bot/bot"
	"discord-bot/config"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	cfg, err := config.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}

	session, err := bot.Start(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	<-exit

	fmt.Println("Shutting down...")
}
