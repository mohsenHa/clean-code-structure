package cli_cleancode_manager

import (
	"bufio"
	"fmt"
	"github.com/mohsenHa/cleancode-cli-manager/blueprint/tmpl_handler"
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
			content:    tmpl_handler.GetHandlerTmpl(),
			tmplPath:   filepath.Join("cmd", "manager", "blueprint", "handler", "handler.tmpl"),
			params:     name,
			targetPath: targetPath,
		},
		{
			filename:   "route.go",
			content:    tmpl_handler.GetRoutTmpl(),
			tmplPath:   filepath.Join("cmd", "manager", "blueprint", "handler", "route.tmpl"),
			params:     name,
			targetPath: targetPath,
		},
		{
			filename:   "sample.go",
			content:    tmpl_handler.GetSampleTmpl(),
			tmplPath:   filepath.Join("cmd", "manager", "blueprint", "handler", "sample.tmpl"),
			params:     name,
			targetPath: targetPath,
		},
	}
	for _, f := range files {
		fmt.Print(filepath.Join(f.targetPath, f.filename))
		err := create(f)
		if err != nil {
			fmt.Println()
			fmt.Println(err)
			continue
		}
		fmt.Println(" Done")
	}

}
