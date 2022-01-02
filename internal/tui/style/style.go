package style

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	DocStyle            = lipgloss.NewStyle().Margin(1, 2)
	ModuleWrapper       = lipgloss.NewStyle().Margin(1, 4)
	FocusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("#79e7e7"))
	BlurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	CursorStyle         = FocusedStyle
	NoStyle             = lipgloss.NewStyle()
	HelpStyle           = BlurredStyle
	CursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))
	TitleStyle          = lipgloss.NewStyle().Foreground(lipgloss.Color("#14044d")).Background(lipgloss.Color("#79e7e7")).Bold(true)
	FooterStyle         = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	ErrStyle            = lipgloss.NewStyle().Foreground(lipgloss.Color("#f00")).Bold(true)

	FocusedButton = FocusedStyle.Render("[ Submit ]")
	BlurredButton = fmt.Sprintf("[ %s ]", BlurredStyle.Render("Submit"))
)
