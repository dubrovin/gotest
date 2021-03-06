package utils

import (
	"archive/zip"
	"fmt"
	"log"
	"os"
)

func CreateTestFile(dir, fileName string) {
	// Create a buffer to write our archive to.
	os.Mkdir(dir, os.ModePerm)
	f, err := os.Create(fmt.Sprintf("%s/%s", dir, fileName))
	if err != nil {
		log.Fatal(err)
	}

	// Create a new zip archive.
	w := zip.NewWriter(f)

	// Add some files to the archive.
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling licence.\nWrite more examples."},
	}
	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}

	err = w.Flush()
	if err != nil {
		log.Fatal(err)
	}
	// Make sure to check the error on Close.
	err = w.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteDir(dir string) {
	os.RemoveAll(dir)
}
