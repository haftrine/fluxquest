package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/haftrine/fluxquest/internal/cli/flagparsers"
)

// RootCmd defines the root fluxquest command
var RootCmd = &cobra.Command{
	Use:   "fluxquest",
	Short: "Fluxquest migrates an InfluxDB database (or part of a database) to QuestDB",
	Long:  "Fluxquest migrates an InfluxDB database (or part of a database) to QuestDB",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Outflux version " + fluxQuestVersion)
		fmt.Println("Run 'fluxquest --help' for usage")
	},
}

// Execute is called to execute the root outflux command
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	RootCmd.PersistentFlags().Bool(flagparsers.QuietFlag, false, "If specified will suppress any log to STDOUT")
	RootCmd.Flags().Bool(flagparsers.VersionFlag, false, "Print the version of Outflux")
	migrateCmd := initMigrateCmd()
	RootCmd.AddCommand(migrateCmd)

	schemaTransferCmd := initSchemaTransferCmd()
	RootCmd.AddCommand(schemaTransferCmd)
}
