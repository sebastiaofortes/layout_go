package domain

import (
	"fmt"
	"time"
)

type Pessoa struct {
	Nome         string
	Idade        int32
	Enfermidades []Enfermidade
	EstadoFisico EstadoFisico
}

func (p *Pessoa) CalcularEstadoFisico() {
	var media int64
	var total int64
	for i := 0; i <= len(p.Enfermidades); i++ {
		total = total + p.Enfermidades[i].Severidade
	}
	media = total/int64(len(p.Enfermidades))
	p.EstadoFisico.Classificaçao = media
	p.atualizarHorario()

}

func (p *Pessoa) atualizarHorario(){
	horario := time.Now()
	p.EstadoFisico.DataAvaliaçao = fmt.Sprintln(horario)
}
