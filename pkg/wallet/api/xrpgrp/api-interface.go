package xrpgrp

import (
	"github.com/btcsuite/btcd/chaincfg"

	"github.com/hiromaily/go-crypto-wallet/pkg/wallet/api/xrpgrp/xrp"
	"github.com/hiromaily/go-crypto-wallet/pkg/wallet/coin"
)

// Rippler Ripple Interface
type Rippler interface {
	RippleAdminer
	RipplePublicer
	RippleAPIer

	// ripple
	Close()
	CoinTypeCode() coin.CoinTypeCode
	GetChainConf() *chaincfg.Params
}

// RippleAPIer is RippleAPI interface
type RippleAPIer interface {
	// RippleAPI
	PrepareTransaction(senderAccount, receiverAccount string, amount float64) (*xrp.TxJSON, error)
	SignTransaction(txJSON *xrp.TxJSON, secret string) (string, string, error)
	SubmitTransaction(signedTx string) (*xrp.SentTxJSON, uint32, error)
	WaitValidation(targetledgerVarsion uint32) (uint32, error)
	GetTransaction(txID string) (string, error)
}

// RipplePublicer is RipplePublic interface
type RipplePublicer interface {
	// public_account
	AccountChannels(sender, receiver string) (*xrp.ResponseAccountChannels, error)
	// public_server_info
	ServerInfo() (*xrp.ResponseServerInfo, error)
}

// RippleAdminer is RippleAdmin interface
type RippleAdminer interface {
	// admin_keygen
	ValidationCreate(secret string) (*xrp.ResponseValidationCreate, error)
	WalletProposeWithKey(seed string, keyType xrp.XRPKeyType) (*xrp.ResponseWalletPropose, error)
	WalletPropose(passphrase string) (*xrp.ResponseWalletPropose, error)
}