package cli_cleancode_manager

import (
	"bufio"
	"fmt"
	"github.com/mohsenHa/cleancode-cli-manager/blueprint/tmpl_param"
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
			content:    tmpl_param.GetSampleTmpl(),
			tmplPath:   filepath.Join("cmd", "manager", "blueprint", "param", "sample.tmpl"),
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
