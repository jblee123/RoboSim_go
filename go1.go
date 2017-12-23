package main

import (
	"fmt"
)

type Behavior interface {
	PrintId()
	Activate()
}

type BehaviorBase struct {
	Id int
}

func (b *BehaviorBase) PrintId() {
	fmt.Printf("Id: %d\n", b.Id)
}

type Behavior1 struct {
	BehaviorBase
}

type Foo1 struct {
	val1 int
}

func (f *Foo1) Bar() {
	fmt.Printf("Foo1.Bar -- val1: %d\n", f.val1)
}

func (f *Foo1) DoIt() {
	fmt.Printf("Foo1.DoIt -- val1: %d\n", f.val1)
	f.Bar()
}

type Foo2 struct {
	Foo1
}

func (f *Foo2) Bar() {
	fmt.Printf("Foo2.Bar -- val1: %d\n", f.val1)
}

// func (f *Foo2) DoIt() {
//     fmt.Printf("Foo2.DoIt -- val1: %d\n", f.val1)
//     f.Bar()
// }

func main() {
	var f1 Foo1
	f1.val1 = 5
	f1.DoIt()

	var f2 Foo2
	f2.val1 = 6
	f2.DoIt()
}
