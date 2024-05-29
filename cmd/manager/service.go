package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func service() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter your service name")
	scanner.Scan()
	name := scanner.Text()

	newService(name)

	fmt.Println("Successfully")

}

func newService(name string) {
	targetPath := filepath.Join("service", name+"service")
	files := []file{
		{
			filename:   "sample.go",
			tmplPath:   filepath.Join("cmd", "manager", "blueprint", "service", "sample.tmpl"),
			targetPath: targetPath,
			params:     name,
		},
		{
			filename:   "service.go",
			tmplPath:   filepath.Join("cmd", "manager", "blueprint", "service", "service.tmpl"),
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
