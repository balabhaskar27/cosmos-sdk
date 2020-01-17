package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
)

// GetQueryCmd returns the query commands for IBC clients
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	ics02ClientQueryCmd := &cobra.Command{
		Use:                        "client",
		Short:                      "IBC client query subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
	}

	ics02ClientQueryCmd.AddCommand(flags.GetCommands(
		GetCmdQueryClientStates(queryRoute, cdc),
		GetCmdQueryClientState(queryRoute, cdc),
		GetCmdQueryConsensusState(queryRoute, cdc),
		GetCmdQueryRoot(queryRoute, cdc),
		GetCmdQueryHeader(cdc),
		GetCmdNodeConsensusState(queryRoute, cdc),
		GetCmdQueryPath(queryRoute, cdc),
	)...)
	return ics02ClientQueryCmd
}

// GetTxCmd returns the transaction commands for IBC clients
func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	ics02ClientTxCmd := &cobra.Command{
		Use:                        "client",
		Short:                      "Client transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
	}

	ics02ClientTxCmd.AddCommand(flags.PostCommands(
		GetCmdCreateClient(cdc),
		GetCmdUpdateClient(cdc),
	)...)

	return ics02ClientTxCmd
}