package module

import tea "github.com/charmbracelet/bubbletea"

type BackMsg struct{}

type Module interface {
	Init() tea.Cmd
	Update(tea.Msg) tea.Cmd
	View() string
}

type DefaultModule struct{}

func (m DefaultModule) Init() tea.Cmd           { return nil }
func (m *DefaultModule) Update(tea.Msg) tea.Cmd { return func() tea.Msg { return BackMsg{} } }
func (m DefaultModule) View() string            { return "" }
