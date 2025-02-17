package main

import (
	"fmt"
	"strings"
)

const PI float64 = 3.14;
const NOME string = "DANYLO HENRIQUE";


/* 
**int8:** Inteiro de 8 bits (valores de -128 a 127).
**int16:** Inteiro de 16 bits (valores de -32768 a 32767).
**int32:** Inteiro de 32 bits (valores de -2147483648 a 2147483647).
**int64:** Inteiro de 64 bits (valores de -9223372036854775808 a 9223372036854775807
*/

func main(){
	var texto string = "Minha variavel";
	texto = "novo texto"
	fmt.Println(texto)

	var var2 = 123
	fmt.Println("var2: ", var2)

	var var3 int //inicializando sem passar um valor, e o GO inicializa ela com um valor padrao para cada tipo
	fmt.Println("var3: ", var3)

	var4 := "Variavel shorthand"
	fmt.Println("var4: ", var4)

	var5 := true
	fmt.Println("var5: ", var5)

	fmt.Println("constante PI: ", PI)
	fmt.Println("constante NOME: ", NOME)

	funcInt()
	funcFloat()
	funcBoolean()
	funcString()
}

func funcInt(){
	var idade int8
	//idade = 200 //overflow
	idade = 23
	println("idade: ", idade)
}

func funcFloat(){
	//NÃP PODE FAZER OPERAÇÃO ENTRE FLOAT32 E FLOAT64

	// var teste float32 = 2.211
	var raio float64 = 1.22
	var resultado float64 = PI * raio * raio
	println("area: ", resultado)
}

func funcBoolean(){
	var GoEhLegal bool = true
	var maior bool = 10 > 5
	var menor bool = 10 < 5


	println("Go é bom?: ", GoEhLegal)
	println("10 > 5: ", maior)
	println("10 < 5: ", menor)

}

func funcString(){
	var hello string = "Bom dia"
	var question string = " Tudo bem?"

	var meet = hello + question

	fmt.Println(meet)
	fmt.Println(strings.ToUpper(meet))
	fmt.Println(strings.Contains(meet, "Bom"))

}