package main

import (
	"fmt"
	"os"

	"github.com/notional-labs/demerklizator"
	"github.com/spf13/cobra"
)

var version = "0.0.1"
var rootCmd = &cobra.Command{
	Use:     "demerk",
	Version: version,
	Short:   "demerk -  a tool to convert iavl merklized data to normal db data",

	Run: func(cmd *cobra.Command, args []string) {

		fromPath := args[0]
		outPath := args[1]
		//convert fromPath before parsing to from field
		fromPath = demerklizator.ApplicationDBPathFromRootDir(fromPath)
		demerklizator.MigrateLatestStateDataToDBStores(fromPath, outPath)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s', please check your path", err)
		os.Exit(1)
	}
}