package cmd

import (
	"dolittle.io/cookie-oidc-client/configuration"
	"dolittle.io/cookie-oidc-client/configuration/viper"
	"github.com/spf13/cobra"
)

var serve = &cobra.Command{
	Use:   "serve",
	Short: "Starts the server",
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := viper.NewViperConfiguration()
		if err != nil {
			return err
		}
		return configuration.NewContainer(config).Server.Run()
	},
}
