package combine

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kwmt/go-utils"
)

// combine the files under dirName directory
func Combine(dirName string, fileName string) error {
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		return err
	}
	total := len(files)

	for i, file := range files {
		file, err := os.Open(fmt.Sprintf("%s/%s", dirName, file.Name()))
		if err != nil {
			return err
		}
		b, err := ioutil.ReadAll(file)
		if err != nil {
			return err
		}
		err = appendToFile(fileName, string(b))
		if err != nil {
			return err
		}

		fmt.Printf("\r%3.0f%%", utils.Percent(i+1, total))
		file.Close()
	}
	fmt.Println()
	return nil
}

// append text to filName file
func appendToFile(fileName string, text string) error {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err = f.WriteString(text); err != nil {
		return err
	}
	return nil
}
