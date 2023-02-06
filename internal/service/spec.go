package service

type EsperançaDeVida interface {
	CalcularExpectativaDeVidaPorPais(int32) (float32, error)
	CalcularExpectativaDeVidaPorIdade(int32) (float32, error)
}
