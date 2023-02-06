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
	db.Save(&dto.EstadoFisico{Id: 1, Classificaçao: 0.9, DataAvaliaçao: "10/10/2020"})
	db.Save(&dto.Enfermidade{Id: 1, Nome: "Teste", Severidade: 0.8})
	db.Save(&dto.Pais{Id: 1, ExpectativaVida: 70, IDH: 0.9, BaseDeCalculo: 90.4})
	
	db.Save(&dto.Pessoa{Id: 2, Nome: "Maria", Idade: 25, Pais: 1})
	db.Save(&dto.EstadoFisico{Id: 2, Classificaçao: 0.9, DataAvaliaçao: "10/10/2020"})
	db.Save(&dto.Enfermidade{Id: 2, Nome: "Teste", Severidade: 0.9})
	db.Save(&dto.Pais{Id: 1, ExpectativaVida: 70, IDH: 0.9, BaseDeCalculo: 90.4})
	return db
}
