package main

import (
	"fmt"
)

// Em Go a unica estrutura de repetição é o for, mas ela é flexive~
// e pode representar todas as outras

func main(){

	//estrutura tradicional
	//for inicio, condição, pós {}
	for i := 0; i < 10; i++{
		fmt.Println(i)
	}

	//for como se fosse while
	// for condição
	var numero = 0
	for numero <= 5{
		fmt.Println("for como'While' - ", numero)
		numero++
	}

	i := 5
	var fatorial = 1
	for i != 0 {
		fatorial = fatorial * i
		if i != 1{
			fmt.Printf("%v * ", i)
		}else{
			fmt.Printf("%v", i)
		}
		i--
	}
	fmt.Println(" = ", fatorial)

	// FOR COMO FOREACH
	//interar sobre coleções
	//for index, valor := range coleção

	pokemons := []string{"Pikachu", "Charizard", "Eevee"}

	for index, value := range pokemons{
		fmt.Printf("%d - %v\n", index+1, value)
	}

	//index pode ser ocultado, ou entao o proprio valor pode ser ocultado
	for _, value := range pokemons{
		fmt.Println(value)
	}

	//for como do while,
	//executado pelo menos uma vez
	var numero2 = 1
	for ok := true; ok; ok = (numero2 % 2 != 0) {
		fmt.Println("Digite um valor par")
		fmt.Scanln(&numero2)
	}

	
	//for aninhado
	//tabuada de 2 e 3
	for i := 2; i <= 3; i++ {
		j := 1
		for j <=10 {
			fmt.Printf("%v * %v = %v\n", i, j, i*j)
			j++
		}

	}
}