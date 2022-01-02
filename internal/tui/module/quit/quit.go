package quit

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jon4hz/ethcli/internal/tui/module"
)

type model struct{}

func NewModel() module.Module {
	return &model{}
}

func (m *model) Init() tea.Cmd          { return tea.Quit }
func (m *model) Update(tea.Msg) tea.Cmd { return nil }
func (m *model) View() string           { return "" }
