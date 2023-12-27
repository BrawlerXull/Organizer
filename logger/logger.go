package logger

import (
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

func SetupLogger() *os.File {
	logFile, err := os.OpenFile("logfile.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Unable to open log file:", err)
	}
	log.SetOutput(logFile)
	return logFile
}

func Out() *log.Logger {
	Stdlog := log.New(os.Stderr)
	log.ErrorLevelStyle = lipgloss.NewStyle().
		SetString("ERROR!!").
		Bold(true).
		Padding(0, 1, 0, 1).
		Background(lipgloss.AdaptiveColor{
			Light: "203",
			Dark:  "204",
		}).
		Foreground(lipgloss.Color("#211f26"))
	return Stdlog
}

func CloseLogger(logFile *os.File) {
	err := logFile.Close()
	if err != nil {
		log.Fatal("Error closing log file:", err)
	}
}
