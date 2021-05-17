package config

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	Extension  string
	Directory  string
	SearchText string
}

func (c Config) Init() Config {
	var extension string
	var directory string
	var searchText string

	flag.StringVar(&extension, "ex", "", "input extension name for parsing\nenable values: kt/py/go")
	flag.StringVar(&searchText, "text", "", "enter the text you want to find\nfor example: @TestRail")
	flag.StringVar(&directory, "dir", "", "input directory for parsing\nfor example: /home/users/")
	flag.Parse()
	if directory == "" {
		printNotFoundFlag("dir")
	}
	if searchText == "" {
		printNotFoundFlag("text")
	}
	return Config{Extension: fmt.Sprintf(".%s", extension), Directory: directory, SearchText: searchText}
}

func printNotFoundFlag(flagName string) {
	fmt.Printf("cannot find flag %s : use flag '-help' for see the flags", flagName)
	os.Exit(1)
}
