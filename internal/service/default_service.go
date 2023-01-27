package service

import "github.com/sebastiaofortes/layout_go/internal/domain"

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

func NewDefaultService() EsperançaDeVida {
	return &DefaultService{}
}