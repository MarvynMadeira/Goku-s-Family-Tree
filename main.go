package main

import (
	"fmt"
)

type familyMember struct {
	name     string
	wife     *familyMember
	children []*familyMember
	father   *familyMember
	mother   *familyMember
}

func newMember(name string, father, mother *familyMember) *familyMember {
	return &familyMember{
		name:   name,
		father: father,
		mother: mother,
	}
}

func (s *familyMember) newChildren(newPerson *familyMember) {
	s.children = append(s.children, newPerson)
	newPerson.father = s
}

func (s *familyMember) show_me_thisFamily(tree int) {
	for i := 0; i < tree; i++ {
		fmt.Print(" ") // " " generates a space for each level
	}
	fmt.Println(s.name)

	for _, child := range s.children {
		child.show_me_thisFamily(tree + 1)
	}
}

func main() {
	bardock := newMember("Bardock", nil, nil)
	gine := newMember("Gine", nil, nil)
	goku := newMember("Goku", bardock, gine)
	raditz := newMember("Raditz", bardock, gine)

	bardock.newChildren(goku)
	bardock.newChildren(raditz)

	chichi := newMember("Chichi", nil, nil)
	gohan := newMember("Gohan", goku, chichi)
	goten := newMember("Goten", goku, chichi)

	goku.newChildren(gohan)
	goku.newChildren(goten)

	videl := newMember("Videl", nil, nil)
	pan := newMember("Pan", gohan, videl)

	gohan.newChildren(pan)

	fmt.Println("Goku's Family Tree:")
	bardock.show_me_thisFamily(0)
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
