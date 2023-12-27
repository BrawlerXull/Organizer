package styles

import "github.com/charmbracelet/lipgloss"

var WelcomeStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#81efc5")).
	Padding(0, 1, 0, 1).
	BorderBottom(true).
	BorderTop(true).
	BorderStyle(lipgloss.NormalBorder()).
	BorderBottomForeground(lipgloss.Color("#3c4056")).
	BorderTopForeground(lipgloss.Color("#3c4056"))

var InputPromptStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#6aaa96"))

var ArrStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#ba9af8"))

var BorderNotif = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderBottom(true).
	BorderTop(true).
	BorderForeground(lipgloss.Color("#3c4056")).
	Padding(0, 1, 0, 1)

var InfoStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#f19da5"))

var QuestionStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#5477a8")).
	Padding(0, 1, 0, 1).
	BorderBottom(true).
	BorderTop(true)
