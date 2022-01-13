package main

import (
	"fmt"
	"func/fabric"
	"func/generator"
)

func main() {
	//// funcOprions
	// Opcns, _ := file.NewFile("C:/Users/77679/Documents/GO/rep-study/myTextFile.txt", file.NewContent("my content in my this file ye boi"), file.NewGID(1000), file.NewUID(1000))
	// fmt.Println(Opcns)

	////builder
	// normBuilder := builder.GetBuilder("normal")

	// director := builder.NewDirector(&normBuilder)
	// normalHouse := director.BuildHouse()

	// iglooBuilder := builder.GetBuilder("igloo")
	// director.SetBuilder(iglooBuilder)
	// iglooHouse := director.BuildHouse()

	// fmt.Println(normalHouse.DoorType)
	// fmt.Println(normalHouse.WindowType)
	// fmt.Println(normalHouse.Floor)

	// fmt.Println(iglooHouse.DoorType)
	// fmt.Println(iglooHouse.WindowType)
	// fmt.Println(iglooHouse.Floor)

	// // for i := 0; i < 5; i++ {
	// // 	go singleton.GetInstance()
	// // }

	// // fmt.Scanln()

	// ak47, _ := fabric.GetGun("ak47")
	// musket, _ := fabric.GetGun("musket")

	// printDetails(ak47)
	// printDetails(musket)

	////strategy
	// lfu := &strategy.Lfu{}
	// cache := strategy.InitCache(lfu)

	// cache.Add("a", "1")
	// cache.Add("b", "2")
	// cache.Add("c", "3")

	// lru := &strategy.Lru{}
	// cache.SetEvicionAlgo(lru)

	// cache.Add("d", "4")

	// fifo := &strategy.Fifo{}
	// cache.SetEvicionAlgo(fifo)

	// cache.Add("e", "5")

	////spec
	// shirtItem := spectator.NewItem("Nike Shirt")

	// observerFirst := &spectator.Customer{Id: "abc@gmail.com"}
	// observerSecond := &spectator.Customer{Id: "xyz@gmail.com"}

	// shirtItem.Register(observerFirst)
	// shirtItem.Register(observerSecond)

	// shirtItem.UpdateAvailability()

	////Generator
	fmt.Println("Нет бутылок пива на стене")

	for i := range generator.Count(1, 99) {
		fmt.Println("Передай эту, поставь ее,", i, "бутылок пива на стене")
	}

	fmt.Println(100, "бутылок пива на стене")

}

func printDetails(g fabric.IGun) {
	fmt.Printf("Gun: %s", g.GetName())
	fmt.Println()
	fmt.Printf("Power: %d", g.GetPower())
	fmt.Println()
}
