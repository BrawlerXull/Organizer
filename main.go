package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/logrusorgru/aurora/v4"
)

func setupLogger() *os.File {
	logFile, err := os.OpenFile("logfile.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Unable to open log file:", err)
	}
	log.SetOutput(logFile)
	return logFile
}

func closeLogger(logFile *os.File) {
	if err := logFile.Close(); err != nil {
		log.Fatal("Error closing log file:", err)
	}
}

func main() {
	logFile := setupLogger()
	defer closeLogger(logFile)
	fmt.Println(aurora.Magenta("Welcome to Ogranizer"))
	fmt.Println(aurora.Cyan("Enter the directory to clean"))

	in := bufio.NewReader(os.Stdin)

	filePathDir, err := in.ReadString('\n')
	filePathDir = strings.TrimSuffix(filePathDir, "\n")
	filePathDir = strings.TrimSuffix(filePathDir, "\r")

	files, err := os.ReadDir(filePathDir)
	if err != nil {
		panic(err)
	}
	fmt.Println(aurora.Cyan("Below files are found"))

	for _, file := range files {
		fmt.Println(aurora.Yellow(file.Name()))
		extensionName := filepath.Ext(file.Name())
		extensionName = strings.TrimPrefix(extensionName, ".")

		if file.Name() == ".DS_Store" || extensionName == "" {
			continue
		}
		destinationDir := filepath.Join(filePathDir, extensionName)
		err := os.MkdirAll(destinationDir, 0755)
		if err != nil {
			fmt.Println("Error creating destination directory:", err)
			log.Fatal("Error creating destination directory:", err)
			return
		}

		sourceFilePath := filepath.Join(filePathDir, file.Name())
		destinationFilePath := filepath.Join(destinationDir, file.Name())
		err = os.Rename(sourceFilePath, destinationFilePath)
		if err != nil {
			fmt.Println("Error moving the file:", err)
			log.Fatal("Error moving the file:", err)
			return
		}
	}
	log.Info(filePathDir)

	fmt.Println(aurora.Magenta("Press Enter key to close"))
	Close, err := in.ReadString('\n')
	fmt.Println(Close)

}
