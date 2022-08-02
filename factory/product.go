package factory

// 实现一个抽象的产品
type Product interface {
	SetName(name string)
	GetName() string
}

// 实现具体的产品1
type Product1 struct {
	name string
}

func (p1 *Product1) SetName(name string) {
	p1.name = name
}

func (p1 *Product1) GetName() string {
	return "产品1的名称：" + p1.name
}

// 实现具体的产品2
type Product2 struct {
	name string
}

func (p2 *Product2) SetName(name string) {
	p2.name = name
}

func (p2 *Product2) GetName() string {
	return "产品2的名称：" + p2.name
}

// 实现简单工厂类
type ProductType int

const (
	productType1 ProductType = iota
	productType2
)

type ProductFactory struct{}

func (pf *ProductFactory) Create(pt ProductType) Product {
	if pt == productType1 {
		return &Product1{}
	}
	if pt == productType2 {
		return &Product2{}
	}
	return nil
}
