package main

import (
	"os"
	"os/signal"

	_ "github.com/mohsenHa/cleancode-cli-manager"
)

func main() {
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt)
	<-exit
}
