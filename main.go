package main

import (
	"fmt"
	"func/file"
)

func main() {
	emptyFile, err := file.New("/tmp/empty.txt")
	if err != nil {
		panic(err)
	}

	fillerFile, err := file.New("/tmp/file.txt", file.UID(1000), file.Contents("My text for file"))
	if err != nil {
		panic(err)
	}

	fmt.Println(emptyFile, fillerFile)
}
