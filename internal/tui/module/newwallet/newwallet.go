package newwallet

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jon4hz/ethcli/internal/ethcli"
	module "github.com/jon4hz/ethcli/internal/tui/module"
)

type Msg ethcli.Wallet

type model struct{}

func NewModel() module.Module {
	return &model{}
}

func (m *model) Init() tea.Cmd {
	return func() tea.Msg {
		return Msg(ethcli.NewWallet())
	}
}

func (m *model) Update(msg tea.Msg) tea.Cmd {
	return nil
}

func (m *model) View() string {
	return ""
}
