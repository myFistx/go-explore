package main

import (
	"fmt"
)

/**
一个方法就是一个带有接受者的函数
类型和作用在它上面定义的方法必须在同一个包里面定义
*/
type interge struct {
	value int
}

func (a interge) compare(b interge) bool {
	return a.value < b.value
}

type Point struct {
	px float32
	py float32
}

func (point *Point) setXY(px, py float32) {
	point.px = px
	point.py = py
}

func (point *Point) getXY() (float32, float32) {
	return point.px, point.py
}

/**
继承
*/
type Point3 struct {
	Point
	pz float32
}

type Node struct {
	value       int
	left, right *Node
}

//当接收者是一个nil的时候，会报：panic: runtime error: invalid memory address or nil pointer dereference
func (node *Node) setValue(value int) {
	if node == nil {
		fmt.Println("node is nil")
		return
	}
	node.value = value
}

func createNode(value int) *Node {
	return &Node{value: value}
}

func main() {

	node1 := createNode(1222)
	fmt.Println(*node1)

	node := new(Node)
	node.setValue(222)
	node.left.setValue(33)
	fmt.Println(node)
	fmt.Println("ok")

	a := interge{1}
	b := interge{2}
	fmt.Println(a.compare(b))

	point := new(Point)
	point.setXY(1.23, 4.56)
	x, y := point.getXY()
	fmt.Println(x, y)

	point3 := new(Point3)
	point3.pz = 7.89
	point3.setXY(1.23, 4.56)
	x1, y1 := point.getXY()
	fmt.Println(x1, y1, point3.pz)

}
