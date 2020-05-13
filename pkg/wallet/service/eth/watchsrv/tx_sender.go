package watchsrv

import (
	"database/sql"
	"strings"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/hiromaily/go-bitcoin/pkg/repository/watchrepo"
	"github.com/hiromaily/go-bitcoin/pkg/tx"
	"github.com/hiromaily/go-bitcoin/pkg/wallet"
	"github.com/hiromaily/go-bitcoin/pkg/wallet/api/ethgrp"
)

// TxSend type
type TxSend struct {
	eth          ethgrp.Ethereumer
	logger       *zap.Logger
	dbConn       *sql.DB
	addrRepo     watchrepo.AddressRepositorier
	txRepo       watchrepo.ETHTxRepositorier
	txDetailRepo watchrepo.EthDetailTxRepositorier
	txFileRepo   tx.FileRepositorier
	wtype        wallet.WalletType
}

// NewTxSend returns TxSend object
func NewTxSend(
	eth ethgrp.Ethereumer,
	logger *zap.Logger,
	dbConn *sql.DB,
	addrRepo watchrepo.AddressRepositorier,
	txRepo watchrepo.ETHTxRepositorier,
	txDetailRepo watchrepo.EthDetailTxRepositorier,
	txFileRepo tx.FileRepositorier,
	wtype wallet.WalletType) *TxSend {

	return &TxSend{
		eth:          eth,
		logger:       logger,
		dbConn:       dbConn,
		addrRepo:     addrRepo,
		txRepo:       txRepo,
		txDetailRepo: txDetailRepo,
		txFileRepo:   txFileRepo,
		wtype:        wtype,
	}
}

// SendTx send signed tx by keygen/sign walet
func (t *TxSend) SendTx(filePath string) (string, error) {

	// get tx_deposit_id from file name
	//payment_5_unsigned_1_1534466246366489473
	actionType, _, txID, _, err := t.txFileRepo.ValidateFilePath(filePath, tx.TxTypeSigned)
	if err != nil {
		return "", errors.Wrap(err, "fail to call txFileRepo.ValidateFilePath()")
	}

	t.logger.Debug("send_tx", zap.String("action_type", actionType.String()))

	// read hex from file
	data, err := t.txFileRepo.ReadFileSlice(filePath)
	if err != nil {
		return "", errors.Wrap(err, "fail to call txFileRepo.ReadFile()")
	}

	//sentTxes := make([]string, 0, len(data))
	for _, txHex := range data {
		// data is csv [rawTx.TxHex, signedRawTx.TxHex]
		// rawTx.TxHex is used to record status by updating database
		tmp := strings.Split(txHex, ",")
		if len(tmp) != 2 {
			return "", errors.New("data format is invalid in file")
		}
		unsignedTx := tmp[0]
		signedTx := tmp[1]

		// sign
		sentTx, err := t.eth.SendSignedRawTransaction(signedTx)
		if err != nil {
			t.logger.Warn("fail to call eth.SendSignedRawTransaction()",
				zap.Error(err),
			)
			continue
		}
		if sentTx == "" {
			t.logger.Warn("no sentTx by calling eth.SendSignedRawTransaction()",
				zap.Error(err),
			)
			continue
		}
		// signedRawTx.TxHex
		//sentTxes = append(sentTxes, fmt.Sprintf("%s,%s", unsignedTx, sentTx))

		// update eth_detail_tx
		affectedNum, err := t.txDetailRepo.UpdateAfterTxSent(txID, unsignedTx, tx.TxTypeSent, signedTx, sentTx)
		if err != nil {
			//TODO: even if error occurred, tx is already sent. so db should be corrected manually
			t.logger.Warn("fail to call repo.Tx().UpdateAfterTxSent() but tx is already sent. So database should be updated manually",
				zap.Int64("tx_id", txID),
				zap.String("tx_type", tx.TxTypeSent.String()),
				zap.Int8("tx_type_value", tx.TxTypeSent.Int8()),
				zap.String("signed_hex_tx", signedTx),
				zap.String("sent_hash_tx", sentTx),
			)
			return "", errors.Wrapf(err, "fail to call updateHexForSentTx(), but tx is sent. txID: %d", txID)
		}
		if affectedNum == 0 {
			t.logger.Info("no records to update tx_table",
				zap.Int64("tx_id", txID),
				zap.String("tx_type", tx.TxTypeSent.String()),
				zap.Int8("tx_type_value", tx.TxTypeSent.Int8()),
				zap.String("signed_hex_tx", signedTx),
				zap.String("sent_hash_tx", sentTx),
			)
			return "", nil
		}
	}

	// update account_pubkey_table
	//if actionType != action.ActionTypePayment {
	//	//skip for that receiver address is anonymous
	//	err = t.updateIsAllocatedAccountPubkey(txID)
	//	if err != nil {
	//		//TODO: even if error occurred, tx is already sent. so db should be corrected manually
	//		return "", err
	//	}
	//}
	return "", nil
}

//func (t *TxSend) updateIsAllocatedAccountPubkey(txID int64) error {
//	// get txOutputs by tx_id
//	txOutputs, err := t.txOutputRepo.GetAllByTxID(txID)
//	if err != nil {
//		return errors.Wrap(err, "fail to call repo.TxOutput().GetAllByTxID()")
//	}
//	if len(txOutputs) == 0 {
//		return errors.New("output tx could not be found in tx_deposit_output")
//	}
//
//	_, err = t.addrRepo.UpdateIsAllocated(true, txOutputs[0].OutputAddress)
//	if err != nil {
//		return errors.Wrap(err, "fail to call repo.Pubkey().UpdateIsAllocated()")
//	}
//
//	return nil
//}