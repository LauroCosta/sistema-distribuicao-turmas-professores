package dados

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	dp "trabalhoLogica/disciplina"
)

func MapearDisciplinas(caminho string) map[string][]*dp.Disciplina {

	arquivo, err := lerTexto(caminho)
	if err != nil {
		fmt.Println("erro ao ler arquivo")
	}
	m := make(map[string][]*dp.Disciplina)
	re := regexp.MustCompile(`[A-Za-z]*`)
	for _, linha := range arquivo {

		result := re.FindAllString(linha, -1)
		d := dp.NovaDisciplina(result[1], result[2])

		if len(m[result[0]]) == 0 {
			m[result[0]] = append(m[result[0]], d)
		} else {
			m[result[0]] = append(m[result[0]], d)
		}
	}
	return m
}

func MapearSemestre(caminho string) map[string][]*dp.Disciplina {

	arquivo, err := lerTexto(caminho)
	if err != nil {
		fmt.Println("erro ao ler arquivo")
	}
	m := make(map[string][]*dp.Disciplina)
	re := regexp.MustCompile(`[A-Za-z0-9]*`)
	for _, linha := range arquivo {

		result := re.FindAllString(linha, -1)
		if len(result) == 4 {
			d := dp.NovaDisciplina(result[1], result[2])
			if result[2] == "cc" {
				if len(m["cc"+result[3]]) == 0 {
					m["cc"+result[3]] = append(m["cc"+result[3]], d)
				} else {
					m["cc"+result[3]] = append(m["cc"+result[3]], d)
				}
			}
			if result[2] == "es" {
				if len(m["es"+result[3]]) == 0 {
					m["es"+result[3]] = append(m["es"+result[3]], d)
				} else {
					m["es"+result[3]] = append(m["es"+result[3]], d)
				}
			}
		}
		if len(result) == 5 {
			d := dp.NovaDisciplina(result[1], result[2])
			if result[2] == "cces" {
				if len(m["cc"+result[3]]) == 0 {
					m["cc"+result[3]] = append(m["cc"+result[3]], d)
				} else {
					m["cc"+result[3]] = append(m["cc"+result[3]], d)
				}
				if len(m["es"+result[4]]) == 0 {
					m["es"+result[4]] = append(m["es"+result[4]], d)
				} else {
					m["es"+result[4]] = append(m["es"+result[4]], d)
				}
			}
		}
	}
	return m
}

func lerTexto(caminhoDoArquivo string) ([]string, error) {
	arquivo, err := os.Open(caminhoDoArquivo)
	if err != nil {
		return nil, err
	}
	defer arquivo.Close()

	var linhas []string
	scanner := bufio.NewScanner(arquivo)
	for scanner.Scan() {
		linhas = append(linhas, scanner.Text())
	}
	return linhas, scanner.Err()
}
