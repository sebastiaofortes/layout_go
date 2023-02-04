package service

type EsperançaDeVida interface {
	CalcularExpectativaDeVidaPorPais(int32) float32
	CalcularExpectativaDeVidaPorIdade(int32) float32
}
