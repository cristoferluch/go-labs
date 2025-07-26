package qrcode

import (
	"github.com/mdp/qrterminal/v3"
	"os"
)

func Generate(url string) {

	config := qrterminal.Config{
		Level:     qrterminal.M,
		Writer:    os.Stdout,
		BlackChar: qrterminal.WHITE,
		WhiteChar: qrterminal.BLACK,
		QuietZone: 1,
	}

	qrterminal.GenerateWithConfig(url, config)
}
