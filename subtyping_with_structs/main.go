// Subtyping with structs. May be useful, but is not real subtyping.
// See also subtyping_with_interfaces, which continues from here.
// Good read: http://spf13.com/post/is-go-object-oriented/
package main

import "fmt"

// Animal is an animal. The supertype, so to speak.
type Animal struct {
	Name string
}

// Describe describes the animal.
func (a *Animal) Describe() {
	fmt.Printf("The animal is called %v.\n", a.Name)
}

// TalkingAnimal is the subtype of Animal.
type TalkingAnimal struct {
	Sound  string
	Animal // An anonymous field; this establishes the "is-a" relationship
}

// Talk makes a TalkingAnimal talk.
func (ta *TalkingAnimal) Talk() {
	fmt.Printf("%v says: \"%s\".\n", ta.Name, ta.Sound)
}

// Describe overrides the Describe method of the supertype.
func (ta *TalkingAnimal) Describe() {
	fmt.Printf("The animal is called %v and it can talk!\n", ta.Name)
}

// Wash washes an Animal.
func Wash(a *Animal) {
	fmt.Printf("%v was washed and is now clean.\n", a.Name)
}

// main is the entry point, as you should know.
func main() {
	// Use the supertype
	gi := &Animal{Name: "Gi, the Giraffe"}
	gi.Describe()

	// Use the subtype
	mimi := &TalkingAnimal{}
	mimi.Name = "Mimi, the kitten"
	mimi.Sound = "Meow"

	mimi.Describe()
	mimi.Talk()

	// The supertype method is still available for the subtype
	mimi.Animal.Describe()

	// Now, let's try some animal washing
	Wash(gi)
	// Wash(mimi) // Doesn't work! This is not real subtyping!
}
