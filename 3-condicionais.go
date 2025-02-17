package main

import "fmt"

func main(){
	nota := 9

	if nota >= 7 {
		fmt.Println("Bom trabalho!")
	}else{
		fmt.Println("Precisa se esforçar mais")
	}

	if nota == 10 {
		fmt.Println("Excelente!!!")
	}else if nota >= 7{
		fmt.Println("Muito bom!")
	}else{
		fmt.Println("Ainda esta ruim")
	}

	goIsGood := true

	if goIsGood {
		fmt.Println("Go é bom")
	}else{
		fmt.Println("Go é ruim")
	}

	goIsSpeed := true

	if goIsGood  && goIsSpeed{
		fmt.Println("Go é bom e rápido")
	}else if goIsGood || goIsSpeed{
		fmt.Println("Go é bom ou rápido")
	}else{
		fmt.Println("Go é ruim e lento")
	}

	var numero int
	fmt.Println("--------------")
	fmt.Println("Digite um valor qualquer")
	fmt.Scanln(&numero)

	if numero % 2 == 0{
		fmt.Println("O número é par")
	}else{
		fmt.Println("O número é impar")
	}


}