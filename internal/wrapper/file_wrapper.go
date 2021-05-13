package wrapper

import (
	"bufio"
	"copypaste_alerter/internal/config"
	"copypaste_alerter/internal/models"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
)

type FileWrapper struct {
	Result models.ParsingResult
	Config config.Config
	Files  []models.FileEntity
}

var preparedFiles = make([]models.FileEntity, 0)

func (fw FileWrapper) FindFiles(currentDir string, extension string) {
	files, err := ioutil.ReadDir(currentDir)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		filePath := path.Join(currentDir, f.Name())
		if f.IsDir() {
			fw.FindFiles(filePath, extension)
		} else {
			if strings.Contains(f.Name(), string(extension)) {
				preparedFiles = append(preparedFiles, models.FileEntity{Filename: f.Name(), Path: filePath})
			}
		}
	}
}

func (fw FileWrapper) ReadFile(wg *sync.WaitGroup, filePath string, res *models.ParsingResult, searchText string) {
	defer wg.Done()
	f, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(f)
	var lineNumber = 1
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " ")
		if strings.Contains(line, searchText) {
			fmt.Printf("finded line with %s: %s, line number : %v \n", searchText, line, lineNumber)
			res.AddFindedValue(scanner.Text(), fmt.Sprintf("file: %s  line: %v", filePath, lineNumber))
		}
		lineNumber++
	}

	if closeErr := f.Close(); closeErr != nil {
		fmt.Println(closeErr)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func (fw FileWrapper) GetParsingResult() (models.ParsingResult, error) {
	wg := sync.WaitGroup{}
	var res = models.ParsingResult{LineMap: map[string][]string{}, Mx: &sync.Mutex{}}
	basePath, _ := filepath.Abs(fw.Config.Directory)
	fw.FindFiles(basePath, fw.Config.Extension)
	for _, file := range preparedFiles {
		wg.Add(1)
		go fw.ReadFile(&wg, file.Path, &res, fw.Config.SearchText)
	}
	wg.Wait()
	fw.Result = res
	return fw.Result, nil
}
