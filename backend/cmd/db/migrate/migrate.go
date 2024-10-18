package migrate

import (
	"log"

	"github.com/RaflyDioniswaraPramono/my-portofolio/api/http/bio"
	"github.com/RaflyDioniswaraPramono/my-portofolio/api/http/experience"
	"github.com/RaflyDioniswaraPramono/my-portofolio/api/http/portofolio"
	"github.com/RaflyDioniswaraPramono/my-portofolio/api/http/role"
	"github.com/RaflyDioniswaraPramono/my-portofolio/api/http/sosmed"
	"github.com/RaflyDioniswaraPramono/my-portofolio/api/http/user"
	"github.com/RaflyDioniswaraPramono/my-portofolio/app"
	"github.com/RaflyDioniswaraPramono/my-portofolio/configs/database"
	"github.com/RaflyDioniswaraPramono/my-portofolio/utils"
	"github.com/spf13/cobra"
)

var Migrate = &cobra.Command{
	Use:   "db:migrate",
	Short: "Database migration command!",
	Run: func(cmd *cobra.Command, args []string) {
		database.DBInit()

		utils.WarningCLIMessage("Running migrations ... \n")

		var db = *database.DB

		models := []interface{}{
			&user.User{},
			&role.Role{},
			&portofolio.Portofolio{},
			&experience.Experience{},
			&bio.Bio{},
			&sosmed.Sosmed{},
		}

		err := app.RunMigrations(&db, models)

		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		utils.SuccessCLIMessage("Migrations succeed!")
	},
}

var UndoMigrate = &cobra.Command{
	Use:   "db:migrate:undo",
	Short: "Undo database migration command!",
	Run: func(cmd *cobra.Command, args []string) {
		database.DBInit()

		utils.WarningCLIMessage("Running undo migrations ... \n")

		var db = *database.DB

		models := []interface{}{
			&user.User{},
			&role.Role{},
			&portofolio.Portofolio{},
			&experience.Experience{},
			&bio.Bio{},
			&sosmed.Sosmed{},
		}

		err := app.RunUndoMigrations(&db, models)

		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		utils.SuccessCLIMessage("Undo migrations succeed!")
	},
}
