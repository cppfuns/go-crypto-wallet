package wallets

//signature-wallet

import (
	"fmt"
	"go.uber.org/zap"

	"github.com/pkg/errors"

	"github.com/hiromaily/go-bitcoin/pkg/account"
	"github.com/hiromaily/go-bitcoin/pkg/enum"
)

// AddMultisigAddressByAuthorization account_key_authorizationテーブルのwallet_addressを認証者として、
// added_pubkey_history_paymentテーブル内のwalletアドレスのmultisigアドレスを生成する
// TODO:第4パラメータに、address_typeを追加する。Bitcoinの場合は、p2sh-segwit とする
func (w *Wallet) AddMultisigAddressByAuthorization(accountType account.AccountType, addressType enum.AddressType) error {
	//TODO:remove it
	if w.wtype != WalletTypeSignature {
		return errors.New("it's available on Coldwallet2")
	}

	//accountチェック
	//multisigであればこのチェックはOK
	//if accountType != enum.AccountTypeReceipt && accountType != enum.AccountTypePayment {
	//	logger.Info("AccountType should be AccountTypeReceipt or AccountTypePayment")
	//	return nil
	//}
	if !account.AccountTypeMultisig[accountType] {
		w.logger.Info("This func is for only account witch uses multiaddress")
		return nil
	}

	//account_key_authorizationテーブルからAuthorizationのwallet_addressを取得
	authKeyTable, err := w.storager.GetOneByMaxIDOnAccountKeyTable(account.AccountTypeAuthorization)
	if err != nil {
		return errors.Errorf("DB.GetOneByMaxIDOnAccountKeyTable(enum.AccountTypeAuthorization) error: %s", err)
	}

	//added_pubkey_history_xxxテーブルからwallet_address(full-pubkeyである必要がある)を取得
	addedPubkeyHistoryTable, err := w.storager.GetAddedPubkeyHistoryTableByNoWalletMultisigAddress(accountType)
	if err != nil {
		return errors.Errorf("DB.GetAddedPubkeyHistoryTableByNoWalletMultisigAddress(%s) error: %s", accountType, err)
	}

	//addmultisigaddress APIをcall
	//FIXME:multisigのN:Mは可変でも可能なようにロジックを組み立てる
	for _, val := range addedPubkeyHistoryTable {
		resAddr, err := w.btc.AddMultisigAddress(
			2,
			[]string{
				val.FullPublicKey, // receipt or payment address
				authKeyTable.P2shSegwitAddress,
			},
			fmt.Sprintf("multi_%s", accountType), //TODO:ここのアカウント名はどうすべきか
			addressType,
		)
		if err != nil {
			//[Error] -5: no full public key for address mkPmdpo59gpU7ZioGYwwoMTQJjh7MiqUvd
			w.logger.Error(
				"fail to call btc.CreateMultiSig(2,,) ",
				zap.String("full public key", val.FullPublicKey),
				zap.String("p2sh segwit address", authKeyTable.P2shSegwitAddress),
				zap.Error(err))
			continue
		}

		//レスポンスをadded_pubkey_history_xxxテーブルに保存
		err = w.storager.UpdateMultisigAddrOnAddedPubkeyHistoryTable(accountType, resAddr.Address,
			resAddr.RedeemScript, authKeyTable.P2shSegwitAddress, val.FullPublicKey, nil, true)
		if err != nil {
			w.logger.Error(
				"fail to call db.UpdateMultisigAddrOnAddedPubkeyHistoryTable()",
				zap.String("accountType", accountType.String()),
				zap.Error(err))
		}
	}

	return nil
}
