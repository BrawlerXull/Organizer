package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/filepicker"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type keyMap struct {
	Sort  key.Binding
	Clean key.Binding
	Help  key.Binding
	Quit  key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Sort, k.Clean, k.Help, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Sort, k.Clean},
		{k.Help, k.Quit},
	}
}

var keys = keyMap{
	Sort: key.NewBinding(
		key.WithKeys("s"),
		key.WithHelp("s", "Sort by extensions"),
	),
	Clean: key.NewBinding(
		key.WithKeys("c"),
		key.WithHelp("c", "Clean old files"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}

type model struct {
	filepicker   filepicker.Model
	selectedFile string
	quitting     bool
	err          error
	keys         keyMap
	help         help.Model
	DateInput    textinput.Model
	changeview   bool
	currentview  string
}

type clearErrorMsg struct{}

func clearErrorAfter(t time.Duration) tea.Cmd {
	return tea.Tick(t, func(_ time.Time) tea.Msg {
		return clearErrorMsg{}
	})
}

func (m model) Init() tea.Cmd {
	return m.filepicker.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		case "s":
			m.quitting = true
			Sortbyext(m.filepicker.CurrentDirectory)
			return m, tea.Quit
		case "c":
			m.changeview = true
			m.currentview = "cleaninp"
			m.DateInput.Focus()
			return m, tea.HideCursor
		case tea.KeyEnter.String():
			if m.currentview == "cleaninp" {
				m.quitting = true
				Clean(m.filepicker.CurrentDirectory, m.DateInput.Value())
				return m, tea.Quit
			}
		}

	case clearErrorMsg:
		m.err = nil
	}

	var cmd tea.Cmd
	m.filepicker, cmd = m.filepicker.Update(msg)
	if m.currentview == "cleaninp" {
		m.DateInput, cmd = m.DateInput.Update(msg)
	}
	// Did the user select a file?
	if didSelect, path := m.filepicker.DidSelectFile(msg); didSelect {
		// Get the path of the selected file.
		m.selectedFile = path
		// clean(m.filepicker.CurrentDirectory)
	}

	// Did the user select a disabled file?
	// This is only necessary to display an error to the user.
	if didSelect, path := m.filepicker.DidSelectDisabledFile(msg); didSelect {
		// Let's clear the selectedFile and display an error.
		m.err = errors.New(path + " is not valid.")
		m.selectedFile = ""
		return m, tea.Batch(cmd, clearErrorAfter(2*time.Second))
	}

	return m, cmd
}

func (m model) View() string {
	var s strings.Builder
	s.WriteString("\n  ")
	if m.quitting {
		return ""
	} else if m.changeview {
		s.WriteString("\n\n" + m.DateInput.View() + "\n" + m.help.View(m.keys))
		return s.String()
	}
	if m.err != nil {
		s.WriteString(m.filepicker.Styles.DisabledFile.Render(m.err.Error()))
	} else if m.selectedFile == "" {
		s.WriteString("Current directory: " + m.filepicker.CurrentDirectory)
	} else {
		s.WriteString("Selected directory: " + m.filepicker.Styles.Selected.Render(m.filepicker.CurrentDirectory))
	}
	s.WriteString("\n\n" + m.filepicker.View() + "\n" + m.help.View(m.keys))
	return s.String()
}

func main() {
	tea.LogToFile("debug.log", "debug")
	fp := filepicker.New()
	fp.AllowedTypes = []string{".mod", ".sum", ".go", ".txt", ".md"}
	fp.CurrentDirectory, _ = os.UserHomeDir()

	ti := textinput.New()
	ti.Placeholder = "01/02/2023"
	ti.CharLimit = 10
	ti.Width = 20
	m := model{
		filepicker: fp,
		keys:       keys,
		help:       help.New(),
		DateInput:  ti,
	}
	tm, _ := tea.NewProgram(&m, tea.WithOutput(os.Stderr)).Run()
	mm := tm.(model)
	fmt.Println("\n Current directory: " + m.filepicker.Styles.Selected.Render(mm.filepicker.CurrentDirectory) + "\n")
}

func Sortbyext(dir string) {

	files, err := os.ReadDir(dir)
	checkerr(err)
	fmt.Println("Below files are found")

	for _, file := range files {
		// fmt.Println(file.Name())
		extensionName := filepath.Ext(file.Name())
		extensionName = strings.TrimPrefix(extensionName, ".")
		// fmt.Println(extensionName)

		if file.Name() == ".DS_Store" || extensionName == "" {
			continue
		}
		destinationDir := filepath.Join(dir, extensionName)
		err := os.MkdirAll(destinationDir, 0755)
		if err != nil {
			fmt.Println("Error creating destination directory:", err)
			return
		}

		sourceFilePath := filepath.Join(dir, file.Name())
		destinationFilePath := filepath.Join(destinationDir, file.Name())
		err = os.Rename(sourceFilePath, destinationFilePath)
		if err != nil {
			fmt.Println("Error moving the file:", err)
			return
		}
	}

}

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}

func Clean(filePathDir string, d string) {
	files, err := os.ReadDir(filePathDir)
	checkerr(err)
	day, err := strconv.Atoi(strings.Split(d, "/")[0])
	checkerr(err)
	month, err := strconv.Atoi(strings.Split(d, "/")[1])
	checkerr(err)
	year, err := strconv.Atoi(strings.Split(d, "/")[2])
	checkerr(err)
	for _, file := range files {
		cleanFileDir := filePathDir + "/" + file.Name()
		fileInfo, err := os.Stat(cleanFileDir)
		if err != nil {
			panic(err)
		}
		onDayAgo := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
		if fileInfo.ModTime().Before(onDayAgo) {
			err := os.RemoveAll(cleanFileDir)
			if err != nil {
				panic(err)
			}
		}
	}
}
