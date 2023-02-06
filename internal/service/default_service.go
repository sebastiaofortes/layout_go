package service

import (
	"log"

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
	var result float32
	pais.CalularExpectativaDeVida()
	pessoa.CalcularEstadoFisico()
	result = pais.ExpectativaVida * float32(pessoa.EstadoFisico.Classificaçao)
	log.Println("Resultado = [expectativa de vida] ", pais.ExpectativaVida, "* [estado fisico]", pessoa.EstadoFisico.Classificaçao, " = ", result)
	return result
}

func (d *DefaultService) CalcularExpectativaDeVidaPorPais(p int32) float32 {
	var media float32
	var total float32
	
	Pais, _ := d.paisR.GetPais(p)
	Listapessoas, _ := d.pessoaR.GetPessoasPorPais(p)
	for _, pessoas := range Listapessoas {
		log.Println("------Nome: ", pessoas.Nome, "Idade: ", pessoas.Idade, "------")
		total = total + d.calcularExpectativaDeVida(pessoas, Pais)
	}
	media = total / float32(len(Listapessoas))
	log.Println("------ resultado ", media, " ------")
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
