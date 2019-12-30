package main

import (
	"fmt"
	"time"
)

type Event struct {
	Data string
}

type Observer interface {
	//更新事件
	Update(*Event)
}

// 被观察的对象接口
type Subject interface {
	//注册观察者
	Regist(Observer)
	//注销观察者
	Deregist(Observer)

	//通知观察者事件
	Notify(*Event)
}

type ConcreteObserver struct {
	Id int
}

func (co *ConcreteObserver) Update(e *Event) {
	fmt.Printf("observer [%d] recieved msg: %s.\n", co.Id, e.Data)
}

type ConcreteSubject struct {
	Observers map[Observer]struct{}
}

func (cs *ConcreteSubject) Regist(ob Observer) {
	//struct{}{}  空结构  struct{} 类型 ，｛｝ 实例化
	cs.Observers[ob] = struct{}{}
}

func (cs *ConcreteSubject) Deregist(ob Observer) {
	delete(cs.Observers, ob)
}

func (cs *ConcreteSubject) Notify(e *Event) {
	for ob, _ := range cs.Observers {
		ob.Update(e)
	}
}

func main() {
	//创建观察者 容器
	cs := &ConcreteSubject{
		//这里的空结构无意义，程序中没有调用。
		Observers: make(map[Observer]struct{}),
	}

	//创建观察者
	cobserver1 := &ConcreteObserver{1}
	cobserver2 := &ConcreteObserver{2}
	//注册观察者到容器
	cs.Regist(cobserver1)
	cs.Regist(cobserver2)

	for i := 0; i < 5; i++ {
		e := &Event{fmt.Sprintf("msg [%d]", i)}
		//notify 调用了 e的update方法
		cs.Notify(e)

		time.Sleep(time.Duration(1) * time.Second)
	}
}
