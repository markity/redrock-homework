package main

import "fmt"

// Animal 保存动物相关信息
type Animal struct {
	// ID 动物的唯一识别码
	ID string
	// Type 动物的品种
	Type string
	// Age 动物的年龄
	Age int
	// Color 动物的颜色
	Color string
}

func (a *Animal) GetInfo() string {
	return fmt.Sprintf("识别码: %v, 品种: %v, 年龄: %v, 颜色, %v", a.ID, a.Type, a.Age, a.Color)
}

func (a *Animal) Outter() {}

// PetAnimal 宠物动物
type PetAnimal struct {
	// Animal 包含Animal的基本信息
	Animal
	// OwnerName 主人名字
	OwnerName string
	// Name 宠物的名字
	Name string
}

func (a *PetAnimal) GetInfo() string {
	return fmt.Sprintf("%v, 名字: %v, 主人名字: %v", a.Animal.GetInfo(), a.OwnerName, a.Name)
}

// PetDog 宠物狗
type PetDog struct {
	PetAnimal
}

func (pd *PetDog) Outter() {
	println("旺旺!")
}

// PetCat 宠物猫
type PetCat struct {
	PetAnimal
}

func (pc *PetCat) Outter() {
	println("喵喵!")
}

type AnimalIface interface {
	Outter()
	GetInfo() string
}

func main() {
	var iface AnimalIface
	iface = &PetDog{
		PetAnimal{
			Name: "小白",
			Animal: Animal{
				ID:    "1",
				Type:  "哈士奇",
				Age:   3,
				Color: "黑白相间",
			},
		},
	}
	fmt.Println(iface.GetInfo())
	iface.Outter()
}
