package serve

import (
	"fmt"

	"github.com/RaflyDioniswaraPramono/my-portofolio/app"
	"github.com/spf13/cobra"
)

var stage string

var Serve = &cobra.Command{
	Use:   "serve",
	Short: "Serve to running the application server!",
	Run: func(cmd *cobra.Command, args []string) {
		app.RunServer()
	},
}

func init() {
	message := fmt.Sprintf("Start the %v server!", stage)

	Serve.Flags().StringVarP(&stage, "stage", "s", "", message)
}
