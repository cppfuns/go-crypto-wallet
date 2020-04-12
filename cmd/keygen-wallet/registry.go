package main

import (
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"

	"github.com/hiromaily/go-bitcoin/pkg/config"
	mysql "github.com/hiromaily/go-bitcoin/pkg/db/rdb"
	"github.com/hiromaily/go-bitcoin/pkg/enum"
	"github.com/hiromaily/go-bitcoin/pkg/logger"
	"github.com/hiromaily/go-bitcoin/pkg/model/rdb"
	"github.com/hiromaily/go-bitcoin/pkg/model/rdb/coldrepo"
	"github.com/hiromaily/go-bitcoin/pkg/tracer"
	"github.com/hiromaily/go-bitcoin/pkg/txfile"
	"github.com/hiromaily/go-bitcoin/pkg/wallets"
	"github.com/hiromaily/go-bitcoin/pkg/wallets/api"
	"github.com/hiromaily/go-bitcoin/pkg/wallets/coldwallet"
	"github.com/hiromaily/go-bitcoin/pkg/wallets/key"
	"github.com/hiromaily/go-bitcoin/pkg/wallets/types"
)

// Registry is for registry interface
type Registry interface {
	NewKeygener() wallets.Keygener
}

type registry struct {
	conf        *config.Config
	mysqlClient *sqlx.DB
	logger      *zap.Logger
	rpcClient   *rpcclient.Client
	walletType  types.WalletType
}

// NewRegistry is to register registry interface
func NewRegistry(conf *config.Config, walletType types.WalletType) Registry {
	return &registry{
		conf:       conf,
		walletType: walletType,
	}
}

// NewKeygener is to register for keygener interface
func (r *registry) NewKeygener() wallets.Keygener {
	//TODO: should be interface
	r.setFilePath()

	return coldwallet.NewColdWalet(
		r.newBTC(),
		r.newLogger(),
		r.newTracer(),
		r.newStorager(),
		r.walletType,
	)

	//FIXME: wallet.NewWallet doesn't have rdb.KeygenStorager
	// How should it be fixed?? NewKeygen should be defined based on NewWallet
	//return keygen.NewKeygen(
	//	r.newBTC(),
	//	r.newLogger(),
	//	r.newTracer(),
	//	r.newStorager(),
	//	r.walletType,
	//	r.newColdWallet(),
	//)
}

func (r *registry) newRPCClient() *rpcclient.Client {
	var err error
	if r.rpcClient == nil {
		r.rpcClient, err = api.NewRPCClient(&r.conf.Bitcoin)
	}
	if err != nil {
		panic(err)
	}
	return r.rpcClient
}

func (r *registry) newBTC() api.Bitcoiner {
	bit, err := api.NewBitcoin(r.newRPCClient(), &r.conf.Bitcoin, r.newLogger(), enum.CoinType(r.conf.CoinType))
	if err != nil {
		panic(err)
	}
	return bit
}

func (r *registry) newLogger() *zap.Logger {
	if r.logger == nil {
		r.logger = logger.NewZapLogger(&r.conf.Logger)
	}
	return r.logger
}

func (r *registry) newTracer() opentracing.Tracer {
	return tracer.NewTracer(r.conf.Tracer)
}

//func (r *registry) newStorager() rdb.KeygenStorager {
func (r *registry) newStorager() rdb.ColdStorager {
	// if there are multiple options, set proper one
	// storager interface as MySQL
	return coldrepo.NewColdRepository(
		r.newMySQLClient(),
		r.newLogger(),
	)
}

func (r *registry) newMySQLClient() *sqlx.DB {
	if r.mysqlClient == nil {
		dbConn, err := mysql.NewMySQL(&r.conf.MySQL)
		if err != nil {
			panic(err)
		}
		r.mysqlClient = dbConn
	}
	return r.mysqlClient
}

//TODO: move to somewhere
func (r *registry) setFilePath() {
	// TxFile
	if r.conf.TxFile.BasePath != "" {
		txfile.SetFilePath(r.conf.TxFile.BasePath)
	}

	// PubkeyCSV
	if r.conf.PubkeyFile.BasePath != "" {
		key.SetFilePath(r.conf.PubkeyFile.BasePath)
	}
}