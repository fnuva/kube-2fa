package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	cfgFile string

	rootCmd = &cobra.Command{
		Use:                "kube-2fa",
		Short:              "mfa control tools",
		Long:               `mfa control tools`,
		DisableFlagParsing: true,
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}
