package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/iyuuya/envet/repo"
)

var addCmd = &cobra.Command{
	Use: "add",
	Short: "add a new environment variable",
	Aliases: []string{"a"},
	Args: cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		repo := repo.NewKeychainRepository()
		ns := args[0]
		keys := args[1:]

		if len(keys) == 0 {
			return fmt.Errorf("at least one key is required")
		}
		for _, key := range keys {
			fmt.Printf("Enter value for %s: ", key)
			var value string
			_, err := fmt.Scanln(&value)
			if err != nil {
				return err
			}
			err = repo.AddEntry(ns, key, value)
			if err != nil {
				return err
			}
		}

		return nil
	},
}

var removeCmd = &cobra.Command{
	Use: "remove",
	Short: "remove an environment variable",
	Aliases: []string{"rm"},
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		repo := repo.NewKeychainRepository()
		return repo.RemoveEntry(args[0], args[1])
	},
}

var listKeysCmd = &cobra.Command{
	Use: "list-keys",
	Short: "list keys in a namespace",
	Aliases: []string{"lsk"},
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		repo := repo.NewKeychainRepository()
		keys, err := repo.ListKeys(args[0])
		if err != nil {
			return err
		}
		for _, k := range keys {
			fmt.Println(k)
		}

		return nil
	},
}
