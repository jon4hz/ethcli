package simpleview

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/jon4hz/ethcli/internal/tui/module"
	"github.com/jon4hz/ethcli/internal/tui/style"
	"golang.org/x/term"
)

var (
	titleStyle  = style.TitleStyle.Copy().MarginBottom(1)
	footerStyle = style.FooterStyle.Copy().MarginTop(1)
)

type Model struct {
	content        string
	width, height  int
	header, footer string
	minWidth       int
}

func NewModel(content, header, footer string) *Model {
	return &Model{
		content: content,
		header:  header,
		footer:  footer,
	}
}

func (m *Model) SetContent(content string) {
	m.content = content
}

func (m *Model) SetHeader(header string) {
	m.header = header
}

func (m *Model) SetFooter(footer string) {
	m.footer = footer
}

func (m *Model) SetMinWidth(minWidth int) {
	m.minWidth = minWidth
}

func (m *Model) Init() tea.Cmd {
	top, right, bottom, left := style.ModuleWrapper.GetMargin()
	availWidth, availHeight, _ := term.GetSize(int(os.Stdout.Fd()))
	availHeight -= top + bottom
	availWidth -= right + left
	m.width = availWidth
	m.height = availHeight
	return nil
}

func (m *Model) Update(msg tea.Msg) tea.Cmd {
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

func (m Model) View() string {
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

	width := m.width
	if width < m.minWidth {
		width = m.minWidth
	}

	content := lipgloss.NewStyle().Width(width).Height(availHeight).Render(m.content)
	sections = append(sections, content)

	if m.footer != "" {
		sections = append(sections, footer)
	}
	return lipgloss.JoinVertical(lipgloss.Left, sections...)
}
