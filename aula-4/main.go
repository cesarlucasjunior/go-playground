//Variáveis em Go

package main

import "fmt"

func main() {
	// Go é uma linguagem de tipagem estática e fortemente tipada, mas podemos inferir o tipo da variável
	// tipo explícito
	var nome string = "César"
	fmt.Println("Nome:", nome)
	// tipo implícito
	var idade = 32
	fmt.Println("Idade:", idade)
	// declaração curta (só pode ser usada dentro de funções) - inferência de tipo
	cidade := "Brasília"
	fmt.Println("Cidade:", cidade)

	var (
		altura float32 = 1.81
		peso  float32 = 75.5
	)

	fmt.Println("Altura:", altura)
	fmt.Println("Peso:", peso)

	corDosOlhos, corDoCabelo := "Castanhos", "Pretos"
	fmt.Println("Cor dos olhos:", corDosOlhos)
	fmt.Println("Cor do cabelo:", corDoCabelo)

	const numeroDePulmoes int = 2
	fmt.Println("Número de pulmões:", numeroDePulmoes)

	//inversao de variáveis
	a, b := 1, 2
	fmt.Println("Antes da troca: a =", a, ", b =", b)
	a, b = b, a
	fmt.Println("Depois da troca: a =", a, ", b =", b)
}