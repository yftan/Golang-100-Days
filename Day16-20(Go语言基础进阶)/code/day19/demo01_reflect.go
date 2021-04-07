package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	Name string
	Age  int
}

func (animal *Animal) Say() {
	fmt.Println("I am a animal")
}

func (animal *Animal) Eat(food string) {
	fmt.Printf("I am eating %s\n", food)
}

func main() {
	//反射操作：通过反射，可以获取一个接口类型变量的 类型和数值
	var x float64 = 3.4

	fmt.Println("type:", reflect.TypeOf(x))   //type: float64
	fmt.Println("value:", reflect.ValueOf(x)) //value: 3.4

	fmt.Println("-------------------")
	//根据反射的值，来获取对应的类型和数值
	v := reflect.ValueOf(x)
	fmt.Println("kind is float64: ", v.Kind() == reflect.Float64)
	fmt.Println("type : ", v.Type())
	fmt.Println("value : ", v.Float())

	fmt.Println("-------------------")
	var a int32 = 22
	t1 := reflect.TypeOf(a)
	v1 := reflect.ValueOf(a)
	fmt.Println("type:", t1)
	fmt.Println("kind:", t1.Kind())
	fmt.Println("value:", v1)

	fmt.Println("-------------------")
	animal := &Animal{
		Name: "贝克",
		Age:  8,
	}
	t2 := reflect.TypeOf(animal)
	v2 := reflect.ValueOf(animal)
	fmt.Println("type:", t2)
	fmt.Println("kind:", t2.Kind())
	fmt.Println("value:", v2)
	vSay := v2.MethodByName("Say")
	vSay.Call(nil)
	vEat := v2.MethodByName("Eat")
	fmt.Printf("Kind : %s, Type : %s\n", vEat.Kind(), vEat.Type())
	values := []reflect.Value{reflect.ValueOf("apple")}
	vEat.Call(values)

	fmt.Println("-------------------")
	fmt.Printf("%p", funtest)
	f := reflect.ValueOf(funtest)
	f.Call(nil)
}

func funtest() {
	fmt.Println("fun test")
}
