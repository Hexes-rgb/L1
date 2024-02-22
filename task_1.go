package main

import "fmt"

// Human structure with simple methods for example
type Human struct {
	Age  int
	Name string
}

func (h *Human) GetAge() int {
	return h.Age
}

func (h *Human) GetName() string {
	return h.Name
}

// Definition of Action structure
type Action struct {
	Human // Embedding the Human structure
}

// Definition of the Action structure method with using parent structure methods
func (a *Action) IntroduceYourself() {
	fmt.Printf("Hello, my name is %s, i'm %d years old.\n", a.GetName(), a.GetAge())
}

func main() {
	// Creation object of the Action type
	actionPerson := Action{
		Human: Human{Name: "Mary", Age: 24},
	}

	// Call method of Action type object with using parent structure methods
	actionPerson.IntroduceYourself()
}
