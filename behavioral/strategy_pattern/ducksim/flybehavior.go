package ducksim

import "fmt"

// FlyBehavior 飞行表现接口
type FlyBehavior interface {
	Fly()
}

// NoFly 无需飞行
type NoFly struct {
}

func (f *NoFly) Fly() {
	fmt.Println("不用飞~")
}

// FlyWithWings 用翅膀飞的行为
type FlyWithWings struct {
}

func (f *FlyWithWings) Fly() {
	fmt.Println("用翅膀飞！")
}
