package combine

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/kwmt/go-utils"
)

// dirname以下にあるファイルをfileNameファイルにまとめる
func Combine(dirname string, fileName string) {
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}
	total := len(files)

	for i, file := range files {
		fmt.Printf("\r%3.0f%%", utils.Percent(i+1, total))

		file, err := os.Open(fmt.Sprintf("%s/%s", dirname, file.Name()))
		if err != nil {
			log.Fatal(err)
		}
		b, err := ioutil.ReadAll(file)
		appendToFile(fileName, string(b))
		file.Close()
	}
	fmt.Println()
}

// filnameファイルにtextをappendする
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
