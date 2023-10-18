package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func walkDir(dir string, firstLetter string) {

	files, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if file.IsDir() {
			subdir := filepath.Join(dir, file.Name())
			walkDir(subdir, firstLetter) // рекурсивный вызов
		} else {
			if strings.HasPrefix(file.Name(), firstLetter) {
				fmt.Println(file.Name())
			}
		}
	}
}

func main() {

	// Буква, с которой должны начинаться имена файлов
	// firstLetter := "1"
	firstLetter := "q"

	// Путь к папке
	folder := "/app/task_6/test"

	walkDir(folder, firstLetter)

}
