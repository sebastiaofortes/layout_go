package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sebastiaofortes/layout_go/internal/platform/dto"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var StorageDB *sql.DB

func init() {
	dataSource := "root:@tcp(localhost)/mercado_fresco_db?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	StorageDB, err = sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	// test connection
	if err = StorageDB.Ping(); err != nil {
		panic(err)
	}

	// set db options
	StorageDB.SetConnMaxLifetime(time.Minute * 1)
	StorageDB.SetMaxOpenConns(10)
	StorageDB.SetMaxIdleConns(10)

	log.Println("Db connected")
}

func Get() *gorm.DB {
	dataSource := "root:@tcp(localhost)/mercado_fresco_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{})
	if err != nil {
		panic(err)
	}

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
	return db
}
