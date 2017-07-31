package cmd

import (
	"github.com/Sirupsen/logrus"
	vault "github.com/hashicorp/vault/api"
	"github.com/jetstack-experimental/vault-helper/pkg/cert"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var certCmd = &cobra.Command{
	Use: "cert [cert role] [common name] [destination path]",
	// TODO: Make short better
	Short: "Create local key to generate a CSR. Call vault with CSR for specified cert role",
	Run: func(cmd *cobra.Command, args []string) {
		logger := logrus.New()
		logger.Level = logrus.DebugLevel
		log := logrus.NewEntry(logger)

		v, err := vault.NewClient(nil)
		if err != nil {
			logger.Fatal(err)
		}

		c := cert.New(v, log)

		if err := c.Run(cmd, args); err != nil {
			logger.Fatal(err)
		}

	},
}

func init() {
	certCmd.PersistentFlags().Int64(cert.FlagKeyBitSize, 2048, "Bit size used for generating key. [int]")
	certCmd.PersistentFlags().String(cert.FlagKeyType, "RSA", "Type of key to generate. [string]")
	certCmd.PersistentFlags().StringSlice(cert.FlagIpSans, []string{}, "IP sans. [[]string] (default none)")
	certCmd.PersistentFlags().StringSlice(cert.FlagSanHosts, []string{}, "Host Sans. [[]string] (default none)")
	certCmd.PersistentFlags().String(cert.FlagOwner, "root", "Owner of created file/directories. [string]")
	certCmd.PersistentFlags().String(cert.FlagGroup, "root", "Group of created file/directories. [string]")

	RootCmd.AddCommand(certCmd)
}
