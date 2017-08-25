package cmd

import (
	"github.com/Sirupsen/logrus"
	vault "github.com/hashicorp/vault/api"
	"github.com/spf13/cobra"

	"github.com/jetstack-experimental/vault-helper/pkg/instanceToken"
)

// initCmd represents the init command
var renewtokenCmd = &cobra.Command{
	Use:   "renew-token [cluster ID]",
	Short: "Renew token on vault server.",
	Run: func(cmd *cobra.Command, args []string) {
		logger := logrus.New()

		n, err := RootCmd.PersistentFlags().GetInt("log-level")
		if err != nil {
			logrus.Fatalf("failed to get log level of flag: %s", err)
		}
		if n < 0 || n > 2 {
			logrus.Fatalf("not a valid log level")
		}
		switch n {
		case 0:
			logger.Level = logrus.FatalLevel
		case 1:
			logger.Level = logrus.InfoLevel
		case 2:
			logger.Level = logrus.DebugLevel
		}
		log := logrus.NewEntry(logger)

		v, err := vault.NewClient(nil)
		if err != nil {
			logger.Fatal(err)
		}

		i := instanceToken.New(v, log)

		if err := i.Run(cmd, args); err != nil {
			logger.Fatal(err)
		}
	},
}

func init() {
	renewtokenCmd.PersistentFlags().String(instanceToken.FlagTokenRole, "", "Set role of token to renew. (default *no role*)")

	RootCmd.AddCommand(renewtokenCmd)
}
