package decorator

type pizza interface {
	GetPrice() int
}

type VeggeMania struct {
}

func (p *VeggeMania) GetPrice() int {
	return 15
}

type TomatoTopping struct {
	Pizza pizza
}

func (c *TomatoTopping) GetPrice() int {
	pizzaPrice := c.Pizza.GetPrice()
	return pizzaPrice + 7
}

type CheeseTopping struct {
	Pizza pizza
}

func (c *CheeseTopping) GetPrice() int {
	pizzaPrice := c.Pizza.GetPrice()
	return pizzaPrice + 10
}
