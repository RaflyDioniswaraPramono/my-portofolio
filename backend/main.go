package main

import (
	"github.com/RaflyDioniswaraPramono/my-portofolio/cmd/db/migrate"
	"github.com/RaflyDioniswaraPramono/my-portofolio/cmd/db/seed"
	"github.com/RaflyDioniswaraPramono/my-portofolio/cmd/serve"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use: "app",
	}

	// app command for running the server
	rootCmd.AddCommand(serve.Serve)

	// db command for migrations and seeding
	rootCmd.AddCommand(migrate.Migrate)
	rootCmd.AddCommand(migrate.UndoMigrate)
	rootCmd.AddCommand(seed.Seed)

	if err := rootCmd.Execute(); err != nil {
		panic("Failed to execute the program!")
	}
}
