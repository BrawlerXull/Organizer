package cleanold

import (
	"bufio"
	"fmt"
	"organizer/logger"
	"organizer/styles"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/log"
)

func CleanOld() {
	in := bufio.NewReader(os.Stdin)

	Stdlog := logger.Out()
	logFile := logger.SetupLogger()
	defer logger.CloseLogger(logFile)
	fmt.Println(styles.InputPromptStyle.Render("Enter the directory to clean"))
	fmt.Print(styles.ArrStyle.Render("> "))

	filePathDir, err := in.ReadString('\n')
	checkerr(err, Stdlog)
	filePathDir = strings.TrimSuffix(filePathDir, "\n")
	filePathDir = strings.TrimSuffix(filePathDir, "\r")

	files, err := os.ReadDir(filePathDir)
	checkerr(err, Stdlog)

	fmt.Println(styles.InputPromptStyle.Render("Enter the number of years old files to clean"))
	fmt.Print(styles.ArrStyle.Render("> "))
	inputyear, err := in.ReadString('\n')
	checkerr(err, Stdlog)
	inputyear = strings.TrimSuffix(inputyear, "\n")
	inputyear = strings.TrimSuffix(inputyear, "\r")
	year, err := strconv.Atoi(inputyear)
	checkerr(err, Stdlog)

	fmt.Println(styles.InputPromptStyle.Render("Enter the number of months old files to clean"))
	fmt.Print(styles.ArrStyle.Render("> "))
	monthInput, err := in.ReadString('\n')
	checkerr(err, Stdlog)
	monthInput = strings.TrimSuffix(monthInput, "\n")
	monthInput = strings.TrimSuffix(monthInput, "\r")
	month, err := strconv.Atoi(monthInput)
	checkerr(err, Stdlog)

	fmt.Println(styles.InputPromptStyle.Render("Enter the number of days old files to clean"))
	fmt.Print(styles.ArrStyle.Render("> "))
	dayInput, err := in.ReadString('\n')
	checkerr(err, Stdlog)
	dayInput = strings.TrimSuffix(dayInput, "\n")
	dayInput = strings.TrimSuffix(dayInput, "\r")
	day, err := strconv.Atoi(dayInput)
	checkerr(err, Stdlog)

	fmt.Println(styles.BorderNotif.Render("Below files are found"))

	for _, file := range files {
		fmt.Println(styles.InfoStyle.Render(fmt.Sprintf("* %s", file.Name())))
		cleanFileDir := filePathDir + "/" + file.Name()
		fileInfo, err := os.Stat(cleanFileDir)
		checkerr(err, Stdlog)
		onDayAgo := time.Now().AddDate(-year, -month, -day)
		if fileInfo.ModTime().Before(onDayAgo) {
			err := os.RemoveAll(cleanFileDir)
			checkerr(err, Stdlog)
		}
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
