package main

import (
	"fmt"
	"time"
)

//Um dos pontos fortes de go é a facilidade para usar concorrência e a sua leveza
//A concorrência é chamada de gorountines e elas são, de forma simples,
//um mecanismo para que instruções sejam executas de forma simultâneas
//isso significa que uma instrução não precisaria esperar a outra acabar para que ela seja executada

//em go as goroutines são criadas usando o comando "go"

/*
"Goroutines são unidades leves de concorrência que são executadas em uma única thread de execução
ou em várias dependendo de como o runtime Go decide gerenciá-las.""
*/

//Se o programa principal terminar antes das goroutines, elas serão encerradas.
//Por isso, o time.Sleep no main ajuda a visualizar a execução concorrente.


func contadorNumeros(){
	for numero := 0; numero < 10; numero++ {
		fmt.Printf("%d ", numero)
		time.Sleep(time.Millisecond * 100)
	}
}

func contadorLetras(){
	for letra := 'a'; letra < 'j'; letra++ {
		fmt.Printf("%c ", letra)
		time.Sleep(time.Millisecond * 100)
	}
}

func main(){
	fmt.Println("EXECUÇÃO DE FORMA TRADICIONAL (SEQUENCIAL)")
	contadorNumeros()
	contadorLetras()
	//função de letra só foi executada após o fim da função de numeros

	fmt.Printf("\n\n")

	fmt.Println("EXECUÇÃO DE FORMA CONCORRENTE (PARALELA)")
	go contadorNumeros()
	contadorLetras()

	//numeros e letras foram embaralhados pois ambas as funções estavam sendo executadas ao mesmo tempo

	//--------------CHANELs------------------
	//São uma forma de fazer com que goroutines diferentes se comuniquem
	//uma boa pratica são criar 2 chanels
	//um para a primeira goroutine escreva e a outra receba
	//e o segundo canal para que a primeira goroutine receba e a outra escreva a mensagem

	//--------SEM BUFFER
	//um goroutine que enviar uma mensagem no canal, fica travada enquanto sua mensagem não é lida
	fmt.Println("\n\n USANDO CANAIS PARA COMUNICAÇÃO DAS GOROUTINES SEM BUFFER")
	chanelNumero := make(chan int)

	go contadorNumerosChanel(chanelNumero)

	for v := range chanelNumero{
		fmt.Printf("Lido com o channel: %d\n", v)
		time.Sleep(time.Millisecond * 180)
	}
	//podemos perceber que a função de numeros nunca escreve um segundo valor antes de um valor já ser lido
	//por isso, o resultado fica em sequencia 

	//--------COM BUFFER
	//aumenta a capacidade das goroutines de trabalharem enquanto suas mensagens nao sao lidas
	//exemplo: buffer de 5, enquanto não tiver 5 mensagens enviadas a goroutine não ficara travada
	fmt.Println("\n\n USANDO CANAIS PARA COMUNICAÇÃO DAS GOROUTINES SEM BUFFER")
	chanelNumeroBuffer := make(chan int, 5)

	go contadorNumerosChanel(chanelNumeroBuffer)

	for v := range chanelNumeroBuffer{
		fmt.Printf("Lido com o channel: %d\n", v)
		time.Sleep(time.Millisecond * 300)
	}
	//PODE SER REPARADO EM COMO A GOROUTINE DE ESCRITA NAO ESPEROU QUE AS MENSAGENS FOSSEM LIDAS
	//TERMINANDO DE EXECUTAR MAIS RAPIDO E SÓ DEPOIS TODAS AS MENSAGENS FORAM LIDAS PELA OUTRA GOROUTINE
}

//passado um canal por parametro
// <- significa que a função só pode escrever no chanel e nao receber
func contadorNumerosChanel(n chan <- int){
	for i := 0; i < 10; i++ {
		n <- i //valor passado para o chanel
		fmt.Printf("Escrito no channel: %d\n", i)
		time.Sleep(time.Millisecond * 150)
	}
	close(n)
}
