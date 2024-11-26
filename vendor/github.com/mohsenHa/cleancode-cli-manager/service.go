package cli_cleancode_manager

import (
	"bufio"
	"fmt"
	"github.com/mohsenHa/cleancode-cli-manager/blueprint/tmpl_service"
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
			content:    tmpl_service.GetSampleTmpl(),
			tmplPath:   filepath.Join("cmd", "manager", "blueprint", "service", "sample.tmpl"),
			params:     name,
			targetPath: targetPath,
		},
		{
			filename:   "service.go",
			content:    tmpl_service.GetServiceTmpl(),
			tmplPath:   filepath.Join("cmd", "manager", "blueprint", "service", "service.tmpl"),
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
