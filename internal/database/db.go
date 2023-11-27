package database

import (
	"project/config"
	"project/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DataBaseConnect() (*gorm.DB, error) {
	cfg := config.GetConfig()
	dsn := cfg.DBConfig.DB_DSN
	//dsn := "host=postgres user=postgres password=admin dbname=jobportal01 port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	//dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// err = db.Migrator().DropTable(&model.Job{}, &model.User{}, &model.Company{})
	// if err != nil {
	// 	return nil, err
	// }

	err = db.Migrator().AutoMigrate(&model.User{}, &model.Company{}, &model.Job{})
	if err != nil {
		// If there is an error while migrating, log the error message and stop the program
		return nil, err
	}
	return db, nil
}
