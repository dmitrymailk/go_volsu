package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	// Расширение файлов для поиска
	ext := ".txt"

	// Корневая папка для поиска
	root := "/app/task_6/test"

	// Рекурсивный обход дерева каталогов
	walkFn := func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && filepath.Ext(path) == ext {
			fmt.Println(path)
		}
		return nil
	}

	// Запускаем обход
	filepath.Walk(root, walkFn)

}
