package main

import (
	"fmt"
)

//被代理的公共函数
//
type ProxyFuncs interface {
	//卖房功能
	SailHouse()
}

type MasterBeijing struct {
	Name     string //北京业主姓名
	Location string //业主所卖房屋的位置
}

func (this *MasterBeijing) SailHouse() {
	fmt.Printf("%s sailing house at %s\n", this.Name, this.Location)
}

type Proxier struct {
	Mofbj *MasterBeijing
}

func (this *Proxier) SailHouse() {
	if this.Mofbj == nil {
		this.Mofbj = &MasterBeijing{}
	}

	this.Mofbj.SailHouse()
}

func main() {
	m := &MasterBeijing{
		Name:     "Lao wang",
		Location: "Xi Cheng",
	}
	//可以直接调用自己的方法
	m.SailHouse()

	//实例化到另一个struct，再调用自己的方法
	proxier := &Proxier{
		Mofbj: m,
	}

	proxier.SailHouse()
}
