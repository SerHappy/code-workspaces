package cli

import (
	"fmt"

	"github.com/serhappy/code-workspaces/internal/workspaces"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all workspaces under projects root",
	RunE: func(cmd *cobra.Command, args []string) error {
		wsList, root, err := getWorkspaces(cmd)
		if err != nil {
			return err
		}

		if len(wsList) == 0 {
			fmt.Printf("No workspaces found under %s\n", root)
			return nil
		}

		keys := workspaces.Keys(wsList)
		for _, key := range keys {
			fmt.Println(key)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
