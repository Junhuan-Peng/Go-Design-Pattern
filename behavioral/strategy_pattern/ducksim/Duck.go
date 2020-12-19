package ducksim

import (
	"fmt"
)

// DuckInterface 定义鸭子接口
type DuckInterface interface {
	Quack()
	Swim()
	Display()
	Fly()
}

// Duck 鸭子主体，实现鸭子接口的默认方法,作为所有具体鸭子的父类
type Duck struct {
	doFly FlyBehavior
}

// Quack 鸭子呱呱叫
func (duck *Duck) Quack() {
	fmt.Println("呱呱~")
}

// Swim 鸭子游泳
func (duck *Duck) Swim() {
	fmt.Println("Duck Swiming")
}

// Display 鸭子外观
func (duck *Duck) Display() {

}

// Fly 鸭子执行飞行动作
func (duck *Duck) Fly() {
	// 将飞行行为托管给飞行策略
	duck.doFly.Fly()
}

// SetFlyBehavior 配置鸭子飞行策略
func (duck *Duck) SetFlyBehavior(flyBehavior FlyBehavior) {
	duck.doFly = flyBehavior
}
