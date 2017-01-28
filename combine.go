package combine

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// dirname以下にあるファイルをfileNameファイルにまとめる
func Combine(dirname string, fileName string) {
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fmt.Println(file.Name())

		file, err := os.Open(fmt.Sprintf("%s/%s", dirname, file.Name()))
		if err != nil {
			log.Fatal(err)
		}
		b, err := ioutil.ReadAll(file)
		appendToFile(fileName, string(b))
		file.Close()
	}
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
