package main

import (
	"github.com/alfredxiao/gofrankie/rest"
)

func main() {
	rest.StartThenWait(":8080")
}
