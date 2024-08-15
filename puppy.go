package puppy

import (
	"fmt"

	"github.com/tanzeelalam/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func from11() {
	fmt.Println("I'm from version 1.1.0")
}
