package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func handler() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter your handler name")
	scanner.Scan()
	name := scanner.Text()

	newHandler(name)

	fmt.Println("Successfully")

}
func newHandler(name string) {
	targetPath := filepath.Join("delivery", "httpserver", name+"handler")
	files := []file{
		{
			filename:   "handler.go",
			tmplPath:   filepath.Join("cmd", "manager", "blueprint", "handler", "handler.tmpl"),
			targetPath: targetPath,
			params:     name,
		},
		{
			filename:   "route.go",
			tmplPath:   filepath.Join("cmd", "manager", "blueprint", "handler", "route.tmpl"),
			targetPath: targetPath,
			params:     name,
		},
		{
			filename:   "sample.go",
			tmplPath:   filepath.Join("cmd", "manager", "blueprint", "handler", "sample.tmpl"),
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
