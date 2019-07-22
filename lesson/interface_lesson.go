package main

import "fmt"

/**
定义一个接口
*/
type Animal interface {
	Fly()
	Run()
}

type Animal2 interface {
	Fly()
}

/**
定义一个结构体(类)
*/
type Bird struct {
}

/**
实现接口对两个方法
*/
func (b Bird) Fly() {
	fmt.Println("Flying")
}

func (b Bird) Run() {
	fmt.Println("Running")
}

func main() {

	//定义一个接口animal
	var animal Animal
	//实例化一个类
	bird := new(Bird)
	//1.将对象实例赋值给接口，该类同时也实现类该接口中定义对所有方法，即该类实现了该接口
	animal = bird

	animal.Run()
	animal.Fly()

	//2.接口赋值给接口，只能大接口赋值给小接口，被包含
	var animal2 Animal2
	animal2 = animal
	animal2.Fly()

	//接口类型查询interface
	var v interface{}
	v = 1.23
	if v, ok := v.(float64); ok {
		fmt.Printf("%v is float32", v)
	}

}
