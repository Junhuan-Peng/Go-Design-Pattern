package main

// 鸭子接口
type Duck interface {
	quack()
	fly()
}

//具体的绿头鸭
type MallardDuck struct {
}

func (receiver MallardDuck) quack() {
	println("绿头鸭呱呱呱")
}

func (receiver MallardDuck) fly() {
	println("绿头鸭起飞！！")
}

//火鸡
type Turkey struct {
}

func (receiver Turkey) gobble() {
	println("火鸡咯咯咯")
}

func (receiver Turkey) fly() {
	println("火鸡起飞！！！")
}

//将火鸡适配到鸭子上的适配器
type TurkeyAdaptor struct {
	turkey Turkey
}

func (receiver TurkeyAdaptor) quack() {
	receiver.turkey.gobble()
}
func (receiver TurkeyAdaptor) fly() {
	receiver.turkey.fly()
}

func main() {
	aMallardDuck := MallardDuck{} // 一只绿头鸭
	aMallardDuck.quack()
	aMallardDuck.fly()

	turkey := Turkey{} // 一只火鸡
	turkey.gobble()
	turkey.fly()

	fakeDuck_which_is_real_a_turkey := TurkeyAdaptor{turkey: turkey}

	// 看起来就和鸭子一样了
	fakeDuck_which_is_real_a_turkey.quack()
	fakeDuck_which_is_real_a_turkey.fly()

	//通过Duck接口使用不同的类型
	var aDuck [2]Duck
	aDuck[0] = aMallardDuck
	aDuck[1] = fakeDuck_which_is_real_a_turkey

}
