package seed

import (
	"github.com/RaflyDioniswaraPramono/my-portofolio/api/http/role"
	"github.com/RaflyDioniswaraPramono/my-portofolio/api/http/user"
	"github.com/RaflyDioniswaraPramono/my-portofolio/configs/database"
	"github.com/RaflyDioniswaraPramono/my-portofolio/utils"
	"github.com/spf13/cobra"
)

var Seed = &cobra.Command{
	Use:   "db:seed",
	Short: "Database seeding command!",
	Run: func(cmd *cobra.Command, args []string) {
		database.DBInit()

		var db = *database.DB

		utils.WarningCLIMessage("Running seeding ... \n")

		role.RoleSeeder(&db)
		user.UserSeeder(&db)

		utils.SuccessCLIMessage("Seeding succeed!")
	},
}
