package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	files     []string
	alligator string
	gopath    string
)

func init() {
	files = []string{
		".drone.yml",
		".gitignore",
		"CHANGELOG.md",
		"Makefile",
		"README.md",
		"glide.yaml",
		"swagger.yaml",
	}
	alligator = os.Getenv("ALLIGATOR")
	if alligator == "" {
		fmt.Print("no repository name exported please export ALLIGATOR=<appname>")
		os.Exit(1)
	}
	gopath = os.Getenv("GOPATH")
	if gopath == "" {
		fmt.Print("no GOPATH exported please export GOPATH=<gopath>")
		os.Exit(1)
	}
	gopath = gopath + "/src/github.com/GannettDigital"
}

func write() {
	// Create each file based on the template
	for _, file := range files {
		template := scrape(file)
		generatedFile := fmt.Sprintf("%s/%s/%s", gopath, alligator, file)
		err := ioutil.WriteFile(generatedFile, template, 0644)
		if err != nil {
			fmt.Printf("error writing %s, %v", generatedFile, err)
			os.Exit(1)
		}
	}
}

func scrape(file string) []byte {
	// Read the given file in as a byte array
	template, err := ioutil.ReadFile(fmt.Sprintf("templates/%s", file))
	if err != nil {
		fmt.Printf("error reading %s, %v", template, err)
		os.Exit(1)
	}

	// Replace {NAME} in the tenplates with the new repository name
	s := string(template)
	s = strings.Replace(s, "{NAME}", alligator, -1)
	template = []byte(s)

	return template
}
