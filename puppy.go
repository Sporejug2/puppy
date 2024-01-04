package puppy

import (
	"fmt"

	dog "github.com/Sporejug2/Dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof"
}

func BigBark() string {
	return dog.WhenGrowUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrowUp(Barks())
}

func From11() {
	fmt.Println("im from version 1.1.0")
}

func From12() {
	fmt.Println("Im from version 1.2.0")
}

func From13() {
	fmt.Println("Im from version 1.3.0")
}
