package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {

	mostrarResult("saida.txt")

}

func mostrarResult(caminho string) {

	arquivo, err := lerTexto(caminho)
	if err != nil {
		fmt.Println("erro ao ler arquivo")
	}
	re := regexp.MustCompile(`[A-Za-z0-9]*`)

	profAnterior := ""
	for _, linha := range arquivo {
		r := re.FindAllString(linha, -1)
		if len(r) == 3 {

			if r[0] == profAnterior {
				fmt.Println(buscarDisciplinas(r[1], r[2]))

			} else {
				fmt.Println(r[0])
				fmt.Println(buscarDisciplinas(r[1], r[2]))
			}

			profAnterior = r[0]
		}

	}

}

func buscarDisciplinas(disciplina string, curso string) string {

	arquivo, err := lerTexto("saida.txt")
	if err != nil {
		fmt.Println("erro ao ler arquivo")
	}
	re := regexp.MustCompile(`[A-Za-z0-9]*`)

	anterior := ""
	cursoAntes := ""
	temp := ""

	for _, linha := range arquivo {
		r := re.FindAllString(linha, -1)
		if len(r) != 0 && len(r) == 4 {
			if disciplina == r[0] && curso == r[1] {

				if r[0] == "ppct" || r[0] == "preCalc" {
					return "  " + r[0] + " " + r[1] + " " + r[2] + " " + r[3]
				} else if anterior == r[0] && cursoAntes == r[1] {

					return temp + " " + r[2] + " " + r[3]

				} else {

					temp = "  " + r[0] + " " + r[1] + " " + r[2]
					anterior = r[0]
					cursoAntes = r[1]

				}

			}
		}
	}
	return ""
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
