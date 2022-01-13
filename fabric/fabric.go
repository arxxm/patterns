package fabric

import "fmt"

type IGun interface {
	setName(name string)
	setPower(power int)
	GetName() string
	GetPower() int
}

type gun struct {
	name  string
	power int
}

func (g *gun) setName(name string) {
	g.name = name
}

func (g *gun) setPower(power int) {
	g.power = power
}

func (g *gun) GetName() string {
	return g.name
}

func (g *gun) GetPower() int {
	return g.power
}

type ak47 struct {
	gun
}

func newAk47() IGun {
	return &ak47{
		gun: gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}

type musket struct {
	gun
}

func newMusket() IGun {
	return &musket{
		gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}

func GetGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	}

	if gunType == "musket" {
		return newMusket(), nil
	}

	return nil, fmt.Errorf("Wrong gun type passed")
}
