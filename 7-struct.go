package main

import "fmt"

//struct: forma de criar a propria estrutura de dado, personalizada de acordo com a necessidade
//metodo: É basicamente uma função comum, mas que está associada a uma struct específica.
//Isso significa que ele pode acessar e modificar os campos dessa struct,
//diferente de uma função normal que precisaria receber a struct como argumento

type Pokemon struct{
	Nome string
	Tipo string
	Pokedex int
	Peso float32
}

// METODOS
func (p Pokemon) Apresentar(){
	fmt.Println("Ola, me chamo", p.Nome, "- sou do tipo ", p.Tipo)
}

func main(){
	//formas de criar dados
	p1 := Pokemon{"Squirtle", "Água", 007, 19.8}
	p2 := Pokemon{Nome: "Charmander", Tipo: "Fogo", Pokedex: 004, Peso: 18.7}//passando os campos
	//com valor faltando (Apenas passando os campos)
	p3 := Pokemon{Nome: "Bulbasaur", Tipo: "Planta", Pokedex: 001} 

	p4 := Pokemon{Nome: "Pikachu", Tipo: "Elétrico", Pokedex: 025, Peso: 13.2}
	p5 := Pokemon{Nome: "Growlithe", Tipo: "Fogo", Pokedex: 58, Peso: 41.9}
	p6 := Pokemon{Nome: "Pidgey", Tipo: "Voador", Pokedex: 016, Peso: 4.0}


	//imprimindo todo os dados
	fmt.Println(p1, p2, p3)

	//imprimindo apenas os nomes
	fmt.Println(p1.Nome, p2.Nome, p3.Nome)

	//acessando o metodo
	p1.Apresentar()

	//CRIANDO UMA LISTA DO TIPO STRUCT
	pokemons := []Pokemon{}//criando a lista
	pokemons = append(pokemons, p1, p2, p3)//adicionando as variaveis ja criadas á lista
	fmt.Println(pokemons)

	//CRIANDO UM MAP DO TIPO STRUCT

	//usando uma lista ja existente
	pokemonsMap := map[string][]Pokemon{} //pokemonsMap = um map de chave (string) e valor (lista do tipo Pokemon)
	pokemonsMap["time1"] = pokemons
	fmt.Println("time 1:", pokemonsMap["time1"])

	//criando dados ao iniciar o map
	pokemonsMap2 := map[string][]Pokemon{
		//criando as instancias ao mesmo tempo
		"time2": {
			{Nome: "Rattata", Tipo: "Normal", Pokedex: 19, Peso: 7.7},
			{Nome: "Zubat", Tipo: "Venenoso", Pokedex: 041, Peso: 16.5},
			{Nome: "Meowth", Tipo: "Normal", Pokedex: 052, Peso: 9.3},
		},

		//utilizando variaveis ja criadas
		"time3": {p4, p5, p6},
	}
	fmt.Println(pokemonsMap2)

	//CRIANDO SUBSTRUCTS
	type PokemonEvoluido struct{
		Pokemon
		PokemonAnterior string
	}

	p7 := PokemonEvoluido{
		Pokemon: Pokemon{Nome: "Wartortle", Tipo: "Água", Pokedex: 8, Peso: 49.6},                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             
		PokemonAnterior: "Squirtle",
	}

	fmt.Println(p7)
	fmt.Println(p7.Pokemon.Nome)
}

