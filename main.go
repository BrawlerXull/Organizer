package main

import (
	"bufio"
	"fmt"
	cleanold "organizer/cleanOld"
	"organizer/logger"
	"organizer/organizer"
	"organizer/styles"
	"os"
	"strings"

	"github.com/charmbracelet/log"
)

func main() {

	in := bufio.NewReader(os.Stdin)

	Stdlog := logger.Out()
	logFile := logger.SetupLogger()
	defer logger.CloseLogger(logFile)
	fmt.Println(styles.WelcomeStyle.Render("Welcome to Organizer !!"))

	fmt.Println(styles.QuestionStyle.Render("What do you want to perform ? (cleanOld / organize)"))
	fmt.Print(styles.ArrStyle.Render("> "))
	userResp, err := in.ReadString('\n')
	checkerr(err, Stdlog)
	userResp = strings.TrimSuffix(userResp, "\n")
	userResp = strings.TrimSuffix(userResp, "\r")

	if userResp == "organize" {
		organizer.Organizer()

	} else if userResp == "cleanOld" {
		cleanold.CleanOld()
	} else {
		customError(styles.InfoStyle.Render("Command not found"), Stdlog)
	}

	fmt.Println(styles.InputPromptStyle.Render("Press Enter key to close"))
	_, err = in.ReadString('\n')
	checkerr(err, Stdlog)
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
