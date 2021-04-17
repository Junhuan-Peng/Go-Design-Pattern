package main

import (
	"go_design_pattern/behavioral/strategy_pattern/ducksim"
)

func main() {
	aMallarDuck := ducksim.MallarDuck{}
	aMallarDuck.SetFlyBehavior(&ducksim.FlyWithWings{})
	aMallarDuck.Swim()
	aMallarDuck.Fly()

	aRubberDuck := ducksim.RubberDuck{}
	aRubberDuck.SetFlyBehavior(&ducksim.NoFly{})
	aRubberDuck.Quack()
	aRubberDuck.Fly()
}
