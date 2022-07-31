package migrations

import (
	"fmt"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2/utils"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func TestAutoMigrate(t *testing.T) {
	var err error
	var db *gorm.DB
	if os.Getenv("CI_FORCE_USE_POSTGRES") == "1" {
		config := gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			DisableForeignKeyConstraintWhenMigrating: true,
		}

		dbPort := os.Getenv("CI_DB_PORT")
		if dbPort == "" {
			dbPort = "5432"
		}
		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
			os.Getenv("CI_DB_HOST"),
			dbPort,
			os.Getenv("CI_DB_USER"),
			os.Getenv("CI_DB_PASS"),
			os.Getenv("CI_DB_NAME"),
		)

		db, err = gorm.Open(postgres.Open(dsn), &config)
	} else {
		db, err = gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Silent),
			DisableForeignKeyConstraintWhenMigrating: true,
			NamingStrategy:                           schema.NamingStrategy{SingularTable: true},
		})
	}
	utils.AssertEqual(t, nil, err, "database must be connected")
	utils.AssertEqual(t, false, db == nil, "database connection must be established")

	err = AutoMigrate(db)
	utils.AssertEqual(t, nil, err, "migration must be succeed")

	err = SeedAll(db)
	utils.AssertEqual(t, nil, err, "seeder must be succeed")

	os.Setenv("SIMULATE_SEEDER_FAIL", "sample error")
	err = SeedAll(db)
	utils.AssertEqual(t, true, err != nil, "seeder must be fail")

	sqlDB, _ := db.DB()
	sqlDB.Close()
}
