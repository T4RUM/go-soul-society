package main

import "fmt"

// address representa um endereço que poderá ser utilizado dentro de outra struct.
type address struct {
	street string
	city   string
}

// user agrupa os dados que representam um usuário.
// O campo address demonstra como uma struct pode conter outra struct.
type user struct {
	name    string
	age     int
	address address
}

func main() {
	// =========================================================================
	// DECLARANDO E PREENCHENDO UMA STRUCT CAMPO A CAMPO
	// =========================================================================

	// Ao declarar uma variável do tipo user sem informar valores,
	// todos os campos começam com o valor zero de seus respectivos tipos.
	var myProfile user

	myProfile.name = "Wesley"
	myProfile.age = 28

	fmt.Println("=== Struct preenchida campo a campo ===")
	fmt.Printf("%+v\n", myProfile)

	// Cada variável do tipo user possui sua própria cópia dos campos.
	var otherProfile user
	otherProfile.name = "Elias"
	otherProfile.age = 23

	fmt.Println("\n=== Outra variável do mesmo tipo ===")
	fmt.Printf("%+v\n", otherProfile)

	// =========================================================================
	// LITERAIS DE STRUCT
	// =========================================================================

	// Uma struct pode ser criada informando os valores pela posição dos campos.
	// Nesse formato, a ordem precisa ser exatamente a mesma da declaração da struct.
	profileByPosition := user{
		"Carlos",
		30,
		address{},
	}

	fmt.Println("\n=== Literal por posição ===")
	fmt.Printf("%+v\n", profileByPosition)

	// A forma com nomes de campos costuma ser mais legível.
	// Campos não informados recebem o valor zero de seu tipo.
	profileByFields := user{
		name: "Mariana",
	}

	fmt.Println("\n=== Literal com nomes de campos ===")
	fmt.Printf("%+v\n", profileByFields)

	// =========================================================================
	// STRUCT ADDRESS
	// =========================================================================

	var myAddress address
	myAddress.street = "Massachusetts Ave NW"
	myAddress.city = "Washington"

	fmt.Println("\n=== Struct de endereço ===")
	fmt.Printf("%+v\n", myAddress)

	// =========================================================================
	// COMO ANINHAR STRUCTS
	// =========================================================================

	// Como user possui um campo chamado address do tipo address,
	// podemos preencher os campos da struct interna usando acesso encadeado.
	var profileWithAddress user
	profileWithAddress.name = "Wesley"
	profileWithAddress.age = 28
	profileWithAddress.address.street = "Massachusetts Ave NW"
	profileWithAddress.address.city = "Washington"

	fmt.Println("\n=== Struct aninhada preenchida campo a campo ===")
	fmt.Printf("%+v\n", profileWithAddress)
	fmt.Println("Cidade:", profileWithAddress.address.city)

	// Também podemos construir a struct externa e a struct interna
	// ao mesmo tempo utilizando literais com nomes de campos.
	completeProfile := user{
		name: "Elias",
		age:  23,
		address: address{
			street: "Avenida Brasil",
			city:   "Washington",
		},
	}

	fmt.Println("\n=== Struct aninhada criada com literal ===")
	fmt.Printf("%+v\n", completeProfile)
	fmt.Println("Rua:", completeProfile.address.street)
	fmt.Println("Cidade:", completeProfile.address.city)

	// Para alterar um campo da struct interna, usamos o mesmo caminho encadeado.
	completeProfile.address.city = "Votorantim"

	fmt.Println("\n=== Alterando um campo da struct interna ===")
	fmt.Println("Nova cidade:", completeProfile.address.city)
}
