package main

func main() {
	// emptyFile, err := file.New("/tmp/empty.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// fillerFile, err := file.New("/tmp/file.txt", file.UID(1000), file.Contents("My text for file"))
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(emptyFile, fillerFile)

	assembly := car.NewBuilder().Color(car.RedColor)

	familyCar := assembly.Wheels(car.SportsWheels).TopSpeed(50 * car.MPH).Build()
	familyCar.Drive()

	sportsCar := assembly.Wheels(car.SteelWheels).TopSpeed(150 * car.MPH).Build()
	sportsCar.Drive()
}
