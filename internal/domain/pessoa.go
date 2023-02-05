package domain

import (
	"fmt"
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
	var media int64
	var total int64
	for i := 0; i <= len(p.Enfermidades); i++ {
		total = total + p.Enfermidades[i].Severidade
	}
	media = total / int64(len(p.Enfermidades))
	p.EstadoFisico.Classificaçao = media
	p.atualizarHorario()

}

func (p *Pessoa) atualizarHorario() {
	horario := time.Now()
	p.EstadoFisico.DataAvaliaçao = fmt.Sprintln(horario)
}
