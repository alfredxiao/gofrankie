package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/alfredxiao/gofrankie/rest"
)

func main() {
	var port = 8080
	if len(os.Args) >= 2 {
		portStr := os.Args[1]
		var err error

		port, err = strconv.Atoi(portStr)
		if err != nil {
			fmt.Printf("port number %s is not integer\n", portStr)
			fmt.Println("Usage: gofrankie <PORT>\n     PORT defaults to 8080")
			return
		}
	}

	rest.StartThenWait(port)
}
