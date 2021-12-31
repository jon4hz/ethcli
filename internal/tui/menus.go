package tui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jon4hz/ethcli/internal/ethcli"
	qrc "github.com/jon4hz/ethcli/internal/qr"
	"github.com/jon4hz/ethcli/internal/tui/modules/simpleview"
	"github.com/jon4hz/ethcli/internal/tui/style"
)

func getInitialItems() []list.Item {
	return []list.Item{
		newMenuItem("New Wallet", "Create a new wallet", withNextState(stateReady), withCallback(newWallet)),
		newMenuItem("Load Mnmemonic", "Load a wallet with a mnemonic seedphrase", withNextState(stateLoadHDWallet)),
		newMenuItem("Load Keystore", "Load a wallet from a keystore file", withNextState(stateLoadKeystore)),
	}
}

func newWallet(tea.Msg) tea.Cmd {
	return func() tea.Msg {
		return newWalletMsg(ethcli.New())
	}
}

func getMainItems(w *ethcli.Wallet) []list.Item {
	return []list.Item{
		newMenuItem("Public Key", "View your public key", withNextState(stateShowPublicKey),
			withModel(simpleview.NewModel(qrc.NewQr(w.PublicKeyString()), w.PublicKeyString(), style.BlurredStyle.Render("Press any key to continue"))),
		),
		newMenuItem("Private Key", "View your private key", withNextState(stateShowPrivateKey)),
		newMenuItem("Show Mnemonic", "Show your mnemonic seedphrase", withNextState(stateShowMnemonic),
			withModel(simpleview.NewModel(w.Mnemonic(), w.PublicKeyString(), style.BlurredStyle.Render("Press any key to continue"))),
		),
		newMenuItem("Balance", "View your balance", withNextState(stateShowBalance)),
		newMenuItem("Token Balance", "View your balance of a paricular token", withNextState(stateShowTokenBalance)),
		newMenuItem("New Transaction", "Create a new transaction", withNextState(stateNewTx)),
		newMenuItem("New Token Transfer", "Create a new token transfer", withNextState(stateNewTokenTx)),
		newMenuItem("Sign Message", "Sign a message", withNextState(stateNewMessage)),
		newMenuItem("Set RPC", "Set the RPC URL", withNextState(stateSetRPC)),
		newMenuItem("Save Keystore", "Save the wallet to a keystore file", withNextState(stateKSStore)),
		newMenuItem("Quit", "Quit the application", withNextState(stateQuit), withCallback(quitCallback)),
	}
}
