package domain

type Pais struct {
	ExpectativaVida float32 `json:"expectativa_vida"`
	IDH             float32 `json:"IDH"`
	BaseDeCalculo 	float32 `json:"base-calculo"`
}

func (p *Pais) CalularExpectativaDeVida() {
	resultado := p.IDH * p.BaseDeCalculo
	p.ExpectativaVida = resultado
}