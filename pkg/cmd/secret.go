package cmd

import (
	"fmt"
	//"strings"

	"github.com/spf13/cobra"

	//"k8s.io/client-go/tools/clientcmd"
	//"k8s.io/client-go/tools/clientcmd/api"

	"k8s.io/cli-runtime/pkg/genericclioptions"
)

type SecretOptions struct {
	configFlags *genericclioptions.ConfigFlags

	genericclioptions.IOStreams
}

// NewSecretOptions provides an instance of SecretOptions with default values
func NewSecretOptions(streams genericclioptions.IOStreams) *SecretOptions {
	return &SecretOptions{
		configFlags: genericclioptions.NewConfigFlags(),

		IOStreams: streams,
	}
}

// NewCmdSecret provides a cobra command wrapping SecretOptions
func NewCmdSecret(streams genericclioptions.IOStreams) *cobra.Command {
	o := NewSecretOptions(streams)

	cmd := &cobra.Command{
		Use:   "secret [action] [flags]",
		Short: "View or set the current namespace",
		//Example:      fmt.Sprintf(namespaceExample, "kubectl"),
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			fmt.Println("run called!")

			return nil
		},
	}

	//cmd.Flags().BoolVar(&o.listSecrets, "list", o.listSecrets, "if true, print the list of all namespaces in the current KUBECONFIG")
	o.configFlags.AddFlags(cmd.Flags())

	return cmd
}
