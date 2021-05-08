package config

import (
	"copypaste_alerter/internal/models"
	"flag"
	"fmt"
	"os"
)

type Config struct {
	Extension  models.Extension
	Directory  string
	SearchText string
}

func (c Config) Init() Config {
	var ext models.Extension
	var language string
	var directory string
	var searchText string

	flag.StringVar(&language, "lang", "", "input programming language name for parsing\nenable values: kotlin/python/go")
	flag.StringVar(&searchText, "text", "", "enter the text you want to find\nfor example: @TestRail")
	flag.StringVar(&directory, "dir", "", "input directory for parsing\nfor example: /home/users/")
	flag.Parse()

	if language == "" || searchText == "" || directory == "" {
		fmt.Println("use flag '-help' for see the flags")
		os.Exit(1)
	}
	ext = ext.GetByString(language)
	return Config{Extension: ext, Directory: directory, SearchText: searchText}
}
