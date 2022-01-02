package qr

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/jon4hz/ethcli/internal/ethcli"
	qrc "github.com/jon4hz/ethcli/internal/qr"
	"github.com/jon4hz/ethcli/internal/tui/bubbles/simpleview"
	"github.com/jon4hz/ethcli/internal/tui/module"
	"github.com/jon4hz/ethcli/internal/tui/style"
)

type model struct {
	content string
	wallet  *ethcli.Wallet
	view    *simpleview.Model
}

func NewModel(content string, wallet *ethcli.Wallet) module.Module {
	s := simpleview.NewModel("", wallet.Address(), style.BlurredStyle.Render("Press any key to continue"))
	s.SetMinWidth(46)
	return &model{
		wallet:  wallet,
		view:    s,
		content: content,
	}
}

func (m *model) Init() tea.Cmd {
	return m.view.Init()
}

func (m *model) Update(msg tea.Msg) tea.Cmd {
	return m.view.Update(msg)
}

func (m *model) View() string {
	var s strings.Builder
	s.WriteString(qrc.NewQr(m.content))
	s.WriteString("\n")
	s.WriteString(m.content)
	m.view.SetContent(s.String())
	return m.view.View()
}
