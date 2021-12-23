package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	module "github.com/jon4hz/ethcli/internal/tui/modules"
)

type (
	menuListOption func(*MenuItem)

	MenuItem struct {
		title     string
		desc      string
		nextState state
		callback  func(tea.Msg) tea.Cmd
		module    module.Module
	}
)

var (
	defaultCallback = func(tea.Msg) tea.Cmd { return nil }
	quitCallback    = func(tea.Msg) tea.Cmd { return tea.Quit }
)

func (i MenuItem) Title() string       { return i.title }
func (i MenuItem) Description() string { return i.desc }
func (i MenuItem) FilterValue() string { return "" }

func (i *MenuItem) SetModel(module module.Module) {
	i.module = module
}

func newMenuItem(title, desc string, opts ...menuListOption) MenuItem {
	m := MenuItem{
		title:    title,
		desc:     desc,
		callback: defaultCallback,
		module:   &module.DefaultModule{},
	}
	for _, opt := range opts {
		opt(&m)
	}
	return m
}

func withNextState(state state) menuListOption {
	return func(m *MenuItem) {
		m.nextState = state
	}
}

func withCallback(callback func(tea.Msg) tea.Cmd) menuListOption {
	return func(m *MenuItem) {
		m.callback = callback
	}
}

func withModel(module module.Module) menuListOption {
	return func(m *MenuItem) {
		m.module = module
	}
}
