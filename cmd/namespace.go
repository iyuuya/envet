package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/iyuuya/envet/repo"
)

var listNamespacesCmd = &cobra.Command{
	Use: "list-namespaces",
	Short: "list all namespaces",
	Aliases: []string{"lsn"},
	RunE: func(cmd *cobra.Command, args []string) error {
		repo := repo.NewKeychainRepository()
		ns, err := repo.ListNamespaces()
		if err != nil {
			return err
		}
		for _, n := range ns {
			fmt.Println(n)
		}
		return nil
	},
}

var removeNamespaceCmd = &cobra.Command{
	Use: "remove-namespace",
	Short: "remove a namespace",
	Aliases: []string{"rmn"},
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		repo := repo.NewKeychainRepository()
		err := repo.RemoveNamespace(args[0])
		if err != nil {
			return err
		}
		return nil
	},
}
