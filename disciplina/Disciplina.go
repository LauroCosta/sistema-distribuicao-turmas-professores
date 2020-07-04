package Disciplina

import (
	"fmt"
)

type Disciplina struct {
	Nome     string
	Curso    string
	Dias     []string
	Horarios []string
}

func NovaDisciplina(nome string, curso string) *Disciplina {
	d := new(Disciplina)
	d.Nome = nome
	d.Curso = curso
	d.Dias = []string{"seg", "ter", "qua", "qui", "sex"}
	d.Horarios = []string{"h1", "h2", "h3", "h4"}

	return d
}

func (d *Disciplina) ParearDisciplina() {

	for i, dia := range d.Dias {
		for _, hora := range d.Horarios {
			if i < 2 {
				fmt.Printf("-" + d.Nome + "_" + d.Curso + "_" + dia + "_" + hora + " ")
				fmt.Printf(d.Nome + "_" + d.Curso + "_" + d.Dias[i+2] + "_" + hora + " ")
			} else if i > 2 {
				fmt.Printf("-" + d.Nome + "_" + d.Curso + "_" + dia + "_" + hora + " ")
				fmt.Printf(d.Nome + "_" + d.Curso + "_" + d.Dias[i-2] + "_" + hora + " ")
			} else {
				fmt.Printf("-" + d.Nome + "_" + d.Curso + "_" + dia + "_" + hora + " ")
				fmt.Printf(d.Nome + "_" + d.Curso + "_" + d.Dias[i-2] + "_" + hora + " ")
				fmt.Printf(d.Nome + "_" + d.Curso + "_" + d.Dias[i+2] + "_" + hora + " ")
			}
			fmt.Println("")
		}
	}

}

func (d *Disciplina) TratarColisao(disciplinas []*Disciplina) {

	for i, dp := range disciplinas {
		for j, dp2 := range disciplinas {
			for _, dia := range d.Dias {
				for _, hora := range d.Horarios {
					if i != j {
						fmt.Printf("-" + dp.Nome + "_" + dp.Curso + "_" + dia + "_" + hora + " ")
						fmt.Printf("-" + dp2.Nome + "_" + dp2.Curso + "_" + dia + "_" + hora)
						fmt.Println("")
					}
				}
			}
		}
	}
}

func (d *Disciplina) ExatementeUma(disciplina *Disciplina) {

	for _, d1 := range d.Dias {
		for _, h1 := range d.Horarios {
			for _, d2 := range d.Dias {
				for _, h2 := range d.Horarios {
					if d1 != d2 || h1 != h2 {
						fmt.Printf("-" + disciplina.Nome + "_" + disciplina.Curso + "_" + d1 + "_" + h1 + " ")
						fmt.Printf("-" + disciplina.Nome + "_" + disciplina.Curso + "_" + d2 + "_" + h2 + " ")
						fmt.Println("")
					}
				}
			}
		}
	}
}
