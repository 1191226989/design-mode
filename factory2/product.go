package factory2

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

// 实现一个抽象工厂
type ProductFactory interface {
	Create() Product
}

// 实现一个具体的工厂：产品1的工厂
type Product01Factory struct {
}

func (p1f *Product01Factory) Create() Product {
	return &Product1{}
}

// 实现一个具体的工厂：产品2的工厂
type Product02Factory struct {
}

func (p1f *Product02Factory) Create() Product {
	return &Product2{}
}
