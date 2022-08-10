package decorator

import "testing"

func TestComponent1_Operate(t *testing.T) {
	c1 := &Component1{}
	c1.Operate()
}

func TestComponentDecorator1_Operate(t *testing.T) {
	cd1 := &ComponentDecorator1{
		c: &Component1{},
	}
	cd1.Operate()
}
