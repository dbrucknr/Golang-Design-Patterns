package examples

func pizzaDecoratorSample() {
	var myPizza Pizza = &ThinCrustPizza{}
	myPizza = &CheeseDecorator{pizza: myPizza}
	myPizza = &OliveDecorator{pizza: myPizza}
	myPizza = &PepperoniDecorator{pizza: myPizza}

}

type Pizza interface {
	GetDescription() string
	GetPrice() float64
}

type ThinCrustPizza struct {
}

func (t *ThinCrustPizza) GetDescription() string {
	return "Thin Crust Pizza"
}

func (t *ThinCrustPizza) GetPrice() float64 {
	return 8.99
}

type ThickCrustPizza struct {
}

func (t *ThickCrustPizza) GetDescription() string {
	return "Thick Crust Pizza"
}

func (t *ThickCrustPizza) GetPrice() float64 {
	return 9.99
}

// Decorators
type CheeseDecorator struct {
	pizza Pizza
}

func (c *CheeseDecorator) GetDescription() string {
	return c.pizza.GetDescription() + " with cheese"
}

func (c *CheeseDecorator) GetPrice() float64 {
	return c.pizza.GetPrice() + 1.99
}

type OliveDecorator struct {
	pizza Pizza
}

func (o *OliveDecorator) GetDescription() string {
	return o.pizza.GetDescription() + " with olive"
}

func (o *OliveDecorator) GetPrice() float64 {
	return o.pizza.GetPrice() + 0.99
}

type PepperoniDecorator struct {
	pizza Pizza
}

func (p *PepperoniDecorator) GetDescription() string {
	return p.pizza.GetDescription() + " with pepperoni"
}

func (p *PepperoniDecorator) GetPrice() float64 {
	return p.pizza.GetPrice() + 2.99
}
