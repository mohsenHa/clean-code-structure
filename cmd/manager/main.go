package main

import (
	_ "github.com/mohsenHa/cleancode-cli-manager"
	"os"
	"os/signal"
)

func main() {
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt)
	<-exit
}
