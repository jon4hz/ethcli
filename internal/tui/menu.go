package tui

import (
	"github.com/jon4hz/ethcli/internal/tui/module"
)

type (
	menuListOption func(*MenuItem)

	MenuItem struct {
		title  string
		desc   string
		state  state
		module module.Module
	}
)

func (i MenuItem) Title() string       { return i.title }
func (i MenuItem) Description() string { return i.desc }
func (i MenuItem) FilterValue() string { return "" }

func (i *MenuItem) SetModel(module module.Module) {
	i.module = module
}

func newMenuItem(title, desc string, opts ...menuListOption) MenuItem {
	m := MenuItem{
		title:  title,
		desc:   desc,
		module: &module.DefaultModule{},
	}
	for _, opt := range opts {
		opt(&m)
	}
	return m
}

func withState(state state) menuListOption {
	return func(m *MenuItem) {
		m.state = state
	}
}

func withModel(module module.Module) menuListOption {
	return func(m *MenuItem) {
		m.module = module
	}
}
