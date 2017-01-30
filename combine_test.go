package combine

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"testing"
)

func TestCombine(t *testing.T) {
	// create dir and files
	// /tmp/combinexxxxxxxxx/files/sample0~9
	dir, err := ioutil.TempDir("", "combine")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir) // clean up

	filesDir, err := ioutil.TempDir(dir, "files")
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 10; i++ {
		file := fmt.Sprintf("%s/%s%d", filesDir, "/sample", i)
		err = writeFile(file, strconv.Itoa(i))
		if err != nil {
			t.Fatal(err)
		}
	}

	// combine the content of /tmp/combinexxxxxxxxx/files/sample0~9 files to /tmp/combinexxxxxxxxx/all.txt
	filename := dir + "/" + "all.txt"

	err = Combine(filesDir, filename)
	if err != nil {
		t.Fatal(err)
	}

	// check
	data := "0123456789"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("ReadFile %s: %v", filename, err)
	}

	if string(contents) != data {
		t.Fatalf("contents = %q\nexpected = %q", string(contents), data)
	}
}

func writeFile(filename string, data string) error {
	err := ioutil.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		return err
	}
	return err
}
