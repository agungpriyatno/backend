package configs

import (
	"log"
	"os"

	"github.com/agungpriyatno/olap-backend/app/configs/clients"
	"github.com/agungpriyatno/olap-backend/app/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() {
	var err error
	clients.DATABASE, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       os.Getenv("DATABASE_DSN"), // data source name
		DefaultStringSize:         256,                       // default size for string fields
		DisableDatetimePrecision:  true,                      // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,                      // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,                      // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,                     // auto configure based on currently MySQL version
	}), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	clients.DATABASE.AutoMigrate(
		models.Confidence{},
		models.Hotspot{},
		models.Location{},
		models.Satelite{},
		models.Time{},
	)

}
