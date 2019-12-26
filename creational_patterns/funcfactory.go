package main

import (
	"fmt"
)

//创建接口  封装进去一个move方法
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

type AnimalFactory interface {
	CreateAnimal() Action
}

type BirdFactory struct {
}

func (this *BirdFactory) CreateAnimal() Action {
	return &Bird{}
}
//定义一个类型  实际是空接口
type FishFactory struct {
}
//上面类型的方法   返回一个实例化的接口
func (this *FishFactory) CreateAnimal() Action {
	return &Fish{}
}

func main() {
	bFactory := &BirdFactory{}
	bird := bFactory.CreateAnimal()
	bird.Move(100)

	fFactory := &FishFactory{}
	fish := fFactory.CreateAnimal()
	fish.Move(200)

}
