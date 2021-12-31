package simpleview

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	module "github.com/jon4hz/ethcli/internal/tui/modules"
	"github.com/jon4hz/ethcli/internal/tui/style"
)

var (
	titleStyle  = style.TitleStyle.Copy().MarginBottom(1)
	footerStyle = style.FooterStyle.Copy().MarginTop(1)
)

type model struct {
	content        string
	width, height  int
	header, footer string
}

func NewModel(content, header, footer string) module.Module {
	return &model{
		content: content,
		header:  header,
		footer:  footer,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return func() tea.Msg { return module.BackMsg{} }

	case tea.WindowSizeMsg:
		top, right, bottom, left := style.ModuleWrapper.GetMargin()
		m.height = msg.Height - top - bottom
		m.width = msg.Width - right - left
	}
	return nil
}

func (m model) View() string {
	var (
		sections    []string
		availHeight = m.height
	)

	if m.header != "" {
		header := titleStyle.Render(m.header)
		sections = append(sections, header)
		availHeight -= lipgloss.Height(header)
	}

	var footer string
	if m.footer != "" {
		footer = footerStyle.Render(m.footer)
		availHeight -= lipgloss.Height(footer)
	}

	content := lipgloss.NewStyle().Height(availHeight).Render(m.content)
	sections = append(sections, content)

	if m.footer != "" {
		sections = append(sections, footer)
	}
	return lipgloss.JoinVertical(lipgloss.Left, sections...)
}
