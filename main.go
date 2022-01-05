package main

import (
	"fmt"
	"func/builder"
)

func main() {
	// Opcns, _ := file.NewFile("C:/Users/77679/Documents/GO/rep-study/myTextFile.txt", file.NewContent("my content in my this file ye boi"), file.NewGID(1000), file.NewUID(1000))
	// fmt.Println(Opcns)

	normBuilder := builder.GetBuilder("normal")

	director := builder.NewDirector(&normBuilder)
	normalHouse := director.BuildHouse()

	iglooBuilder := builder.GetBuilder("igloo")
	director.SetBuilder(iglooBuilder)
	iglooHouse := director.BuildHouse()

	fmt.Println(normalHouse.DoorType)
	fmt.Println(normalHouse.WindowType)
	fmt.Println(normalHouse.Floor)

	fmt.Println(iglooHouse.DoorType)
	fmt.Println(iglooHouse.WindowType)
	fmt.Println(iglooHouse.Floor)

	// for i := 0; i < 5; i++ {
	// 	go singleton.GetInstance()
	// }

	// fmt.Scanln()

}
