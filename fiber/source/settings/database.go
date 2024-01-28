package settings

import (
	"myapp/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB
var DATABASE_URI string = "host=pgbouncer user=postgres password=postgres dbname=django port=6432 sslmode=disable TimeZone=Asia/Bangkok"

func Connect() error {
	var err error

	Database, err = gorm.Open(postgres.Open(DATABASE_URI), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		panic(err)
	}

	Database.AutoMigrate(&entities.Blog{})

	// query blog if not exist add blog 5 rows
	// var count int64
	// Database.Model(&entities.Blog{}).Count(&count)
	// if count == 0 {
	// 	for i := 0; i < 5; i++ {
	// 		Database.Create(&entities.Blog{
	// 			Title:   fmt.Sprint("Blog Title ", i),
	// 			Content: fmt.Sprint("Blog Content ", i),
	// 		})
	// 	}
	// }

	return nil
}
