package dto

import "github.com/sebastiaofortes/layout_go/internal/domain"

type Enfermidade struct {
	Id         int32  `gorm:"primary_key;not null" json:"-"`
	Nome       string `gorm:"primary_key;not null" json:"nome"`
	Severidade int64  `json:"severidade"`
}

func (e Enfermidade) ToDomain() domain.Enfermidade {
	return domain.Enfermidade{
		Nome: e.Nome,
		Severidade: e.Severidade, 
	}
}

type EstadoFisico struct {
	Id            int32  `gorm:"primary_key;not null" json:"-"`
	Classificaçao int64  `json:"classificacao"`
	DataAvaliaçao string `json:"data_avaliacao"`
}

func (e EstadoFisico) ToDomain() domain.EstadoFisico {
	return domain.EstadoFisico{
		Classificaçao: e.Classificaçao,
		DataAvaliaçao: e.DataAvaliaçao, 
	}
}

type Pessoa struct {
	Id           int32         `gorm:"primary_key;not null" json:"-"`
	Nome         string        `json:"nome"`
	Idade        int32         `json:"idade"`
	Enfermidades []Enfermidade `gorm:"ForeignKey:Id;AssociationForeignKey:Id" json:"enfermidades"`
	EstadoFisico EstadoFisico  `gorm:"ForeignKey:Id;AssociationForeignKey:Id" json:"estado_fisico"`
	Pais         int64         `json:"pais"`
}

func (p Pessoa) ToDomain() domain.Pessoa {
	return domain.Pessoa{
		Nome: p.Nome,
		Idade: p.Idade,
		Enfermidades: []domain.Enfermidade{},
		EstadoFisico: domain.EstadoFisico{},
		Pais: p.Pais,
	}
}