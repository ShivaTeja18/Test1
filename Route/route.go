package Route

import (
	"EMP_API/Details"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Connect() {
	const DNS = "host = 'localhost' port = '5432' dbname = 'API' user = 'postgres' password = 'Shiva@205101' sslmode = 'prefer'"
	Database, err := gorm.Open(postgres.Open(DNS), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}
	if err == nil {
		log.Print("Connected")
	}
	_ = Database.AutoMigrate(&Details.EMP{})

	DB = Database
}

func Initialize() {

	R.GET("/fetch", Fetching)
	R.GET("/fetch/:id", Fbyid)
	R.POST("/create", Creating)
	R.PUT("/change/:id", Updating)
	//r.DELETE("/delete", deleting)
	R.Run(":8000")
}
