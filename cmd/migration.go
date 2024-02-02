/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/Marrquito/GoToDoAPP/database"

	"github.com/spf13/cobra"
)

var disableForeingKeys bool

// migrationCmd represents the migration command
var migrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "DB migration",
	Long:  `Runs a migration on postgres DB, by default`,
	Run: func(cmd *cobra.Command, args []string) {
		database.MigrateAll(disableForeingKeys)
	},
}

func init() {
	rootCmd.AddCommand(migrationCmd)
	migrationCmd.Flags().BoolVarP(&disableForeingKeys, "disable-foreing-keys", "d", true, "Disable automigrate with foreing keys")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrationCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
