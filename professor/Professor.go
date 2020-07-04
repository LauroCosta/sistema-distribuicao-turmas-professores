package Professor

import (
	"bufio"
	"fmt"
	"os"
	dp "trabalhoLogica/disciplina"
)

type Professor struct {
	Nome        string
	Disciplinas []*dp.Disciplina
}

func NovoProfessor(nome string, disciplinas []*dp.Disciplina) *Professor {
	p := new(Professor)
	p.Nome = nome
	p.Disciplinas = disciplinas
	return p
}

func (p *Professor) ListarDisciplinas() {

	for _, d := range p.Disciplinas {
		fmt.Println(p.Nome + "_" + d.Nome + "_" + d.Curso + " ")
		if d.Nome == "ppct" || d.Nome == "preCalc" {
			d.ExatementeUma(d)
		}

		for _, dia := range d.Dias {
			for _, hora := range d.Horarios {
				if len(p.Disciplinas) == 1 {
					if dia != "ter" && dia != "qui" && dia != "qua" {
						fmt.Printf(d.Nome + "_" + d.Curso + "_" + dia + "_" + hora + " ")
					}
				} else {
					fmt.Printf(d.Nome + "_" + d.Curso + "_" + dia + "_" + hora + " ")
				}
			}
		}

		fmt.Println("")
		if d.Nome != "ppct" && d.Nome != "preCalc" {
			d.ParearDisciplina()
		}
		d.TratarColisao(p.Disciplinas)
		p.TratarDiasDisciplinas()
		p.TratarFolga()
	}

}

func LerNomesProfessores() []string {

	arquivo, err := os.Open("professor/nomes.txt")
	if err != nil {
		return nil
	}
	defer arquivo.Close()
	var linhas []string
	scanner := bufio.NewScanner(arquivo)
	for scanner.Scan() {
		linhas = append(linhas, scanner.Text())
	}
	return linhas
}

func (p *Professor) TratarFolga() {

	horarios := p.Disciplinas[0].Horarios

	for i, dp := range p.Disciplinas {
		for j, dp2 := range p.Disciplinas {
			for _, hora := range horarios {
				for _, hora2 := range horarios {

					if i != j {
						fmt.Printf("-" + dp.Nome + "_" + dp.Curso + "_seg_" + hora + " ")
						fmt.Printf("-" + dp2.Nome + "_" + dp2.Curso + "_sex_" + hora2)
						fmt.Println("")

						fmt.Printf("-" + dp.Nome + "_" + dp.Curso + "_sex_" + hora + " ")
						fmt.Printf("-" + dp2.Nome + "_" + dp2.Curso + "_seg_" + hora2)
						fmt.Println("")
					}
				}
			}
		}
	}
}

func (p *Professor) TratarDiasDisciplinas() {

	dias := []string{"seg", "qua", "sex"}
	horarios := []string{"h1", "h2", "h3", "h4"}

	for i, dp := range p.Disciplinas {
		for _, h := range horarios {
			fmt.Printf("-" + dp.Nome + "_" + dp.Curso + "_ter_" + h + " ")
			for j, dp2 := range p.Disciplinas {
				for _, dia := range dias {
					for _, h2 := range horarios {
						if i != j {
							fmt.Printf(dp2.Nome + "_" + dp2.Curso + "_" + dia + "_" + h2 + " ")
						}
					}
				}
			}
			fmt.Println("")
		}
	}

	for i, dp := range p.Disciplinas {
		for _, h := range horarios {
			fmt.Printf("-" + dp.Nome + "_" + dp.Curso + "_qui_" + h + " ")
			for j, dp2 := range p.Disciplinas {
				for _, dia := range dias {
					for _, h2 := range horarios {
						if i != j {
							fmt.Printf(dp2.Nome + "_" + dp2.Curso + "_" + dia + "_" + h2 + " ")
						}
					}
				}
			}
			fmt.Println("")
		}
	}
}
