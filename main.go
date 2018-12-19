package main

import (
	"fmt"
	"strings"

	"github.com/jacobsa/go-serial/serial"
	log "github.com/sirupsen/logrus"
)

func checkCommand(c []byte) bool {
	//fmt.Println(c)
	//TODO: Search how to compare two strings, or []byte with a string
	if strings.Contains(string(c), ">WAIT\r\n") {
		fmt.Println(string(c))
		fmt.Println("Esperado!")
		return true
	}

	return false
}

func main() {

	fmt.Println("** TESTING SERIAL READER **")

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

	cIndex := 0
	stringBuf := make([]byte, 64)

	for {

		charBuf := make([]byte, 1)

		_, err := port.Read(charBuf)
		if err != nil {
			fmt.Println("Error")
		}

		if charBuf[0] != byte(0) {
			//fmt.Printf("%d %c\r\n", charBuf, charBuf)    \r=10  \n=13
			//fmt.Printf("%s\r\n", stringBuf)
			stringBuf[cIndex] = charBuf[0]
			cIndex++

			if charBuf[0] == byte(10) {
				//fmt.Printf("%s\r\n", stringBuf)
				checkCommand(stringBuf)
				cIndex = 0
				stringBuf = make([]byte, 64)
			}

		}

	}
}
