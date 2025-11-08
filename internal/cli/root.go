package cli

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/serhappy/code-workspaces/internal/workspaces"
	"github.com/spf13/cobra"
)

var rootDir string
var rootCmd = &cobra.Command{
	Use:   "cw [relative-directory]",
	Short: "Open VS Code workspace by relative directory name",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		target := args[0]

		rootDir, err := cmd.Flags().GetString("root")
		if err != nil {
			return fmt.Errorf("get root directory: %w", err)
		}

		root, err := workspaces.Root(rootDir)
		if err != nil {
			return fmt.Errorf("determine projects root: %w", err)
		}

		wsList, err := workspaces.Scan(root)
		if err != nil {
			return fmt.Errorf("scan workspaces: %w", err)
		}

		if len(wsList) == 0 {
			return fmt.Errorf("no workspaces found under %s", root)
		}

		index := workspaces.BuildIndexByRelPath(wsList)
		ws, ok := index[target]
		if !ok {
			fmt.Printf("Workspace %q not found.\n", target)
			fmt.Println("Available keys:")
			for _, w := range wsList {
				fmt.Println(" -", w.RelDir)
			}
			return nil
		}

		fmt.Printf("Opening workspace: %s\n", ws.FileAbs)

		cmdCode := exec.Command("code", "-r", ws.FileAbs)
		cmdCode.Stdout = os.Stdout
		cmdCode.Stderr = os.Stderr

		if err := cmdCode.Run(); err != nil {
			return fmt.Errorf("run code: %w", err)
		}

		return nil
	},
}

func init() {
	rootCmd.ValidArgsFunction = completeWorkspaces
	rootCmd.PersistentFlags().StringVarP(&rootDir, "root", "r", "", "custom root directory")
}

func completeWorkspaces(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	rootDir, err := cmd.Flags().GetString("root")
	if err != nil {
		return nil, cobra.ShellCompDirectiveError
	}

	root, err := workspaces.Root(rootDir)
	if err != nil {
		return nil, cobra.ShellCompDirectiveError
	}

	wsList, err := workspaces.Scan(root)
	if err != nil {
		return nil, cobra.ShellCompDirectiveError
	}

	keys := workspaces.Keys(wsList)
	var suggestions []string
	for _, key := range keys {
		if toComplete == "" || strings.HasPrefix(key, toComplete) {
			suggestions = append(suggestions, key)
		}
	}

	return suggestions, cobra.ShellCompDirectiveNoFileComp
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
