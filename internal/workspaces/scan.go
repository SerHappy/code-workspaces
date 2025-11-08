package workspaces

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

const workspaceExt = ".code-workspace"

var ignoreDirs = []string{".links", "python_wrappers"}

func Root(customRoot string) (string, error) {
	if customRoot == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(homeDir, "Projects"), nil
	}

	root := customRoot

	if strings.HasPrefix(root, "~") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("expand home directory: %w", err)
		}
		root = filepath.Join(homeDir, root[1:])
	}

	absRoot, err := filepath.Abs(root)
	if err != nil {
		return "", fmt.Errorf("convert to absolute path: %w", err)
	}

	info, err := os.Stat(absRoot)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("directory does not exist: %s", absRoot)
		}
		return "", fmt.Errorf("stat directory: %w", err)
	}

	if !info.IsDir() {
		return "", fmt.Errorf("not a directory: %s", absRoot)
	}

	return absRoot, nil
}

func Scan(root string) ([]Workspace, error) {
	return scanDir(root, root)
}

func scanDir(root string, dirAbsPath string) ([]Workspace, error) {
	entries, err := os.ReadDir(dirAbsPath)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if filepath.Ext(entry.Name()) == workspaceExt {
			relDir, err := filepath.Rel(root, dirAbsPath)
			if err != nil {
				return nil, err
			}
			return []Workspace{{
				FileAbs: filepath.Join(dirAbsPath, entry.Name()),
				RelDir:  relDir,
			},
			}, nil
		}
	}

	var result []Workspace
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		localName := entry.Name()
		if slices.Contains(ignoreDirs, localName) {
			continue
		}

		subdirPath := filepath.Join(dirAbsPath, localName)
		subResult, err := scanDir(root, subdirPath)
		if err != nil {
			return nil, err
		}
		if len(subResult) > 0 {
			result = append(result, subResult...)
		}
	}

	return result, nil
}
