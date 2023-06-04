package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zarix908/passwords_store/pkg/manager"
)

func newInitCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "generate new master key encrypted with provided password",
		Run: func(*cobra.Command, []string) {
			if err := manager.GenerateMasterKey(); err != nil {
				panic(err)
			}
		},
	}
}
