package service

import (
	"github.com/sebastiaofortes/layout_go/internal/domain"
	"github.com/sebastiaofortes/layout_go/internal/platform/repository"
)

// services são camadas que contém a lógica relacionanda a dois ou mais agregados
// no nosso exeplo os agregados são pais e pessoa
type DefaultService struct {
	pessoaR domain.PessoaRepository
	paisR   domain.PaisRepository
}

func (d *DefaultService) calcularExpectativaDeVida(pessoa domain.Pessoa, pais domain.Pais) float32 {
	var total float32
	pessoa.CalcularEstadoFisico()
	total = pais.ExpectativaVida * float32(pessoa.EstadoFisico.Classificaçao)
	return total
}

func (d *DefaultService) CalcularExpectativaDeVidaPorPais(p int32) float32 {
	Pais, _ := d.paisR.GetPais(p)
	Listapessoas, _ := d.pessoaR.GetPessoasPorPais(p)
	var media float32
	var total float32
	for _, v := range Listapessoas {
		total = total + d.calcularExpectativaDeVida(v, Pais)
	}
	media = total / float32(len(Listapessoas))
	return media
}
func (d *DefaultService) CalcularExpectativaDeVidaPorIdade(i int32) float32 {
	Listapessoas, _ := d.pessoaR.GetPessoasPorPais(i)
	var media float32
	var total float32
	for _, v := range Listapessoas {
		Pais, _ := d.paisR.GetPais(int32(v.Pais))
		total = total + d.calcularExpectativaDeVida(v, Pais)
	}
	media = total / float32(len(Listapessoas))
	return media
}

// Recomenda-se criar métodos de build das implemtaçoes das interfaces para verificar se todos os métodos estão de fato sendo implemetados
// nosse função tem como tipo uma interface para que suas impermentações sejam aceitas como objeto de retorno.
func NewDefaultService() EsperançaDeVida {
	pessoaRepo := repository.NewImplementsPessoaRepository()
	paisRepo := repository.NewImplementsPaisRepository()
	return &DefaultService{pessoaR: pessoaRepo, paisR: paisRepo}
}
