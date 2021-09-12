package dashboard

import (
	"context"

	"github.com/codefresh-io/argo-rollouts/pkg/kubectl-argo-rollouts/options"
	"github.com/codefresh-io/argo-rollouts/server"
	"github.com/spf13/cobra"
)

func NewCmdDashboard(o *options.ArgoRolloutsOptions) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "dashboard",
		Short: "Start UI dashboard",
		RunE: func(c *cobra.Command, args []string) error {
			namespace := o.Namespace()
			kubeclientset := o.KubeClientset()
			rolloutclientset := o.RolloutsClientset()

			opts := server.ServerOptions{
				Namespace:         namespace,
				KubeClientset:     kubeclientset,
				RolloutsClientset: rolloutclientset,
				DynamicClientset:  o.DynamicClientset(),
			}

			for {
				ctx := context.Background()
				ctx, cancel := context.WithCancel(ctx)
				argorollouts := server.NewServer(opts)
				argorollouts.Run(ctx, 3100, true)
				cancel()
			}
		},
	}

	return cmd
}
