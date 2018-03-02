package main

import (
	"github.com/spf13/cobra"

	cmds "project/cmd/commands"
)

func main() {

	var RootCmd = &cobra.Command{
		Use:   "kchain",
		Short: "Tendermint Core (BFT Consensus) in Go",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
			return nil
		},
	}
	RootCmd.AddCommand(
		cmds.Hello,
	)
	if err := RootCmd.Execute(); err != nil {
		panic(err)
	}
}
