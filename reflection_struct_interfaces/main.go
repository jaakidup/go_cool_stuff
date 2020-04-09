package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var intSlice []int
	// var strSlice []string
	// fmt.Println(reflect.ValueOf(intSlice).Type())
	// fmt.Println(reflect.ValueOf(strSlice).Kind())

	random := struct {
		Name string
	}{
		Name: "Freaky",
	}
	Money(random)

	Tams := Animal{
		Name: "Tammy",
		Abilities: map[string]Abilities{
			"run":  &Run{},
			"walk": &Walk{},
		},
	}
	Money(Tams)
	Tams.Abilities["run"].Do()
	Tams.Abilities["run"].Stop()

	Jaaki := Person{
		Name: "Jaaki",
		Abilities: map[string]Abilities{
			"run":  &Run{},
			"walk": &Walk{},
		}}
	Money(Jaaki)
	Jaaki.Abilities["walk"].Do()
	Jaaki.Abilities["walk"].Stop()

}

// Money ...
func Money(input interface{}) {
	switch reflect.TypeOf(input).Name() {
	case "Person":
		fmt.Println("Type is Person", "Choices to be made here")
	case "Animal":
		fmt.Println("Type is Animal", "Other option on this one")
	default:
		fmt.Println("Type is Default", "Default options exist")
	}
}
