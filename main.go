package main

import (
	"strconv"
	data "trabalhoLogica/dados"
	pf "trabalhoLogica/professor"
	sm "trabalhoLogica/semestre"
)

func main() {

	disciplinasProfessor := data.MapearDisciplinas("dados/teste.txt")
	disciplinasSemestre := data.MapearSemestre("dados/teste.txt")

	nomesProf := pf.LerNomesProfessores()
	m := make(map[string]*pf.Professor)
	s := make(map[string]*sm.Semestre)

	for _, nome := range nomesProf {
		m[nome] = pf.NovoProfessor(nome, disciplinasProfessor[nome])
		m[nome].ListarDisciplinas()
	}

	for i := 1; i <= 7; i++ {
		s["cc"+strconv.Itoa(i)] = sm.NovoSemestre("cc"+strconv.Itoa(i), disciplinasSemestre["cc"+strconv.Itoa(i)])
		s["cc"+strconv.Itoa(i)] = sm.NovoSemestre("cces"+strconv.Itoa(i), disciplinasSemestre["cc"+strconv.Itoa(i)])

		s["es"+strconv.Itoa(i)] = sm.NovoSemestre("es"+strconv.Itoa(i), disciplinasSemestre["es"+strconv.Itoa(i)])
		s["es"+strconv.Itoa(i)] = sm.NovoSemestre("es"+strconv.Itoa(i), disciplinasSemestre["es"+strconv.Itoa(i)])

		s["cc"+strconv.Itoa(i)].ListarDisciplinas()
		s["es"+strconv.Itoa(i)].ListarDisciplinas()

	}

}
