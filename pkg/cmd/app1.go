package cmd

import (
	"context"

	"github.com/nonchan7720/sample-nginx-auth-request/pkg/handler"
	"github.com/spf13/cobra"
)

func app1Command() *cobra.Command {
	cmd := cobra.Command{
		Use: "app1",
		Run: func(cmd *cobra.Command, args []string) {
			app1Server(cmd.Context())
		},
	}

	return &cmd
}

func app1Server(ctx context.Context) {
	engine := handler.NewApp1Server()
	_ = engine.Run(":3030")
}
