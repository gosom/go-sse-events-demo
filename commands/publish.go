package commands

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/gosom/go-sse-events-demo/publisher"
	"github.com/gosom/go-sse-events-demo/services"
)

func init() {
	var publishCmd = &cobra.Command{
		Use:   "publish",
		Short: "publish",
		Run: func(cmd *cobra.Command, args []string) {
			container, err := services.NewContainer()
			if err != nil {
				panic(err)
			}
			pub, err := publisher.New(container)
			if err != nil {
				panic(err)
			}
			if err := pub.Start(context.Background()); err != nil {
				panic(err)
			}
		},
	}
	RootCmd.AddCommand(publishCmd)
}
