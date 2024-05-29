package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func validator() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter your name")
	scanner.Scan()
	name := scanner.Text()

	newValidator(name)

	fmt.Println("Successfully")

}

func newValidator(name string) {
	targetPath := filepath.Join("validator", name+"validator")
	files := []file{
		{
			filename:   "sample.go",
			tmplPath:   filepath.Join("cmd", "manager", "blueprint", "validator", "sample.tmpl"),
			targetPath: targetPath,
			params:     name,
		},
		{
			filename:   "validator.go",
			tmplPath:   filepath.Join("cmd", "manager", "blueprint", "validator", "validator.tmpl"),
			targetPath: targetPath,
			params:     name,
		},
	}
	for _, f := range files {
		fmt.Print(f.filename)
		err := create(f)
		if err != nil {
			fmt.Println()
			fmt.Println(err)
			continue
		}
		fmt.Println(" Done")
	}

}
