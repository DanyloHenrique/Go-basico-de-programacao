package main

import "fmt"

func main(){
	funcMap()
}

//MAP (CHAVE, VALOR)
func funcMap(){
	fmt.Println("=======MAP=======")
	idade := map[string]int{} //criando um map de string e inteiro
	
	//inserindo dados no map
	idade["Danylo"] = 23
	idade["Emily"] = 24

	fmt.Println(idade)//printando todo o map
	fmt.Println(idade["Danylo"])//printando apenas o valor da chave Danylo
	fmt.Println(idade["Quel"]) //nao existe

	idade["Raquel"] = 12
	
	delete(idade, "Raquel")
	fmt.Println(idade)

	//iniciando um map com valores
	elementoPokemon := map[string]string{
		"Bulbasauro": "Planta",
		"Charmander": "Fogo",
		"Squirtle": "Água",
	}

	fmt.Println(elementoPokemon)

	verificaExistenciaChave(idade, "Clhothilde")
	verificaExistenciaChave(idade, "Emily")
}

func verificaExistenciaChave(idade map[string]int, chave string){
	//VERIFICAR A EXISTENCIA DE UMA CHAVE
	//Sintaxe
	//valor, existe := myMap[chave]
	// ok := idade["Raquel"]; verifica se existe, 
	// se sim armazena o valor da chave na variavel recebeValor
	// e ok recebe um booleano de acordo com a existencia da chave
	if recebeValor, ok := idade[chave]; ok{
		fmt.Println("Pessoa existe no map. Idade: ", recebeValor, ok)
	}else{
		fmt.Println("Pessoa NÃO existe no map")
	}
}