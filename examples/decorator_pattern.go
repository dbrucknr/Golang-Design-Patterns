package examples

func decoratorSample() {
	darkRoast := &DarkRoastCoffee{}
	houseBlend := &HouseBlendCoffee{}
	espresso := &Espresso{}

	// Blanket apply all options onto the concrete beverages
	// - These could be conditional orders in a more sophisticated system
	milkDarkRoast := &MilkDecorator{darkRoast}
	milkHouseBlend := &MilkDecorator{houseBlend}
	milkEspresso := &MilkDecorator{espresso}

	mochaMilkDarkRoast := &MochaDecorator{milkDarkRoast}
	mochaMilkHouseBlend := &MochaDecorator{milkHouseBlend}
	mochaMilkEspresso := &MochaDecorator{milkEspresso}

	whipMochaMilkDarkRoast := &WhipDecorator{mochaMilkDarkRoast}
	whipMochaMilkHouseBlend := &WhipDecorator{mochaMilkHouseBlend}
	whipMochaMilkEspresso := &WhipDecorator{mochaMilkEspresso}

	println(whipMochaMilkDarkRoast.Description(), whipMochaMilkDarkRoast.Price())
	println(whipMochaMilkHouseBlend.Description(), whipMochaMilkHouseBlend.Price())
	println(whipMochaMilkEspresso.Description(), whipMochaMilkEspresso.Price())

	doubleMochaDarkRoast := &DarkRoastCoffee{}                                // 2.0
	firstMochaDarkRoast := &MochaDecorator{doubleMochaDarkRoast}              // + 1.0
	secondMochaDarkRoast := &MochaDecorator{firstMochaDarkRoast}              // + 1.0
	println(secondMochaDarkRoast.Description(), secondMochaDarkRoast.Price()) // Expect 4.0
}

type Beverage interface {
	Description() string
	Price() float64
}

// Concrete Beverages
type DarkRoastCoffee struct {
}

func (d *DarkRoastCoffee) Description() string {
	return "DarkRoastCoffee"
}

func (d *DarkRoastCoffee) Price() float64 {
	return 2.0
}

type HouseBlendCoffee struct {
}

func (h *HouseBlendCoffee) Description() string {
	return "HouseBlendCoffee"
}

func (h *HouseBlendCoffee) Price() float64 {
	return 2.5
}

type Espresso struct {
}

func (e *Espresso) Description() string {
	return "Espresso"
}

func (e *Espresso) Price() float64 {
	return 1.5
}

// Decorators (wrappers around beverages)
// - These are used to compose additional features onto beverages
type MilkDecorator struct {
	Beverage
}

func (m *MilkDecorator) Description() string {
	return m.Beverage.Description() + ", Milk"
}

func (m *MilkDecorator) Price() float64 {
	return m.Beverage.Price() + 0.5
}

type MochaDecorator struct {
	Beverage
}

func (m *MochaDecorator) Description() string {
	return m.Beverage.Description() + ", Mocha"
}

func (m *MochaDecorator) Price() float64 {
	return m.Beverage.Price() + 1.0
}

type WhipDecorator struct {
	Beverage
}

func (w *WhipDecorator) Description() string {
	return w.Beverage.Description() + ", Whip"
}

func (w *WhipDecorator) Price() float64 {
	return w.Beverage.Price() + 0.5
}
