package utils

import (
	"errors"
	"log"
	"net"
	"os"
	"strconv"
)

func ParseArgs() (string, int, error) {

	var folder string = "./"
	var port int = 8080
	var err error

	for i, arg := range os.Args {

		if arg == "-f" || arg == "--folder" {
			folder = os.Args[i+1]
		}

		if arg == "-p" || arg == "--port" {
			port, err = strconv.Atoi(os.Args[i+1])
			if err != nil {
				return "", 0, errors.New("invalid port number")
			}
		}
	}

	return folder, port, nil
}

func GetLocalIpAddress() net.IP {

	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	return conn.LocalAddr().(*net.UDPAddr).IP
}
