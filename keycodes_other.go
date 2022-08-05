//go:build !windows
// +build !windows

package promptui

import "github.com/byzk-project-deploy/readline"

var (
	// KeyBackspace is the default key for deleting input text.
	KeyBackspace rune = readline.CharBackspace
)
