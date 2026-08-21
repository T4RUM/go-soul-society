package main

import "fmt"

// person representa os dados básicos de uma pessoa.
type person struct {
	name     string
	age      int
	weight   int
	lastName string
}

// student representa um estudante.
//
// O campo `person` está embutido dentro da struct `student`.
//
// Isso é chamado de struct embedding.
//
// Diferente de um campo comum, não declaramos um nome para o campo.
// Informamos apenas o tipo:
//
//	person
//
// Com isso, os campos de `person` ficam promovidos e podem ser
// acessados diretamente através de uma variável do tipo student.
//
// Importante:
// Go NÃO possui herança de classes.
// Esse comportamento é baseado em composição e embedding.
type student struct {
	person
	course string
}

func main() {
	// Criando uma pessoa utilizando um literal de struct.
	person1 := person{
		name:     "John",
		age:      20,
		weight:   40,
		lastName: "Jim",
	}

	// Criando um estudante.
	//
	// Como student possui person embutida, precisamos inicializar
	// a struct person dentro do literal de student.
	person2 := student{
		person: person{
			name:     "Jack",
			age:      20,
			weight:   40,
			lastName: "Jim",
		},
		course: "Harvard Data Science",
	}

	// Exibe todos os dados de person1.
	fmt.Println(person1)

	// Acessando um campo declarado diretamente em student.
	fmt.Println(person2.course)

	// Como person está embutida em student, seus campos são promovidos.
	// Por isso podemos acessar `name` diretamente através de person2.
	fmt.Println(person2.name)

	// Também podemos acessar explicitamente através da struct embutida.
	fmt.Println(person2.person.name)
}
