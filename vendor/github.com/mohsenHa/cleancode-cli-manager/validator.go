package cli_cleancode_manager

import (
	"bufio"
	"fmt"
	"github.com/mohsenHa/cleancode-cli-manager/blueprint/tmpl_validator"
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
			content:    tmpl_validator.GetSampleTmpl(),
			tmplPath:   filepath.Join("cmd", "manager", "blueprint", "validator", "sample.tmpl"),
			params:     name,
			targetPath: targetPath,
		},
		{
			filename:   "validator.go",
			content:    tmpl_validator.GetValidatorTmpl(),
			tmplPath:   filepath.Join("cmd", "manager", "blueprint", "validator", "validator.tmpl"),
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
