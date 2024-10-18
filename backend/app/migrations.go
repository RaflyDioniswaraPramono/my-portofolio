package app

import "gorm.io/gorm"

func RunMigrations(db *gorm.DB, models []interface{}) error {
	for _, model := range models {
		err := db.AutoMigrate(&model)

		if err != nil {
			return err
		}
	}

	return nil
}

func RunUndoMigrations(db *gorm.DB, models []interface{}) error {
	for _, model := range models {
		err := db.Migrator().DropTable(&model)

		if err != nil {
			return err
		}
	}

	return nil
}
