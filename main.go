package main

import (
	"fmt"

	"github.com/jacobsa/go-serial/serial"
	log "github.com/sirupsen/logrus"
)

func main() {

	// Set up options.
	options := serial.OpenOptions{
		PortName:        "COM3",
		BaudRate:        57600,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}

	// Open the port.
	port, err := serial.Open(options)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}

	// Make sure to close it later.
	defer port.Close()

	for {
		charBuf := make([]byte, 1)
		_, err := port.Read(charBuf)
		if err != nil {
			fmt.Println("Error")
		}
		if charBuf[0] != byte(0) {
			fmt.Printf("%s", string(charBuf))
		}

	}
}
