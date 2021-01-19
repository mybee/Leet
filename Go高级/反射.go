package main

import (
	"fmt"
	"strconv"
	"time"
)

//typeof 拿到 type
//
//valueof 拿到 value
//
//type 是 类型 如: person 类型
//
//kind 是 系统类型 如: strct

func main() {
	a := &struct {
		Name string
		Age int
	}{
		Name: "mafeng",
		Age: 1,
	}

	for i:=0; i< 100; i++ {
		go func() {
			a.Name="fdasdad"+strconv.Itoa(i)
			a.Age++
			fmt.Println(a)
		}()
	}

	time.Sleep(10 * time.Second)

}