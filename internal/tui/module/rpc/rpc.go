package rpc

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jon4hz/ethcli/internal/config"
	"github.com/jon4hz/ethcli/internal/ethcli"
	"github.com/jon4hz/ethcli/internal/tui/bubbles/simpleview"
	"github.com/jon4hz/ethcli/internal/tui/module"
	"github.com/jon4hz/ethcli/internal/tui/style"
)

type (
	Msg    string
	errMsg error

	keyMap struct {
		Confirm key.Binding
		Back    key.Binding
		Quit    key.Binding
	}
)

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Confirm, k.Back, k.Quit}
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// key.Map interface.
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		k.ShortHelp(),
	}
}

var keys = keyMap{
	Confirm: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "confirm"),
	),
	Back: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "back"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "quit"),
	),
}

type state int

const (
	stateReady state = iota
	stateChecking
	stateDone
)

type Model struct {
	state   state
	url     string
	input   textinput.Model
	spinner spinner.Model
	view    *simpleview.Model
	errs    string
}

func NewModel(cfg *config.Config) *Model {
	placeholder := "none"
	if cfg.RPC != "" {
		placeholder = cfg.RPC
	}

	i := textinput.NewModel()
	i.CursorStyle = style.FocusedStyle
	i.Placeholder = placeholder
	i.Focus()
	i.SetValue(placeholder)

	s := spinner.NewModel()
	s.Spinner = spinner.Points
	s.Style = style.FocusedStyle

	help := help.NewModel()

	return &Model{
		url:     placeholder,
		input:   i,
		spinner: s,
		view:    simpleview.NewModel("", "", help.View(keys)),
	}
}

func (m *Model) Init() tea.Cmd {
	return tea.Batch(
		textinput.Blink,
		m.view.Init(),
	)
}

func (m *Model) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.state {
		case stateReady:
			switch {
			case key.Matches(msg, keys.Quit):
				return tea.Quit
			case key.Matches(msg, keys.Back):
				m.state = stateReady
				return module.Back
			case key.Matches(msg, keys.Confirm):
				m.state = stateChecking
				m.errs = ""
				return tea.Batch(
					spinner.Tick,
					m.checkRPC(),
				)
			}

		case stateChecking:
			switch {
			case key.Matches(msg, keys.Quit):
				return tea.Quit
			case key.Matches(msg, keys.Back):
				m.state = stateReady
				return module.Back
			}

		case stateDone:
			switch {
			case key.Matches(msg, keys.Quit):
				return tea.Quit
			case key.Matches(msg, keys.Back):
				return module.Back
			}
		}

	case tea.WindowSizeMsg:
		_, right, _, left := style.ModuleWrapper.GetMargin()
		m.input.Width = msg.Width - left - right
		return m.view.Update(msg)

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return cmd

	case errMsg:
		m.errs = msg.Error()
		m.input.Reset()
		m.state = stateReady
		return nil
	}

	switch m.state {
	case stateReady:
		var cmd tea.Cmd
		m.input, cmd = m.input.Update(msg)
		if cmd != nil {
			return cmd
		}
	}
	return nil
}

func (m *Model) View() string {
	switch m.state {
	case stateReady:
		var s strings.Builder
		s.WriteString(m.input.View())
		if m.errs != "" {
			s.WriteString(fmt.Sprintf("\n\n%s", style.ErrStyle.Render(m.errs)))
		}
		m.view.SetContent(s.String())
	case stateChecking:
		m.view.SetContent(m.spinner.View())
	}
	return m.view.View()
}

func (m *Model) checkRPC() tea.Cmd {
	return func() tea.Msg {
		url := strings.TrimSpace(m.input.Value())
		client, err := ethcli.NewClient(url)
		if err != nil {
			// TODO log
			return errMsg(fmt.Errorf("invalid rpc: %s", url))
		}
		chainID, err := client.ChainID()
		if err != nil {
			// TODO log
			return errMsg(fmt.Errorf("invalid rpc: %s", url))
		}
		_ = chainID
		return ""
	}
}
