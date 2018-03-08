package main

import (
	"fmt"
	"reflect"
)

type Sayer interface {
	Say() string
}

type Cat struct{}

func (c Cat) Say() string { return "meow" }

type Dog struct{}

func (d Dog) Say() string { return "woof" }

func main() {
	c := Cat{}
	d := Dog{}
	animals := []Sayer{c, d}
	for _, a := range animals {
		fmt.Println(reflect.TypeOf(a).Name(), "says:", a.Say())
	}
}
