package main

import "fmt"

func main(){
	funcArrays()
	funcSlice()
}

//ARRAYS
func funcArrays(){
	fmt.Println("=======ARRAYS=======")

	var array [4]string

	array[0] = "Primeiro"
	fmt.Println("Primeira posição", array[0])
	fmt.Println("Array inteiro", array)

	array[1] = "Segundo"
	array[2] = "terceiro"
	fmt.Println(array)
}

//SLICES (Não precisa determinar o tamanho da lista e usa função para operações)
func funcSlice(){
	fmt.Println("=======SLICES=======")

	var mySlice []string

	mySlice = append(mySlice, "Cachorro", "Gato")
	fmt.Println("tamanho: ", len(mySlice))

	fmt.Println(mySlice)
	fmt.Println("primeira posição", mySlice[0])

	//operacoes
	mySlice = append(mySlice, "Papagaio", "Calopsita")
	//slice[x:x-1]
	fmt.Println("itens 2 e 3: ", mySlice[1:3])

	//remover o ultimo item
	mySlice = mySlice[:3]
	fmt.Println(mySlice)

}
