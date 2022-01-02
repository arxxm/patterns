package main

import (
	"fmt"
)

func main() {
	// Opcns, _ := file.NewFile("C:/Users/77679/Documents/GO/rep-study/myTextFile.txt", file.NewContent("my content in my this file ye boi"), file.NewGID(1000), file.NewUID(1000))
	// fmt.Println(Opcns)

	// normalBuilder := builder.GetBuilder("normal")
	// iglooBuilder := builder.GetBuilder("igloo")

	// director := builder.NewDirector(normalBuilder)
	// normalHouse := director.BuildHouse()

	// fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorType)
	// fmt.Printf("Normal House Window Type: %s\n", normalHouse.windowType)
	// fmt.Printf("Normal House Num Floor: %d\n", normalHouse.floor)

	// director.SetBuilder(iglooBuilder)
	// iglooHouse := director.BuildHouse()

	// fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.doorType)
	// fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.windowType)
	// fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.floor)

	for i := 0; i < 5; i++ {
		go singleton.getInstance()
	}

	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()

}
