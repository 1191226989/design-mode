package strategy

import "fmt"

type Context struct {
	Strategy
}

type Strategy interface {
	Do()
}

type Strategy1 struct {
}

func (s1 *Strategy1) Do() {
	fmt.Println("策略1执行了...")
}

type Strategy2 struct {
}

func (s2 *Strategy2) Do() {
	fmt.Println("策略2执行了...")
}
