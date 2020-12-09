# Go实现设计模式——单例模式
- [Go实现设计模式——单例模式](#go实现设计模式单例模式)
  - [单例模式（Singleton Pattern）](#单例模式singleton-pattern)
    - [懒汉式](#懒汉式)
    - [饿汉式](#饿汉式)
    - [双重检查](#双重检查)
    - [Sync.Once](#synconce)
## 单例模式（Singleton Pattern）

在软件中，个别对象仅仅需要创建一次（实例化一次），例如线程池、日志对象、状态管理器等。如果程序运行期间出现了多个实例，那么往往会造成诸多问题，比如程序的行为异常、资源的浪费以及创建对象所造成的性能影响等。简单的情况下，开发人员之间通过约定来保证对象的独一性，比如静态变量、全局变量等手段来达成目的。但是，在更复杂的情况下，我们应当采用单例模式来强制保证对象的独一性。

开发人员使用语法（比如关键字`new`）来实例化一个类（或结构体），产物是一个对象。单例模式指导我们分隔开“实例化/创建”对象和“获取”对象两步操作——通过设计屏蔽实例化的接口，并提供用户新的获取对象的接口。

### 懒汉式

在面向对象的语言中，可以通过将构造方法声明为**private**来实现第一步，并提供新的方法来返回对象。

比如 Java 实现如下：

```java
class Singleton{
    private static Singleton instance;
    private Singleton(){
        // do something
    }
    public Singleton getInstance(){
        if(instance == null){
            instance = new Singleton();
        }
        return instance;
    }
}
```

Go语言不是严格的面向对象的语言，没有构造方法，通过大小写来控制外部可见性。

```go
type singleton struct{
    // some value
}

var singletonObj *singleton
func GetInstance() *singleton{
    if singletonObj == nil{
        singletonObj = &singleton{
            // assign value
        }
    }
    return singletonObj
}
```

我们实现了一个非常简单的满足单例模式的例子，这种的实现被称之为**懒汉式**，因为对象在第一次被获取的时候才会创建。这还有改进的空间。

我们没有对整个过程添加额外的判断和保障，所以，我们的代码如果运行在多线程下，那么依旧可能出现对象的多次实例化。所以我们需要对整个过程进行加锁来保障线程安全。Java可以通过对方法添加**synchronized**关键字来保障线程安全。Go的实现如下：

```go
var lock sync.Mutex
var singletonObj *singleton
func GetInstance() *singleton{
    lock.Lock()
    defer lock.Unlock()
    if singletonObj == nil{
        singletonObj = &singleton{
            // assign value
        }
    }
    return singletonObj
}
```

这种方式的优点是*保证了线程安全性*，缺点是*在高频率调用时，会反复加锁释锁，对性能造成一定的影响*。

### 饿汉式

相较于懒汉式，饿汉式通过在一开始就创建对象的方式，来避免对对象进行判空的操作，从而避免了线程的不安全性（任何时候进入，对象都是相同状态）。Go的实现如下：

```go
type singleton struct{
    // some value
}

// init global object （valid in package）
var singletonObj *singleton = &singleton{
            // assign value
} 

func GetInstance() *singleton{
    return singletonObj
}
```

饿汉式的优点在于*保证了线程的安全性*，缺点是*对象从程序开始时就一直创建在内存中，持续占有内存*。

### 双重检查

基于线程安全的饿汉式，如果我们能够想办法减少加锁释锁的次数，就能既保证线程安全也能避免对象过早创建带来的内存占用。

反思我们加锁的原因，是因为**创建对象**的步骤应该是独占的，如果多个线程分别检测到对象不存在，那么就会造成对象的多次创建，单例模式失效。我们可以通过增加检查步骤来避免这个问题（相较于加锁释锁，一次比较所增加的性能可以忽略不记）。Go的实现如下

```go
var lock sync.Mutex
var singletonObj *singleton
func GetInstance() *singleton{
    if singletonObj == nil{
        lock.Lock()
        if singletonObj == nil{
            singletonObj = &singleton{
                // assign value
            }
        }
        lock.Unlock()
    }
    return singletonObj
}
```

Java实现如下：

```java
class Singleton{
    private volatile static Singleton instance;
    private Singleton(){
        // do something
    }
    public Singleton getInstance(){
        if(instance == null){
            synchronized(Singleton.class){
                if(instance==null){
                	instance = new Singleton();
                }
            }
        }
        return instance;
    }
}
```

### Sync.Once

Go 的Sync包提供了Once类型，使用其**Do**方法来保证作为参数传入的方法仅仅执行一次，且是线程安全。实现如下：

```go
var once Sync.Once
var singletonObj *singleton
func GetInstance() *singleton{
    once.Do(func(){
        singletonObj = &singleton{
            // assign value
        }
    })
    return singleObj
}
```





