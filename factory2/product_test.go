package factory2

import (
	"fmt"
	"testing"
)

func TestProductFactory_Create(t *testing.T) {
	product01Factory := &Product01Factory{}
	p1 := product01Factory.Create()
	p1.SetName("iPhone")

	fmt.Println(p1.GetName())
}
