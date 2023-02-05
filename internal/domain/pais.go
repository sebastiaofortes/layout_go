package domain

type Pais struct {
	ExpectativaVida float32 `json:"expectativa_vida"`
	IDH             float32 `json:"IDH"`
	baseDeCalculo 	float32 `json:"base-calculo"`
}

func (p *Pais) CalularExpectativaDeVida() {
	resultado := p.IDH * p.baseDeCalculo
	p.ExpectativaVida = resultado
}