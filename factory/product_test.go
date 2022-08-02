package factory

import (
	"testing"
)

// 在factory目录执行测试命令 go test -v
// === RUN   TestProduct_Create
//     product_test.go:10: 产品1的名称：iPhone
//     product_test.go:14: 产品2的名称：Motorola
// --- PASS: TestProduct_Create (0.00s)
// === RUN   TestProductFactory_Create
//     product_test.go:22: 产品1的名称：Sumsung
//     product_test.go:26: 产品2的名称：Realme
// --- PASS: TestProductFactory_Create (0.00s)
// PASS
// ok      design-mode/factory     0.002s

func TestProduct_Create(t *testing.T) {
	product1 := &Product1{}
	product1.SetName("iPhone")
	t.Log(product1.GetName())

	product2 := &Product2{}
	product2.SetName("Motorola")
	t.Log(product2.GetName())
}

func TestProductFactory_Create(t *testing.T) {
	productFactory := &ProductFactory{}

	product1 := productFactory.Create(productType1)
	product1.SetName("Sumsung")
	t.Log(product1.GetName())

	product2 := productFactory.Create(productType2)
	product2.SetName("Realme")
	t.Log(product2.GetName())
}
