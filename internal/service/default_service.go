package service

import "github.com/sebastiaofortes/layout_go/internal/domain"

//services são camadas que contém a lógica relacionanda a dois ou mais agregados
// no nosso exeplo os agregados são pais e pessoa
type DefaultService struct{
	pessoaR domain.PessoaRepository
	paisR domain.PaisRepository
}

func (d *DefaultService) CalcularExpectativaDeVida(){

}

func (d *DefaultService) CalcularExpectativaDeVidaPorPais(){
	
}
func (d *DefaultService) CalcularExpectativaDeVidaPorIdade(){
	
}

//Recomenda-se criar métodos de build das implemtaçoes das interfaces para verificar se todos os métodos estão de fato sendo implemetados
//nosse função tem como tipo uma interface para que suas impermentações sejam aceitas como objeto de retorno.
func NewDefaultService() EsperançaDeVida {
	pessoaRepo := domain.NewImplementsPessoaRepository()
	paisRepo := domain.NewImplementsPaisRepository()
	return &DefaultService{pessoaR: pessoaRepo, paisR: paisRepo}
}