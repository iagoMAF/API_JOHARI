package database

import (
	"log"

	"github.com/iagoMAF/API_JOHARI/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaDataBase() {
	stringConnection := "host=localhost user=root password=root dbname=root port=5434 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringConnection))

	if err != nil {
		log.Panic("Erro ao conectar com o banco de dados.")
	}

	DB.AutoMigrate(&models.Atleta{}, &models.Admin{})
}
