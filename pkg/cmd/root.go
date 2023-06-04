package cmd

import "github.com/spf13/cobra"

func New() *cobra.Command {
	root := &cobra.Command{
		Use: "sec",
	}
	root.AddCommand(newInitCmd())

	return root
}
