package qr

import (
	tea "github.com/charmbracelet/bubbletea"
	qrc "github.com/jon4hz/ethcli/internal/qr"
	module "github.com/jon4hz/ethcli/internal/tui/modules"
)

type model struct {
	key string
}

func NewModel(key string) module.Module {
	return &model{
		key: key,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) tea.Cmd {
	switch msg.(type) {
	case tea.KeyMsg:
		return func() tea.Msg { return module.BackMsg{} }
	}
	return nil
}

func (m model) View() string {
	// var s strings.Builder

	return qrc.NewQr(m.key)
}
