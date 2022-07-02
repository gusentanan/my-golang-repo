package main

import "fmt"

func main() {
	var newStudent student
	newStudent.name = "Bagus Merta"
	newStudent.grade = 100
	fmt.Println("Name: ", newStudent.name)
	fmt.Println("Grade: ", newStudent.grade)

	var superStudent = specialStudent{}
	superStudent.name = "Bagus"
	superStudent.yourHeroName = "strawbery milk"
	superStudent.specialPower.power = "KamehameHA"
	superStudent.specialPower.weakness = "i hate cockorocos"

	var newVillain = villain{group{"penjahat", 30}, []string{"tell lies"}}
	fmt.Println("villain weakness:", newVillain.weakness)

	var badVillain *villain = &newVillain
	fmt.Println("new villain Group Name: ", newVillain.group.nameGroup) // penjahat
	badVillain.group.nameGroup = "JOKERS"
	fmt.Println("bad villain Group Name: ", badVillain.group.nameGroup) // JOKERS

	// all student
	var allStudents = []person{
		{name: "Budi", age: 23},
		{name: "Joko", age: 23},
		{name: "Popowi", age: 22},
	}
	for _, student8 := range allStudents {
		fmt.Println(student8.name, "age is", student8.age)
	}

	// all student 2
	var allStudents2 = []struct {
		person
		grade int
	}{
		{person: person{"Aji", 21}, grade: 2},
		{person: person{"Bang", 21}, grade: 3},
		{person: person{"Powpowi", 21}, grade: 3},
	}
	for _, student9 := range allStudents2 {
		fmt.Println(student9)
	}

}

type student struct {
	name  string
	grade int
}

type specialStudent struct {
	name         string
	yourHeroName string
	specialPower
}

type specialPower struct {
	power    string
	weakness string
}

type villain struct {
	group struct {
		nameGroup string
		stars     int
	}
	weakness []string
}

type group struct {
	nameGroup string
	stars     int
}

type person struct {
	name string
	age  int
}
