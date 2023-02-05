package domain

import (
	"fmt"
	"log"
	"time"
)

type Pessoa struct {
	Nome         string        `json:"nome"`
	Idade        int32         `json:"idade"`
	Enfermidades []Enfermidade `json:"enfermidades"`
	EstadoFisico EstadoFisico  `json:"estado_fisico"`
	Pais         int64         `json:"pais"`
}

func (p *Pessoa) CalcularEstadoFisico() {
	var media float64
	var total float64
	for i := 0; i < len(p.Enfermidades); i++ {
		total = total + p.Enfermidades[i].Severidade
	}
	media = total / float64(len(p.Enfermidades))
	p.EstadoFisico.Classificaçao = media
	p.atualizarHorario()
	log.Println("Estado fisico = media de todas as enfermidades = ", media)

}

func (p *Pessoa) atualizarHorario() {
	horario := time.Now()
	p.EstadoFisico.DataAvaliaçao = fmt.Sprintln(horario)
}
