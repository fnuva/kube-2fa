package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"kube-2fa/internal/tools"
)

func init() {
	rootCmd.AddCommand(applyCommand)
}

var (
	applyCommand = &cobra.Command{
		Use:   "apply",
		Short: "mfa control tools",
		Long:  `mfa control tools`,
		Run: func(cmd *cobra.Command, args []string) {

			apiType := viper.GetString("mfa_config.selected")
			api := tools.Tools[apiType]
			err := api.Init()
			if err != nil {
				fmt.Println(err)
				return
			}
			response, err := api.Run()
			if err != nil {
				fmt.Println(err)
				return
			}
			if err = api.Apply([]string{tools.FileName}, &response); err != nil {
				fmt.Println(err)
				return
			}
		},
	}
)

func init() {
	cobra.OnInitialize(initConfig)
	applyCommand.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kube_mfa.yaml)")
	applyCommand.Flags().StringVarP(&tools.Code, "code", "c", "", "mfa code")
	applyCommand.Flags().StringVarP(&tools.FileName, "fileName", "f", "", "fileName")
	applyCommand.MarkFlagRequired("fileName")
}

func initConfig() {
	viper.SetConfigType("yaml")
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			er(err)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".kube_mfa")
	}

	viper.AutomaticEnv()
	viper.ReadInConfig()

}
