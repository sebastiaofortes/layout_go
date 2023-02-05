package dto

type Pais struct {
	ExpectativaVida int32   `json:"expectativa_de_vida"`
	IDH             float64 `json:"idh"`
	baseDeCalculo   float64 `json:"bse_de_calculo"`
}
