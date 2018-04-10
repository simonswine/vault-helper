package cmd

import (
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/jetstack/vault-helper/pkg/cert"
	"github.com/jetstack/vault-helper/pkg/kubeconfig"
)

// initCmd represents the init command
var kubeconfCmd = &cobra.Command{
	Use:   "kubeconfig [cert role] [common name] [cert path] [kubeconfig path]",
	Short: "Create local key to generate a CSR. Call vault with CSR for specified cert role. Write kubeconfig to yaml file.",
	Run: func(cmd *cobra.Command, args []string) {
		log := LogLevel(cmd)

		if len(args) != 4 {
			log.Fatal("Wrong number of arguments given.\nUsage: vault-helper kubeconfig [cert role] [common name] [cert path] [kubeconfig path]")
		}

		i, err := newInstanceToken(cmd)
		if err != nil {
			i.Log.Fatal(err)
		}
		if err := i.TokenRenewRun(); err != nil {
			i.Log.Fatal(err)
		}

		c := cert.New(i.Log, i)
		if err := setFlagsCert(c, cmd, args); err != nil {
			c.Log.Fatal(err)
		}

		abs, err := filepath.Abs(args[3])
		if err != nil {
			log.Fatalf("error generating absolute path from destination '%s': %v", args[3], err)
		}

		u := kubeconfig.New(log, c)
		u.SetKubeConfigPath(abs)

		if err := c.RunCert(); err != nil {
			c.Log.Fatal(err)
		}

		if err := u.RunKube(); err != nil {
			u.Log.Fatal(err)
		}
	},
}

func init() {
	InitCertCmdFlags(kubeconfCmd)
}
