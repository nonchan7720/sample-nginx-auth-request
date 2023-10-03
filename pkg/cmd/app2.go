package cmd

import (
	"context"

	"github.com/nonchan7720/sample-nginx-auth-request/pkg/handler"
	"github.com/spf13/cobra"
)

func app2Command() *cobra.Command {
	cmd := cobra.Command{
		Use: "app2",
		Run: func(cmd *cobra.Command, args []string) {
			app2Server(cmd.Context())
		},
	}

	return &cmd
}

func app2Server(ctx context.Context) {
	engine := handler.NewApp2Server()
	_ = engine.Run(":3020")
}
