package main

import "fmt"

//ponteiros são variaveis que armazenam um endereço da memória de outra variavel,
//e não um valor em si (como um int), 
// permitindo que o programa acesse ou modifique o conteúdo dessa variável indiretamente

//"Os ponteiros são usados para alocar memória dinamicamente,
//passar argumentos para funções por referência,
//criar estruturas de dados complexas, manipular cadeias de caracteres
//e para acessar recursos de hardware diretamente."

//Pontos positivos: eficiencia de memoria, flexibilidade, passagem por referencia
//e acesso a memoria do hardware

//de maneira abstrata
//uma variável é como uma caixa que armazena um valor
//um ponteiro é como se armazenassa um codigo (endereço)
//capaz de acessar e modificador o valor do dado que esta na caixa

//glossario
// & é usado para acessar o endereço da memoria da variavel
// * na criação de ponteiro ou na declaração de uma função indica que vai ser recebido 
// um endereço da memoria
// já * em um ponteiro significa que vai ser acessado o valor da variavel que o ponteiro esta apontando

func main(){

	//==========PONTEIROS=========================================
	var a int = 10 //variavel a que recebe o valor 10
	fmt.Println("Valor de a:", a)//acessando o valor de a
	fmt.Println("Endereço de a:", &a)//acessando o endereço de a (atraves do &)

	//ponteiro é uma variável que vai receber o endereço da memoria (indicado pelo *) de um int
	var ponteiro *int 
	//ponteiro recebendo o endereço da variável a
	ponteiro = &a
	fmt.Println("Valor de ponteiro:", *ponteiro) //o valor do ponteiro passa a ser o proprio valor de a
	fmt.Println("Endereço de ponteiro:", ponteiro)//o endereço da memoria passa a ser o mesmo endereço de a

	//alterando o valor de A atraves do ponteiro
	//"Essa operação de colocar o operador * antes do nome do ponteiro é chamada de “desreferência de ponteiro”"
	*ponteiro = 20
	fmt.Println("Valor de a após alteração pelo *ponteiro:", a)//acessando o valor de a


	//==========PONTEIROS E FUNÇÕES=========================================
	//cada variavel só existe dentro do seu próprio escopo
	//ao passar um valor para uma função, a variavel não é passada, apenas o valor que era daquela variavel
	//com isso, qualquer alteração no valor feita na outra função, 
	//não alterara o valor da variavel na primeira função
	

	var numero = 22
	fmt.Println("numero: ", numero)
	incrementaUm(numero)
	fmt.Println("numero após função de incrementar um:", numero)//(sem alteração)
	//o valor apenas foi incrementado dentro da função incrementaUm,
	//nao resultando em alteração na função original

	//AGORA PASSANDO O ENDEREÇO DA MEMORIA
	//Passa o endereço da memoria é chamado de "passagem por referencia"
	incrementaUmReferencia(&numero)//passando o enderço da memoria
	fmt.Println("numero após função de incrementar um POR REFERENCIA:", numero)//(valor alterado)

}


func incrementaUm(valor int){
	valor = valor + 1
}

//* para indicar que vai ser recebido o endereço de uma memoria
func incrementaUmReferencia(valor *int){
	*valor = *valor + 1
}