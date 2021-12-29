package main

import (
	"fmt"
	"func/file"
)

func main() {
	Opcns, _ := file.NewFile("C:/Users/77679/Documents/GO/rep-study/myTextFile.txt", file.NewContent("my content in my this file ye boi"), file.NewGID(1000), file.NewUID(1000))
	fmt.Println(Opcns)
}
