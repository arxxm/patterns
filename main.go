package main

import (
	"fmt"
	"func/singleton"
)

func main() {
	// Opcns, _ := file.NewFile("C:/Users/77679/Documents/GO/rep-study/myTextFile.txt", file.NewContent("my content in my this file ye boi"), file.NewGID(1000), file.NewUID(1000))
	// fmt.Println(Opcns)

	normalBuilder = builder.getBuilder("normal")

	for i := 0; i < 5; i++ {
		go singleton.GetInstance()
	}

	fmt.Scanln()

}
