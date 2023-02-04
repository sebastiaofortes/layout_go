package domain

type Pais struct {
	ExpectativaVida float32
	IDH             float32
	baseDeCalculo 	float32
}

func (p *Pais) CalularExpectativaDeVida() {
	resultado := p.IDH * p.baseDeCalculo
	p.ExpectativaVida = resultado
}