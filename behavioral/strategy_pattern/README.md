# Go实现设计模式——策略模式
- [Go实现设计模式——策略模式](#go实现设计模式策略模式)
  - [策略模式](#策略模式)
  - [模拟鸭子应用（DuckSimulation）](#模拟鸭子应用ducksimulation)
    - [初始项目](#初始项目)
    - [新的需求——让鸭子飞起来](#新的需求让鸭子飞起来)
    - [我们遇到了新的问题——一只会飞的橡皮鸭](#我们遇到了新的问题一只会飞的橡皮鸭)
      - [解决方案一](#解决方案一)
      - [解决方案二](#解决方案二)
    - [给我们的小鸭子装上飞行插件](#给我们的小鸭子装上飞行插件)
- [总结](#总结)
## 策略模式
在生活中，我们经常会处于需要进行某种策略选择的场景——例如,购物节有不同的打折方案供我们选择、出行旅游也有多种交通方式等。一种策略即是一种“方案”，同一类策略都具有相同类型的输入和输出。在程序设计过程中，我们也会遇到需要进行策略选择的情况，比如棋盘游戏算法设计中，针对不同的难度选用不同类型的算法等。那么如何来实现策略的灵活变更呢？
策略模式通过定义算法族的方式，分别封装不同的算法，让它们之间可以互相替换，此模式让算法的变化独立于使用算法的客户。

## 模拟鸭子应用（DuckSimulation）
### 初始项目
我们来实现一个鸭子模拟的应用，鸭子会游泳、会呱呱叫。依据面向对象的设计方法，我们暂时设计出了下面的产品。

![初始类图](../../out/uml/behavioral/strategy_pattern/duckInit/Initial%20Duck%20Class%20Diagram.png)
<details>
<pre>
@startuml Initial Duck Class Diagram
abstract class Duck{
    +Swin()
    +Quack()
    {abstract} +Display()
}

class MallarDuck extends Duck{
    +Display()
}

class RedHeadDuck extends Duck{
    +Display()
}
@enduml
</pre>
</details>

### 新的需求——让鸭子飞起来
现在我们需要给鸭子们添加新的功能，让鸭子飞起来！
最简单的方法便是给父类`Duck`添加`Fly`方法。

![带Fly的Duck](../../out/uml/behavioral/strategy_pattern/duckWithFlyFunction/Duck%20Class%20With%20Fly.png)
<details>
<pre>
@startuml Duck Class With Fly
abstract class Duck{
+Swin()
+Quack()
+Fly()
{abstract} +Display()
}
class MallarDuck extends Duck{
+Display()
}

class RedHeadDuck extends Duck{
+Display()
}
@enduml
</pre>
</details>
很好，至少目前看来我们实现了这个需求，而且只增加了少量的代码，同时我们没有额外修改已经存在的代码，干得漂亮！

### 我们遇到了新的问题——一只会飞的橡皮鸭
我们又遇到了问题——当我们增加了一只橡皮鸭（RubberDuck）的时候，它飞了起来！我们可不能让橡皮鸭飞起来。

![会飞的橡皮鸭](../../out/uml/behavioral/strategy_pattern/RubberDuckCanFly/Rubber%20Duck%20Can%20Fly.png)
<details>
<pre>
@startuml Rubber Duck Can Fly
abstract class Duck{
    +Swin()
    +Quack()
    +Fly()
    {abstract} +Display()
}

class MallarDuck extends Duck{
    +Display() 
}

class RedHeadDuck extends Duck{
    +Display()
}

class RubberDuck extends Duck{
    +Quack()
    +Display()
}
@enduml
</pre>
</details>

```go
aRubberDuck := &RubberDuck{}
aRubberDuck.Fly() // it works!
```
#### 解决方案一
利用**继承**的特性，我们尝试让橡皮鸭重写（`override`）`Duck`的`Fly`方法，来让它老老实实呆在地上。

![不会飞的橡皮鸭](../../out/uml/behavioral/strategy_pattern/RubberDuckCannotFly/Rubber%20Duck%20Cannot%20Fly.png)
<details>
<pre>
@startuml Rubber Duck Cannot Fly
abstract class Duck{
    +Swin()
    +Quack()
    +Fly()
    {abstract} +Display()
}

class MallarDuck extends Duck{
    +Display() 
}

class RedHeadDuck extends Duck{
    +Display()
}

class RubberDuck extends Duck{
    +Quack()
    +Display()
    +Fly()
}
@enduml
</pre>
</details>

橡皮鸭总算是不会自己飞起来了！
可是仔细想想，如果我们有新的不会飞的小鸭子们，大家都要重写一遍`Fly`方法吗？会飞的姿势五花八门，不会飞的灵魂千篇一律。明明大家都是一样的不会飞！这样重复的代码可太多了。况且调用一个“没有用”的方法可真让人摸不着头脑。

#### 解决方案二

利用**接口**，让需要飞的鸭子们各自实现自己的飞行姿势！

![实现了飞行接口的鸭子们](../../out/uml/behavioral/strategy_pattern/DucksWithFlyInterface/Ducks%20With%20Fly%20Interface.png)

<details>
<pre>
@startuml Ducks With Fly Interface

interface Flyable{
    Fly()
}

abstract class Duck{
    +Swin()
    +Quack()
    {abstract} +Display()
}

class MallarDuck extends Duck implements Flyable {
    +Display()
}

class RedHeadDuck extends Duck implements Flyable{
    +Display()
}

class RubberDuck extends Duck{
    +Quack()
    +Display()
}
@enduml
</pre>
</details>

Oh! 我们避免了重写大量`Fly`方法——现在只需要大量的`实现`……

### 给我们的小鸭子装上飞行插件
两种解决方案都不太完美。我的意思是，既然有很多类不需要`Fly`，同时，剩下的类里面，有的需要同样形式的`Fly`，有的又需要不同形式的`Fly`，那我们何不把`Fly`抽象出来呢？不是让它作为一个低级的方法，而是把它变成一个emmm**插件**？这样我就可以方便**定制**类的行为了！不是吗？

![鸭子的策略模式](../../out/uml/behavioral/strategy_pattern/DucksWithFlyStrategy/Duck%20With%20Fly%20Strategy.png)
<details>
<pre>
@startuml Duck With Fly Strategy

interface FlyBehavior{
    Fly()
}

abstract class Duck{
    -FlyBehavior doFly
    +Swin()
    +Quack()
    +Fly()
    +setFlyBehavior(Flyable)
    {abstract} +Display()
}

Duck::doFly -l-> FlyBehavior


class MallarDuck extends Duck {
    +Display()
}

class RedHeadDuck extends Duck {
    +Display()
}

class RubberDuck extends Duck {
    +Quack()
    +Display()
}

class NoFly implements FlyBehavior{
    +Fly()
}
class FlyWithWings implements FlyBehavior{
    +Fly()
}

@enduml
</pre>
</details>

这样我们就可以方便定制我们的飞行方式了！而且不需要为了飞行方式的变动而修改小鸭子们了！换言之，我们将鸭子的飞行行为的具体实现托管给了`FlyBehavior`，我们成功的分开的会变化和不会变化的部分。现在我们甚至可以在程序运行过程中对鸭子们的飞行表现进行配置！
```go
type FlyBehavior interface {
	Fly()
}
// NoFly 无需飞行
type NoFly struct {
}

func (f *NoFly) Fly() {
	fmt.Println("不用飞~")
}


type Duck struct {
	doFly FlyBehavior
}
func (duck *Duck) Fly() {
	// 将飞行行为托管给飞行策略
	duck.doFly.Fly()
}

func (duck *Duck) SetFlyBehavior(flyBehavior FlyBehavior) {
	duck.doFly = flyBehavior
}
``` 

```go 
aMallarDuck := ducksim.MallarDuck{}
aMallarDuck.SetFlyBehavior(&ducksim.FlyWithWings{})
aMallarDuck.Swim()
aMallarDuck.Fly()

aRubberDuck := ducksim.RubberDuck{}
aRubberDuck.SetFlyBehavior(&ducksim.NoFly{})
aRubberDuck.Quack()
aRubberDuck.Fly()
```
程序输出（go build main.go）
```shell
Duck Swiming
用翅膀飞！
吱吱
不用飞~
```

# 总结
通过策略模式，我们将飞行这一行为从方法提升到了类。将飞行的具体行为托管给了`FlyBehavior`接口，这样，我们就可以专心控制鸭子和飞行这两件事了。结合策略模式的定义——通过定义算法族的方式，分别封装不同的算法，让它们之间可以互相替换。在这里，`FlyBehavior`便是我们定义的算法族。
策略模式很好的体现了“**面向接口编程**”，避免类和一个具体实现高度耦合。同时也符合**多用组合少用继承**的设计原则。