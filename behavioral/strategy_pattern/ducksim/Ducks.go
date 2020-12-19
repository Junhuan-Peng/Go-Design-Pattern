package ducksim

import "fmt"

// MallarDuck 绿头鸭
type MallarDuck struct {
	Duck
}

// Display 绿头鸭外观描述
func (duck *MallarDuck) Display() {
	fmt.Println("头部是绿色的")
}

// RedHeadDuck 红头鸭
type RedHeadDuck struct {
	Duck
}

// Display 红头鸭外观描述
func (duck *RedHeadDuck) Display() {
	fmt.Println("头部是红色的")
}

// RubberDuck 橡皮鸭
type RubberDuck struct {
	Duck
}

// Display 橡皮鸭外观描述
func (duck *RubberDuck) Display() {
	fmt.Println("橡皮材质的黄色小鸭")
}

// Quack 橡皮鸭吱吱叫
func (duck *RubberDuck) Quack() {
	fmt.Println("吱吱")
}
