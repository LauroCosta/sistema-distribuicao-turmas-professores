package Semestre

import (
	dp "trabalhoLogica/disciplina"
)

type Semestre struct {
	id          string
	Disciplinas []*dp.Disciplina
}

func NovoSemestre(id string, disciplinas []*dp.Disciplina) *Semestre {
	s := new(Semestre)
	s.id = id
	s.Disciplinas = disciplinas
	return s
}

func (s *Semestre) ListarDisciplinas() {
	for _, d := range s.Disciplinas {
		d.TratarColisao(s.Disciplinas)
	}
}
