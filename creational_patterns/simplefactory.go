package main

import (
	"fmt"
)

type Action interface {
	Move(int) int
}

type Animal struct {
	Name string
}

type Bird struct {
	Animal
}

func (this *Bird) Move(num int) int {
	fmt.Printf("I am a bird, I flyed %d meters.\n", num)
	return num
}

type Fish struct {
	Animal
}

func (this *Fish) Move(num int) int {
	fmt.Printf("I am a fish, I swimmed %d meters.\n", num)
	return num
}

type Dog struct {
	Animal
}

func (this *Dog) Move(num int) int {
	fmt.Printf("I am a dog, I flying %d meters.\n", num)
	return num
}
//定义一个类型  实际是空接口
type AnimalFactory struct {
}

//创建上面类型的 函数。返回一个上面的类型
func NewAnimalFactory() *AnimalFactory {
	return &AnimalFactory{}
}

//上面类型的方法   返回一个实例化的接口
func (this *AnimalFactory) CreateAnimal(name string) Action {
	switch name {
	case "bird":
		return &Bird{}
	case "fish":
		return &Fish{}
	case "dog":
		return &Dog{}
	default:
		panic("error animal type")
		return nil
	}
}

func main() {
	//创建一个类型调，并用类型的方法。
	bird := NewAnimalFactory().CreateAnimal("bird")
	bird.Move(100)

	fish := NewAnimalFactory().CreateAnimal("fish")
	fish.Move(200)

	dog := NewAnimalFactory().CreateAnimal("dog")
	dog.Move(300)
}
