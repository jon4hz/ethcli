package ethcli

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/tyler-smith/go-bip39"
)

const defaultBasePath = "m/44'/60'/0'/0/"

type Wallet struct {
	mnemonic           string
	wallet             *hdwallet.Wallet
	baseDerivationPath string
	derivationCounter  int
	privateKey         *ecdsa.PrivateKey
	publicKey          *ecdsa.PublicKey
}

const defaultEntropySize = 256

func New() Wallet {
	entropy, _ := bip39.NewEntropy(defaultEntropySize)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	wallet, _ := hdwallet.NewFromMnemonic(mnemonic)
	w := Wallet{
		mnemonic:           mnemonic,
		wallet:             wallet,
		baseDerivationPath: defaultBasePath,
		derivationCounter:  0,
	}
	w.LoadFirstAccount()
	return w
}

func (w *Wallet) LoadFirstAccount() {
	w.derivationCounter = 0
	w.loadAccount()
}

func (w *Wallet) loadAccount() {
	path := hdwallet.MustParseDerivationPath(fmt.Sprintf("%s%d", w.baseDerivationPath, w.derivationCounter))
	a, _ := w.wallet.Derive(path, false)
	w.privateKey, _ = w.wallet.PrivateKey(a)
	w.publicKey, _ = w.wallet.PublicKey(a)
}

func (w *Wallet) PrivateKeyString() string {
	privateKeyBytes := crypto.FromECDSA(w.privateKey)
	return hexutil.Encode(privateKeyBytes)[2:]
}

func (w *Wallet) PublicKeyString() string {
	return crypto.PubkeyToAddress(*w.publicKey).Hex()
}
