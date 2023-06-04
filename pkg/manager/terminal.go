package manager

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func readPassFromTerminal(prompt string) (string, error) {
	fmt.Print(prompt)

	pass, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return "", fmt.Errorf("failed to read password: %w", err)
	}

	return string(pass), nil
}

func readPassWithConfirm() (string, error) {
	pass, err := readPassFromTerminal("Password:")
	if err != nil {
		return "", fmt.Errorf("failed to read password: %w", err)
	}

	fmt.Println()

	confirm, err := readPassFromTerminal("Confirm:")
	if err != nil {
		return "", fmt.Errorf("failed read password confirmation: %w", err)
	}

	if pass != confirm {
		return "", fmt.Errorf("password and confirmation not match")
	}

	fmt.Println()

	return pass, nil
}
