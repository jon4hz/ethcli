package tui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jon4hz/ethcli/internal/config"
	"github.com/jon4hz/ethcli/internal/ethcli"
	"github.com/jon4hz/ethcli/internal/tui/module"
	"github.com/jon4hz/ethcli/internal/tui/module/newwallet"
	"github.com/jon4hz/ethcli/internal/tui/style"
)

type state int

const (
	stateInit state = iota
	stateLoadHDWallet
	stateLoadKeystore
	stateShowPublicKey
	stateShowPrivateKey
	stateShowMnemonic
	stateShowBalance
	stateShowTokenBalance
	stateNewMessage
	stateNewTx
	stateNewTokenTx
	stateKSStore
	stateSetRPC
	stateReady
	stateQuit
)

type model struct {
	state        state
	wallet       *ethcli.Wallet
	config       *config.Config
	currentModel tea.Model
	list         list.Model
	header       string
	footer       string
}

func Start() error {
	del := list.NewDefaultDelegate()
	del.Styles.SelectedDesc.Foreground(style.FocusedStyle.GetForeground()).BorderForeground(style.FocusedStyle.GetForeground())
	del.Styles.SelectedTitle.Foreground(style.FocusedStyle.GetForeground()).BorderForeground(style.FocusedStyle.GetForeground())
	m := model{
		list:   list.NewModel(getInitialItems(), del, 0, 0),
		wallet: new(ethcli.Wallet),
	}
	m.list.Styles.FilterCursor = style.FocusedStyle
	m.list.SetFilteringEnabled(false)
	m.list.SetShowStatusBar(false)
	m.list.Title = "ethcli"
	m.list.Styles.Title = style.TitleStyle

	m.config = config.Get()

	return tea.NewProgram(m, tea.WithAltScreen()).Start()
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			switch m.state {
			case stateInit, stateReady:
				m.state = m.list.SelectedItem().(MenuItem).state
				cmd := m.list.SelectedItem().(MenuItem).module.Init()
				return m, cmd
			}
		}

	case tea.WindowSizeMsg:
		top, right, bottom, left := style.DocStyle.GetMargin()
		m.list.SetSize(msg.Width-left-right, msg.Height-top-bottom)

	case newwallet.Msg:
		*m.wallet = ethcli.Wallet(msg)
		m.list.SetItems(m.getMainItems())
		m.list.Title = m.wallet.Address()

	case module.BackMsg:
		m.state = stateReady
	}

	var cmd tea.Cmd
	switch m.state {
	case stateInit, stateReady:
		m.list, cmd = m.list.Update(msg)
	default:
		cmd = m.list.SelectedItem().(MenuItem).module.Update(msg)
	}
	return m, cmd
}

func (m model) View() string {
	switch m.state {
	case stateInit, stateReady:
		return style.DocStyle.Render(m.list.View())
	default:
		return style.ModuleWrapper.Render(m.selectedView())
	}
}

func (m model) selectedView() string {
	return m.list.SelectedItem().(MenuItem).module.View()
}
