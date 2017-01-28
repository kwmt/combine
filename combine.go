package combine

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/kwmt/go-utils"
)

// combine the files under dirName directory
func Combine(dirName string, fileName string) {
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}
	total := len(files)

	for i, file := range files {
		file, err := os.Open(fmt.Sprintf("%s/%s", dirName, file.Name()))
		if err != nil {
			log.Fatal(err)
		}
		b, err := ioutil.ReadAll(file)
		appendToFile(fileName, string(b))

		fmt.Printf("\r%3.0f%%", utils.Percent(i+1, total))
		file.Close()
	}
	fmt.Println()
}

// append text to filName file
func appendToFile(fileName string, text string) {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err = f.WriteString(text); err != nil {
		panic(err)
	}
}
