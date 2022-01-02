package tui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/jon4hz/ethcli/internal/config"
	"github.com/jon4hz/ethcli/internal/tui/module/mnemonic"
	"github.com/jon4hz/ethcli/internal/tui/module/newwallet"
	"github.com/jon4hz/ethcli/internal/tui/module/qr"
	"github.com/jon4hz/ethcli/internal/tui/module/quit"
	"github.com/jon4hz/ethcli/internal/tui/module/rpc"
)

func getInitialItems() []list.Item {
	return []list.Item{
		newMenuItem("New Wallet", "Create a new wallet", withState(stateReady),
			withModel(newwallet.NewModel()),
		),
		newMenuItem("Load Mnmemonic", "Load a wallet with a mnemonic seedphrase", withState(stateLoadHDWallet)),
		newMenuItem("Load Keystore", "Load a wallet from a keystore file", withState(stateLoadKeystore)),
	}
}

func (m model) getMainItems() []list.Item {
	return []list.Item{
		newMenuItem("Public Key", "View your public key", withState(stateShowPublicKey),
			withModel(qr.NewModel(m.wallet.Address(), m.wallet)),
		),
		newMenuItem("Private Key", "View your private key", withState(stateShowPrivateKey),
			withModel(qr.NewModel(m.wallet.PrivateKeyString(), m.wallet)),
		),
		newMenuItem("Show Mnemonic", "Show your mnemonic seedphrase", withState(stateShowMnemonic),
			withModel(mnemonic.NewModel(m.wallet)),
		),
		newMenuItem("Balance", "View your balance", withState(stateShowBalance)),
		newMenuItem("Token Balance", "View your balance of a particular token", withState(stateShowTokenBalance)),
		newMenuItem("New Transaction", "Create a new transaction", withState(stateNewTx)),
		newMenuItem("New Token Transfer", "Create a new token transfer", withState(stateNewTokenTx)),
		newMenuItem("Sign Message", "Sign a message", withState(stateNewMessage)),
		newMenuItem("Set RPC", "Set the RPC URL", withState(stateSetRPC),
			withModel(rpc.NewModel(config.Get())),
		),
		newMenuItem("Save Keystore", "Save the wallet to a keystore file", withState(stateKSStore)),
		newMenuItem("Quit", "Quit the application", withState(stateQuit), withModel(quit.NewModel())),
	}
}
