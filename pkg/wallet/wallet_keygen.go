package wallet

import (
	"github.com/hiromaily/go-bitcoin/pkg/account"
	"github.com/hiromaily/go-bitcoin/pkg/enum"
	"github.com/hiromaily/go-bitcoin/pkg/model"
	"github.com/hiromaily/go-bitcoin/pkg/wallet/api"
	"github.com/hiromaily/go-bitcoin/pkg/wallet/key"
)

// Keygener is for keygen wallet service interface
type Keygener interface {
	SignatureFromFile(filePath string) (string, bool, string, error)
	GenerateSeed() ([]byte, error)
	GenerateAccountKey(accountType account.AccountType, coinType enum.CoinType, seed []byte, count uint32) ([]key.WalletKey, error)
	ImportPrivateKey(accountType account.AccountType) error
	ExportAccountKey(accountType account.AccountType, keyStatus enum.KeyStatus) (string, error)
	ImportMultisigAddrForColdWallet1(fileName string, accountType account.AccountType) error
	Done()
	GetDB() *model.DB
	GetBTC() api.Bitcoiner
	GetType() WalletType
	GetSeed() string
}

// KeygenWallet keygen wallet object
type KeygenWallet struct {
	BTC  api.Bitcoiner
	DB   *model.DB //TODO:should be interface
	Type WalletType
	Seed string
}

func NewKeygenWallet(bit api.Bitcoiner, rds *model.DB, typ WalletType, seed string) *KeygenWallet {
	return &KeygenWallet{
		BTC:  bit,
		DB:   rds,
		Type: typ,
		Seed: seed,
	}
}

// Done should be called before exit
func (w *KeygenWallet) Done() {
	w.DB.RDB.Close()
	w.BTC.Close()
}

func (w *KeygenWallet) GetDB() *model.DB {
	return w.DB
}

func (w *KeygenWallet) GetBTC() api.Bitcoiner {
	return w.BTC
}

func (w *KeygenWallet) GetType() WalletType {
	return w.Type
}

func (w *KeygenWallet) GetSeed() string {
	return w.Seed
}
