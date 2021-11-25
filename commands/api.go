package commands

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/gosom/go-sse-events-demo/api"
	"github.com/gosom/go-sse-events-demo/services"
)

func init() {
	var apiCmd = &cobra.Command{
		Use:   "api",
		Short: "api",
		Run: func(cmd *cobra.Command, args []string) {
			container, err := services.NewContainer()
			if err != nil {
				panic(err)
			}
			srv, err := api.New(container)
			if err != nil {
				panic(err)
			}
			if err := srv.Start(context.Background()); err != nil {
				panic(err)
			}
		},
	}
	RootCmd.AddCommand(apiCmd)
}
