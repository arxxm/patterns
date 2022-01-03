package builder

type normalBuilder struct{
	windowType string
	doorType string
	floor int
}

type house struct {
    windowType string
    doorType   string
    floor      int
}

type iBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse()
}

func getBuilder(builderType string) iBuilder{
	if iBuilder == "noraml" {
		return &normalBuilder{}
	}

	// if builderType == "igloo" {
	// 	return
	// }

	return nil
}

func newNormalBuilder() *normalBuilder {
	return &normalBuilder{}
}

func (b *normalBuilder) setWindowType() {
	return b.windowType = "Wooden window"
}

func (b *normalBuilder) setDoorType() {
	return b.doorType = "Wooden door"
}

func (b *normalBuilder) setFloor() {
	return b.floor = 2
}

func (b *normalBuilder) getHouse() {
	return house {
		doorType: b.doorType,
		windowType: b.windowType,
		floor: b.floor,
	}
}

