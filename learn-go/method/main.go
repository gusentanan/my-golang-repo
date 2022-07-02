package main

import (
	"fmt"
	"strings"
)

func main() {

	var hero1 = heroName{"Bruce Wayne", "Batman"}
	hero1.sayhello()

	var name = hero1.getNameAt(2)
	fmt.Println("His hero name is ", name)

	fmt.Println("Hero Name (before changes): ", hero1.heroNickName)
	hero1.changeName("Vigilante")
	fmt.Println("Hero Name (after changes): ", hero1.heroNickName)

	hero1.changeNameByPointer("Hokage")
	fmt.Println("Hero name after changes: ", hero1.heroNickName)
}

func (h heroName) sayhello() {
	fmt.Println("Hello ", h.name)
}

func (h heroName) getNameAt(i int) string {
	return strings.Split(h.name, " ")[i-1]
}

func (h heroName) changeName(nickName string) {
	fmt.Println(" --> on changeName, name change to ", nickName)
	h.heroNickName = nickName
}

func (h *heroName) changeNameByPointer(nickName string) {
	fmt.Println(" --> on changeNameByPointer, Hero Nick Name change to ", nickName)
	h.heroNickName = nickName
}

type heroName struct {
	name         string
	heroNickName string
}
