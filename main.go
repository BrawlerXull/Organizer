package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func main() {
	currentUser, err := user.Current()
	if err != nil {
		panic(err)
	}

	filePathDir := filepath.Join(currentUser.HomeDir, "Desktop/wdfvs")

	files, err := os.ReadDir(filePathDir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())

		extensionName := filepath.Ext(file.Name())
		extensionName = strings.TrimPrefix(extensionName, ".")
		fmt.Println(extensionName)

		if file.Name() == ".DS_Store" || extensionName == "" {
			continue
		}
		destinationDir := filepath.Join(filePathDir, extensionName)
		err := os.MkdirAll(destinationDir, 0755)
		if err != nil {
			fmt.Println("Error creating destination directory:", err)
			return
		}

		sourceFilePath := filepath.Join(filePathDir, file.Name())
		destinationFilePath := filepath.Join(destinationDir, file.Name())
		err = os.Rename(sourceFilePath, destinationFilePath)
		if err != nil {
			fmt.Println("Error moving the file:", err)
			return
		}
	}
}
