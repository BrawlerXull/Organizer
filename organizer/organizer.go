package organizer

import (
	"bufio"
	"fmt"
	"organizer/logger"
	"organizer/styles"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/log"
)

func Organizer() {
	fmt.Println(styles.InputPromptStyle.Render("Enter the directory to clean"))
	fmt.Print(styles.ArrStyle.Render("> "))
	
	in := bufio.NewReader(os.Stdin)

	Stdlog := logger.Out()
	logFile := logger.SetupLogger()
	defer logger.CloseLogger(logFile)

	filePathDir, err := in.ReadString('\n')
	checkerr(err, Stdlog)
	filePathDir = strings.TrimSuffix(filePathDir, "\n")
	filePathDir = strings.TrimSuffix(filePathDir, "\r")

	files, err := os.ReadDir(filePathDir)
	checkerr(err, Stdlog)
	fmt.Println(styles.BorderNotif.Render("Below files are found"))

	for _, file := range files {
		fmt.Println(styles.InfoStyle.Render(fmt.Sprintf("* %s", file.Name())))
		extensionName := filepath.Ext(file.Name())
		extensionName = strings.TrimPrefix(extensionName, ".")

		if file.Name() == ".DS_Store" || extensionName == "" {
			continue
		}
		destinationDir := filepath.Join(filePathDir, extensionName)
		err := os.MkdirAll(destinationDir, 0755)
		checkerr(err, Stdlog)

		sourceFilePath := filepath.Join(filePathDir, file.Name())
		destinationFilePath := filepath.Join(destinationDir, file.Name())
		err = os.Rename(sourceFilePath, destinationFilePath)
		checkerr(err, Stdlog)

		log.Info(filePathDir)
	}
}

func checkerr(err error, Stdlog *log.Logger) {
	if err != nil {
		Stdlog.Error(err)
		log.Fatal(err)
		return
	}
}

func customError(err string, Stdlog *log.Logger) {
	Stdlog.Error(err)
	log.Fatal(err)
}
