package commands

import (
	"os"

	"github.com/spf13/cobra"

	// Kube client doesn't support all auth providers by default.
	// this ensures we include all backends supported by the client.
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

const mainUsage = `Interact with the Brigade cluster service.

Brigade is a tool for scripting cluster workflows, and 'brig' is the command
line client for interacting with Brigade.

The most common use for this tool is to send a Brigade JavaScript file to the
cluster for execution. This is done with the 'brigade run' command.

	$ brig run -f my.js my-project

But the 'brig' command can also be used for learning about projects and
builds as well.

By default, Brigade learns about your Kubernetes cluster by inspect the $KUBECONFIG
environment variable.
`

var (
	globalNamespace  = ""
	globalVerbose    = false
	globalKubeConfig = ""
)

func init() {
	f := Root.PersistentFlags()
	f.StringVarP(&globalNamespace, "namespace", "n", "default", "The Kubernetes namespace for Brigade")
	f.StringVar(&globalKubeConfig, "kubeconfig", "", "The path to a KUBECONFIG file, overrides $KUBECONFIG.")
	f.BoolVarP(&globalVerbose, "verbose", "v", false, "Turn on verbose output")
}

// Root is the root command
var Root = &cobra.Command{
	Use:   "brig",
	Short: "The Brigade client",
	Long:  mainUsage,
}

func kubeConfigPath() string {
	if globalKubeConfig != "" {
		return globalKubeConfig
	}
	if v, ok := os.LookupEnv(kubeConfig); ok {
		return v
	}
	return os.ExpandEnv("$HOME/.kube/config")
}
