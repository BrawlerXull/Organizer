package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/logrusorgru/aurora/v4"
)

func main() {
	fmt.Println(aurora.Magenta("Welcome to Ogranizer"))
	fmt.Println(aurora.Cyan("Enter the directory to clean"))

	// currentUser, err := user.Current()
	// if err != nil {
	// 	panic(err)
	// }

	in := bufio.NewReader(os.Stdin)

	filePathDir, err := in.ReadString('\n')
	filePathDir = strings.TrimSuffix(filePathDir, "\n")
	filePathDir = strings.TrimSuffix(filePathDir, "\r")

	// filePathDir := filepath.Join(currentUser.HomeDir, "Desktop/kk")

	files, err := os.ReadDir(filePathDir)
	if err != nil {
		panic(err)
	}
	fmt.Println(aurora.Cyan("Below files are found"))

	for _, file := range files {
		fmt.Println(aurora.Yellow(file.Name()))
		extensionName := filepath.Ext(file.Name())
		extensionName = strings.TrimPrefix(extensionName, ".")
		// fmt.Println(extensionName)

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

	fmt.Println(aurora.Magenta("Press Enter key to close"))
	Close, err := in.ReadString('\n')
	fmt.Println(Close)

}
