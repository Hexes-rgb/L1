package main

import "fmt"

type Car interface { //define Car interface
	Drive()
}

type Horse struct{}

func (horse *Horse) Ride() string { //Horse structure do not implement Car interface
	return "Riding.."
}

type HorseAdapter struct { //Adapter for Horse structure
	horse *Horse
}

func (horseAdapter *HorseAdapter) Drive() string { //HorseAdapter implement Car interface and call Horse method Ride
	return horseAdapter.horse.Ride()
}

func main() {
	horse := &Horse{} //create horse

	horseAdapter := &HorseAdapter{ //adapter
		horse: horse,
	}

	fmt.Println(horseAdapter.Drive()) //call Drive method
}
