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
		root, err := workspaces.Root()
		if err != nil {
			return fmt.Errorf("determine projects root: %w", err)
		}

		wsList, err := workspaces.Scan(root)
		if err != nil {
			return fmt.Errorf("scan workspaces: %w", err)
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
