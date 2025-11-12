package main;

import (
	"fmt"
	"errors"
)

func main() {
	// Tipos de dados no Go:
	fmt.Println("--------- INTEIROS ---------");
	// Inteiro: int8, int16, int32 e int64
	var idade int8 = 32;
	fmt.Println("Idade do usuário:", idade);

	var altura int16 = 181;
	fmt.Println("Altura do usuário:", altura);

	var algumNumeroNegativo int64 = -1234567;
	fmt.Println("Alguma número negativo:", algumNumeroNegativo);

	// Caso uma variável só tenha número inteiro positivo, podemos usar o uint:
	// uint8, uint16, uint32 e uint64
	var inteiroPositivoPequeno uint8 = 3;
	fmt.Println("Inteiro positivo pequeno:", inteiroPositivoPequeno);
	var inteiroPositivoGrande uint64 = 123456789;
	fmt.Println("Inteiro positivo grande:", inteiroPositivoGrande);

	//Temos dois apelidos para tipo inteiros: rune -> int32 e byte -> uint8
	var apelidoInteiro32 rune = 324;
	fmt.Println("Apelido para int32:", apelidoInteiro32);
	var apelidoInteiroPositivo8 byte = 33;
	fmt.Println("Apelido para uint8:", apelidoInteiroPositivo8);

	fmt.Println("--------- REAIS ---------");
	//Para os tipos reais temos apenas dois: float32 e float64
	var numeroReal32 float32 = 123.45;
	fmt.Println("Número real de 32:", numeroReal32);

	var numeroReal64 float64 = 123456.99;
	fmt.Println("Número real de 64:", numeroReal64);

	fmt.Println("--------- STRINGS ---------");
	
	var nomeCompleto string = "Ana Maria José";
	fmt.Println("Nome:", nomeCompleto);

	// Não há caractere no Go - podemos usar aspas simples com um caractere que será transformado num número inteiro
	// que é relativo aquele caractere na tabelas ASCII
	caractereASCII := 'A';
	fmt.Println("Caractere ASCII:", caractereASCII);

	fmt.Println("--------- BOOLEAN ---------");
	var casado bool = false;
	fmt.Println("Casado:", casado);
	var namorando bool = true;
	fmt.Println("Solteiro:", namorando);

	//No Go temos o tipo error muito usado nas funções
	fmt.Println("--------- ERROR ---------");
	
	var erro error = errors.New("Erro interno");
	fmt.Println("Erro:", erro);

	fmt.Println("--------- VALOR ZERO ---------");
	// Quando uma variável é declarada, mas não é inicializada o compilador dá um valor padrão de inicialização conforme o tipo:
	var stringZero string;
	var inteiroZero int8;
	var inteiroPositivoZero uint16;
	var floatZero float32;
	var booleanZero bool;
	var errorZero error;

	fmt.Println("String zero:", stringZero);
	fmt.Println("Inteiro zero:", inteiroZero);
	fmt.Println("Inteiro positivo zero:", inteiroPositivoZero);
	fmt.Println("Float zero:", floatZero);
	fmt.Println("Boolean zero", booleanZero);
	fmt.Println("Erro zero", errorZero);

}