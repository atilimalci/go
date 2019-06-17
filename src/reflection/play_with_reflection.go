package main

import (
	"fmt"
	"reflect"
)

func playWithReflection() {
	fmt.Println("//Playing with Reflection")
	var x = 3.4
	fmt.Println("valueOf:", reflect.ValueOf(x).String())

	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())
	fmt.Println("===")

	var y uint8 = 'y'
	fmt.Println("y:", y)
	z := reflect.ValueOf(y)
	fmt.Println("type:", z.Type())
	fmt.Println("kind is uint8:", z.Kind() == reflect.Uint8)
	y = uint8(z.Uint())
	fmt.Println("y:", y)
	fmt.Println("===")

	fmt.Println("Second Law: Reflection object to Interface")
	type MyInt int
	var t MyInt = 7
	u := reflect.ValueOf(t)
	fmt.Println("type:", u.Type())
	fmt.Println("kind:", u.Kind())
	s := reflect.ValueOf(u)
	fmt.Println("Stored Value:", s.Interface())
	fmt.Println("===")

	fmt.Println("Third Law: To Modify a reflection object, the value must be settable")
	var l = 3.4
	k := reflect.ValueOf(l)
	fmt.Println("Settability of k:", k.CanSet())

	k = reflect.ValueOf(&l)
	fmt.Println("Settability of k:", k.Elem().CanSet())
	k.Elem().SetFloat(7.1)
	fmt.Println("l after reflection set:", l)
}

func playWithReflectionOfStructs() {
	type T struct {
		A int
		B string
	}

	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n",
			i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
	s.Field(0).SetInt(77)
	s.Field(1).SetString("Sunset Strip")
	fmt.Println("t is now ", t)
}

func main() {
	//playWithReflection()
	playWithReflectionOfStructs()
}
