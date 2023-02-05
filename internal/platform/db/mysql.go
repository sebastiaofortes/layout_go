package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/sebastiaofortes/layout_go/internal/platform/dto"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Storege *gorm.DB

func Get() *gorm.DB {
	dataSource := "root:@tcp(localhost)/ddd_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	Storege = db

	db.Migrator().DropTable(
		&dto.Enfermidade{},
		&dto.EstadoFisico{},
		&dto.Pais{},
		&dto.Pessoa{},
	)

	db.AutoMigrate(
		&dto.Enfermidade{},
		&dto.EstadoFisico{},
		&dto.Pais{},
		&dto.Pessoa{},
	)

	db.Save(&dto.Pessoa{Id: 1, Nome: "João", Idade: 33, Pais: 1})
	db.Save(&dto.EstadoFisico{Id: 1, Classificaçao: 2, DataAvaliaçao: "10/10/2020"})
	db.Save(&dto.Enfermidade{Id: 1, Nome: "Teste", Severidade: 2})
	db.Save(&dto.Pais{Id: 1, ExpectativaVida: 60, IDH: 3.5, BaseDeCalculo: 0.4})
	return db
}
