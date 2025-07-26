package main

import (
	"fmt"
	"http-fileserver/internal/server"
	"http-fileserver/internal/utils"
	"log"
)

func main() {

	folder, port, err := utils.ParseArgs()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	s := server.New(port, folder)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
