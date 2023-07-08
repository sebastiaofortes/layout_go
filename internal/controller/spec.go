package controller

//go:generate mockery --name=EsperançaDeVida --output=./impl_mocks --outpkg=impl_mocks
type EsperançaDeVida interface {
	CalcularExpectativaDeVidaPorPais(int32) (float32, error)
	CalcularExpectativaDeVidaPorIdade(int32) (float32, error)
}
