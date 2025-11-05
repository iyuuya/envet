package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "envet",
	Short: "manage environment variables in keychain",
	Long: `envet is a tool to manage environment variables stored in the system keychain.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(listNamespacesCmd)
	rootCmd.AddCommand(listKeysCmd)
	rootCmd.AddCommand(exportCmd)
	rootCmd.AddCommand(unsetCmd)
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(removeNamespaceCmd)
	rootCmd.AddCommand(removeCmd)
}
