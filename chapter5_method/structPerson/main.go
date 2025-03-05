package main

import (
	"fmt"
	"strings"
	"unsafe"
)

type Person struct {
	firstName string
	lastName  string
	name      struct{}
}

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func main() {

	var person1 Person
	person1.firstName = "Chris"
	person1.lastName = "Woodward"
	upPerson(&person1)

	// 第二种
	pers2 := new(Person)
	pers2.firstName = "Chris"
	pers2.lastName = "Woodward"
	(*pers2).lastName = "Woodward"
	upPerson(pers2)
	fmt.Printf("The name of the person is %s %s\n", pers2.lastName, pers2.firstName)

	// 第三种
	pers3 := &Person{firstName: "Chrs", lastName: "Woodward"}
	upPerson(pers3)
	fmt.Printf("The name of the person is %s %s\n", pers3.lastName, pers3.firstName)
	size := unsafe.Sizeof(pers3)
	fmt.Println(size)

}
