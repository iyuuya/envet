package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"

	"github.com/iyuuya/envet/repo"
)

var exportCmd = &cobra.Command{
	Use:     "export",
	Short:   "print export environment variables",
	Aliases: []string{"exp", "e"},
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		repo := repo.NewKeychainRepository()
		entries, err := repo.ListKeyValues(args[0])
		if err != nil {
			return err
		}
		for _, entry := range entries {
			fmt.Printf("export %s=%s\n", entry.Key, entry.Value)
		}
		return nil
	},
}

var unsetCmd = &cobra.Command{
	Use:     "unset",
	Short:   "print unset environment variables",
	Aliases: []string{"u"},
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		repo := repo.NewKeychainRepository()
		keys, err := repo.ListKeys(args[0])
		if err != nil {
			return err
		}
		for _, key := range keys {
			fmt.Printf("unset %s\n", key)
		}
		return nil
	},
}

var runCmd = &cobra.Command{
	Use:     "run",
	Short:   "run command with environment variables",
	Aliases: []string{"r", "sh", "exec"},
	Args:    cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		repo := repo.NewKeychainRepository()
		entries, err := repo.ListKeyValues(args[0])
		if err != nil {
			return err
		}
		cmdArgs := args[1:]
		if len(cmdArgs) == 0 {
			return fmt.Errorf("command is required")
		}
		execCmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
		execCmd.Env = os.Environ()
		for _, entry := range entries {
			execCmd.Env = append(execCmd.Env, fmt.Sprintf("%s=%s", entry.Key, entry.Value))
		}
		execCmd.Stdin = os.Stdin
		execCmd.Stdout = os.Stdout
		execCmd.Stderr = os.Stderr
		return execCmd.Run()
	},
}
