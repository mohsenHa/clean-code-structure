package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func param() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter your param name")
	scanner.Scan()
	name := scanner.Text()

	newParam(name)

	fmt.Println("Successfully")

}

func newParam(name string) {
	targetPath := filepath.Join("param", name+"param")
	files := []file{
		{
			filename:   "sample.go",
			tmplPath:   filepath.Join("cmd", "manager", "blueprint", "param", "sample.tmpl"),
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
