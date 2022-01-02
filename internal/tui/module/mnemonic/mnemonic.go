package mnemonic

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jon4hz/ethcli/internal/ethcli"
	"github.com/jon4hz/ethcli/internal/tui/bubbles/simpleview"
	"github.com/jon4hz/ethcli/internal/tui/module"
	"github.com/jon4hz/ethcli/internal/tui/style"
)

type model struct {
	wallet *ethcli.Wallet
	view   *simpleview.Model
}

func NewModel(wallet *ethcli.Wallet) module.Module {
	return &model{
		wallet: wallet,
		view:   simpleview.NewModel("", wallet.Address(), style.BlurredStyle.Render("Press any key to continue")),
	}
}

func (m *model) Init() tea.Cmd {
	return m.view.Init()
}

func (m *model) Update(msg tea.Msg) tea.Cmd {
	return m.view.Update(msg)
}

func (m *model) View() string {
	var content string
	if m.wallet.Mnemonic() != "" {
		content = m.wallet.Mnemonic()
	} else {
		content = style.BlurredStyle.Render("no mnemonic found")
	}
	m.view.SetContent(content)
	return m.view.View()
}
