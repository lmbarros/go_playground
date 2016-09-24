// Subtyping with interfaces. Maybe a clumsy example, but shows the mechanisms.
// This build on subtyping_with_structs.
// A good read: http://spf13.com/post/is-go-object-oriented/

package main

import "fmt"

// NamedCreature is any living creature that has a name. Incidentaly, it is also
// the main change between this example and the subtyping_with_structs one.
type NamedCreature interface {
	Name() string
}

// Animal is an animal. The supertype, so to speak.
type Animal struct {
	TheName string
}

// Name returns the animal name. This is new in relation with
// subtyping_with_structs: we need our Animal to implement the NamedCreature
// interface.
func (a *Animal) Name() string {
	return a.TheName
}

// Describe describes the animal.
func (a *Animal) Describe() {
	fmt.Printf("The animal is called %v.\n", a.Name())
}

// TalkingAnimal is the subtype of Animal. Notice that we don't have to manually
// implement the NamedCreature interface here, as the implementation is
// "inherited" from the Animal supertype.
type TalkingAnimal struct {
	Sound  string
	Animal // An anonymous field; this establishes the "is-a" relationship
}

// Talk makes a TalkingAnimal talk.
func (ta *TalkingAnimal) Talk() {
	fmt.Printf("%v says: \"%s\".\n", ta.Name(), ta.Sound)
}

// Describe overrides the Describe method of the supertype.
func (ta *TalkingAnimal) Describe() {
	fmt.Printf("The animal is called %v and it can talk!\n", ta.Name())
}

// Wash washes a NamedCreature. Now our Wash method does not take a concrete
// type as parameter, but an interface instead. And since both Animal and
// TalkingAnimal implement this interface, we'll be able to wash any animal we
// want. Even a cat.
func Wash(a NamedCreature) {
	fmt.Printf("%v was washed and is now clean.\n", a.Name())
}

// main is the entry point, as you should know.
func main() {
	// Use the supertype
	gi := &Animal{TheName: "Gi, the Giraffe"}
	gi.Describe()

	// Use the subtype
	mimi := &TalkingAnimal{}
	mimi.TheName = "Mimi, the kitten"
	mimi.Sound = "Meow"

	mimi.Describe()
	mimi.Talk()

	// The supertype method is still available for the subtype
	mimi.Animal.Describe()

	// Now, let's try some animal washing
	Wash(gi)
	Wash(mimi) // Now it works! We interfaces we get real subtyping!
}
