package dto

type Pais struct {
	Id              int32   `gorm:"primary_key;not null" json:"-"`
	ExpectativaVida int32   `json:"expectativa_de_vida"`
	IDH             float64 `json:"idh"`
	BaseDeCalculo   float64 `json:"bse_de_calculo"`
}
