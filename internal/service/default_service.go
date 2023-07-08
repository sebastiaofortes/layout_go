package service

import (
	"log"

	"github.com/sebastiaofortes/layout_go/internal/domain"
)

// services são camadas que contém a lógica relacionanda a dois ou mais agregados
// no nosso exeplo os agregados são pais e pessoa
type DefaultService struct {
	pessoaR PessoaRepository
	paisR   PaisRepository
}

func (d *DefaultService) calcularExpectativaDeVida(pessoa *domain.Pessoa, pais *domain.Pais) float32 {
	var result float32
	pais.CalularExpectativaDeVida()
	pessoa.CalcularEstadoFisico()
	result = pais.ExpectativaVida * float32(pessoa.EstadoFisico.Classificaçao)
	log.Println("Resultado = [expectativa de vida] ", pais.ExpectativaVida, "* [estado fisico]", pessoa.EstadoFisico.Classificaçao, " = ", result)
	return result
}

func (d *DefaultService) CalcularExpectativaDeVidaPorPais(p int32) (float32, error) {
	var media float32
	var total float32

	Pais, err := d.paisR.GetPais(p)
	if err != nil {
		return media, err
	}
	Listapessoas, err := d.pessoaR.GetPessoasPorPais(p)
	if err != nil {
		return media, err
	}
	for _, pessoa := range Listapessoas {
		total = total + d.calcularExpectativaDeVida(&pessoa, &Pais)
		log.Println("------Nome: ", pessoa.Nome, "Idade: ", pessoa.Idade, "------")
	}
	media = total / float32(len(Listapessoas))
	log.Println("------ resultado ", media, " ------")
	return media, nil
}
func (d *DefaultService) CalcularExpectativaDeVidaPorIdade(i int32) (float32, error) {
	var media float32
	var total float32

	Listapessoas, err := d.pessoaR.GetPessoasPorIdade(i)
	if err != nil {
		return media, err
	}
	for _, pessoa := range Listapessoas {
		Pais, err := d.paisR.GetPais(int32(pessoa.Pais))
		if err != nil {
			return media, err
		}
		total = total + d.calcularExpectativaDeVida(&pessoa, &Pais)
		log.Println("------Nome: ", pessoa.Nome, "Idade: ", pessoa.Idade, "------")
	}
	media = total / float32(len(Listapessoas))
	log.Println("------ resultado ", media, " ------")
	return media, nil
}

// Recomenda-se criar métodos de build das implemtaçoes das interfaces para verificar se todos os métodos estão de fato sendo implemetados
// nosse função tem como tipo uma interface para que suas impermentações sejam aceitas como objeto de retorno.
func NewDefaultService(pessoaRepo PessoaRepository, paisRepo PaisRepository) *DefaultService {
	return &DefaultService{pessoaR: pessoaRepo, paisR: paisRepo}
}
