package decorator

import "fmt"

// 定义一个抽象的装饰者
type ComponentDecorator interface {
	Component
	Do() // 这个是额外的方法
}

// 实现一个具体的装饰者
type ComponentDecorator1 struct {
	c Component
}

func (cd1 *ComponentDecorator1) Do() {
	fmt.Println("调用装饰方法")
}

func (cd1 *ComponentDecorator1) Operate() {
	cd1.Do()
	cd1.c.Operate()
	cd1.Do()
}
