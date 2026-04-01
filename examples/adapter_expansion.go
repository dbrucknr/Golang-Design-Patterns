package examples

/// Adapter Pattern Sample
// - Makes incompatible interfaces work together
// Example: Adapts a Turkey interface to a Duck interface

func adapterSample() {
	mallard := &MallardDuck{}

	simulator := &DuckSimulator{duck: mallard}
	simulator.Simulate()

	// Turkey's are not compatible with DuckSimulators
	// - They satisfy a Turkey interface
	// - They must be adapted to satisfy the Duck interface
	//   - Note that Quack and Fly method have been customized
	turkey := &WildTurkey{}
	turkeyAdapter := &TurkeyAdapter{Turkey: turkey} // composition (satisfies Duck interface)
	simulator.duck = turkeyAdapter
	simulator.Simulate()

	// Drone's are also not compatible with DuckSimulators
	// - They satisfy a Drone interface
	// - We use an adapter to make the 2 interfaces compatible
	//   - Note that the Quack and Fly methods have also been overwritten
	drone := &SuperDrone{}
	droneAdapter := &DroneAdapter{Drone: drone} // composition (satisfies Duck interface)
	simulator.duck = droneAdapter
	simulator.Simulate()
}

type DuckSimulator struct {
	duck Duck
}

func (s *DuckSimulator) Simulate() {
	s.duck.Quack()
	s.duck.Fly()
}

type Duck interface {
	Quack()
	Fly()
}

type MallardDuck struct {
}

func (d *MallardDuck) Quack() {
	println("Quack")
}

func (d *MallardDuck) Fly() {
	println("MallardDuck - Fly")
}

type Turkey interface {
	Gobble()
	Fly()
}

type WildTurkey struct {
}

func (t *WildTurkey) Gobble() {
	println("Gobble")
}

func (t *WildTurkey) Fly() {
	println("WildTurkey - Fly")
}

// Satisfies the Duck interface by adapting the Turkey interface
type TurkeyAdapter struct {
	Turkey Turkey
}

func (a *TurkeyAdapter) Quack() {
	a.Turkey.Gobble()
}

func (a *TurkeyAdapter) Fly() {
	for range 5 {
		a.Turkey.Fly()
	}
}

type Drone interface {
	Beep()
	SpinRotors()
	TakeOff()
}

type DroneAdapter struct {
	Drone Drone
}

func (a *DroneAdapter) Quack() {
	a.Drone.Beep()
}

func (a *DroneAdapter) Fly() {
	a.Drone.SpinRotors()
	a.Drone.TakeOff()
}

type SuperDrone struct {
}

func (d *SuperDrone) Beep() {
	println("SuperDrone - Beep")
}

func (d *SuperDrone) SpinRotors() {
	println("SuperDrone - SpinRotors")
}

func (d *SuperDrone) TakeOff() {
	println("SuperDrone - TakeOff")
}
