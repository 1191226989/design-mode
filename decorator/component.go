package decorator

import "fmt"

// 定义一个抽象的组件
type Component interface {
	Operate()
}

// 实现一个具体的组件：组件1
type Component1 struct {
}

func (c1 *Component1) Operate() {
	fmt.Println("c1 operate...")
}
