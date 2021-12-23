package style

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	DocStyle            = lipgloss.NewStyle().Margin(1, 2)
	FocusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("#79e7e7"))
	BlurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Italic(true)
	CursorStyle         = FocusedStyle
	NoStyle             = lipgloss.NewStyle()
	HelpStyle           = BlurredStyle
	CursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))
	TitleStyle          = lipgloss.NewStyle().Foreground(lipgloss.Color("#14044d")).Background(lipgloss.Color("#79e7e7")).Bold(true)

	FocusedButton = FocusedStyle.Render("[ Submit ]")
	BlurredButton = fmt.Sprintf("[ %s ]", BlurredStyle.Render("Submit"))
)
