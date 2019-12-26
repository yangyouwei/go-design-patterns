package main

import "fmt"

//定义一个类型。  为decorate的一个参数类型。
type Object func(string) string

func Decorate(fn Object) Object {
    //返回一个函数。可以继续被装饰。返回的函数执行结果就是decorate的执行结果。
    return func(base string) string {

        ret := fn(base)

        ret = ret + " and Tshirt"
        return ret
    }
}

func Dressing(cloth string) string {
    return "dressing " + cloth
}

func main() {
    f := Decorate(Dressing)

    fmt.Println(f("shoes"))
}
