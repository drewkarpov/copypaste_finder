package models

import (
	"fmt"
	"sync"
)

type ParsingResult struct {
	Mx      *sync.Mutex
	LineMap map[string][]string
}

func (ps ParsingResult) AddFindedValue(text, filePath string) {
	ps.Mx.Lock()
	ps.LineMap[text] = append(ps.LineMap[text], filePath)
	ps.Mx.Unlock()
}

func (ps ParsingResult) Print() {
	var countOfFilesWithDuplicateText = 0
	for key, value := range ps.LineMap {
		if len(value) > 1 {
			fmt.Println("_______________________________\n")
			fmt.Printf("string '%s' was finded in files:\n", key)
			for _, v := range value {
				fmt.Printf("%v\n", v)
			}
			fmt.Println("_______________________________\n")
			countOfFilesWithDuplicateText++
		}
	}
	if countOfFilesWithDuplicateText == 0 {
		fmt.Println("not find files with duplicate text")
	}
}
