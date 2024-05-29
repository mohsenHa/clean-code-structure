package main

import (
	"bufio"
	"fmt"
	"os"
)

func newAll() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter your name")
	scanner.Scan()
	name := scanner.Text()

	newHandler(name)
	newValidator(name)
	newParam(name)
	newService(name)

	fmt.Println("Successfully")

}
