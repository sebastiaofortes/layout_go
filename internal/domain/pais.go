package domain

type Pais struct {
	ExpectativaVida int32
	IDH             float64
	baseDeCalculo 	float64
}

func (p *Pais) CalularExpectativaDeVida() {
	resultado := p.IDH * p.baseDeCalculo
	p.ExpectativaVida = int32(resultado)
}