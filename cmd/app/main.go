package main

import (
	"copypaste_alerter/internal/config"
	w "copypaste_alerter/internal/wrapper"
	"fmt"
)

func main() {
	var conf config.Config
	conf = conf.Init()
	var wrapper = w.FileWrapper{Config: conf}
	res, err := wrapper.GetParsingResult()
	if err != nil {
		fmt.Println(err)
	}
	res.Print()
}
