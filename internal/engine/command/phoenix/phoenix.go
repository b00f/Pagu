package phoenix

import (
	"fmt"

	"github.com/pagu-project/Pagu/internal/engine/command"
	"github.com/pagu-project/Pagu/internal/entity"
	"github.com/pagu-project/Pagu/internal/repository"
	"github.com/pagu-project/Pagu/pkg/client"
	"github.com/pagu-project/Pagu/pkg/wallet"
)

const (
	CommandName         = "phoenix"
	FaucetCommandName   = "faucet"
	WalletCommandName   = "wallet"
	StatusCommandName   = "status"
	HealthCommandName   = "health"
	NodeInfoCommandName = "node-info"
	HelpCommandName     = "help"
)

type Phoenix struct {
	wallet       wallet.IWallet
	db           repository.Database
	clientMgr    client.Manager
	faucetAmount uint
}

func NewPhoenix(wallet wallet.IWallet, faucetAmount uint, clientMgr client.Manager, db repository.Database) Phoenix {
	return Phoenix{
		wallet:       wallet,
		clientMgr:    clientMgr,
		db:           db,
		faucetAmount: faucetAmount,
	}
}

func (pt *Phoenix) GetCommand() command.Command {
	middlewareHandler := command.NewMiddlewareHandler(pt.db, pt.wallet)

	subCmdStatus := command.Command{
		Name:        StatusCommandName,
		Help:        "Phoenix Testnet statistics",
		Args:        []command.Args{},
		SubCommands: nil,
		AppIDs:      entity.AllAppIDs(),
		Middlewares: []command.MiddlewareFunc{middlewareHandler.CreateUser},
		Handler:     pt.networkStatusHandler,
		TargetFlag:  command.TargetMaskTest,
	}

	subCmdFaucet := command.Command{
		Name: FaucetCommandName,
		Help: fmt.Sprintf("Get %d tPAC Coins on Phoenix Testnet for Testing your code or project", pt.faucetAmount),
		Args: []command.Args{
			{
				Name:     "address",
				Desc:     "your testnet address [example: tpc1z...]",
				Optional: false,
			},
		},
		SubCommands: nil,
		AppIDs:      entity.AllAppIDs(),
		Middlewares: []command.MiddlewareFunc{middlewareHandler.CreateUser, middlewareHandler.WalletBalance},
		Handler:     pt.faucetHandler,
		TargetFlag:  command.TargetMaskTest,
	}

	cmdPhoenix := command.Command{
		Name:        CommandName,
		Help:        "Phoenix Testnet tools and utils for developers",
		Args:        nil,
		AppIDs:      entity.AllAppIDs(),
		SubCommands: make([]command.Command, 0),
		Handler:     nil,
		TargetFlag:  command.TargetMaskTest,
	}

	cmdPhoenix.AddSubCommand(subCmdFaucet)
	cmdPhoenix.AddSubCommand(subCmdStatus)

	return cmdPhoenix
}
