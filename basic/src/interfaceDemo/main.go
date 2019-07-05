package main

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"time"
)

/*
Go 语言提供了另外一种数据类型即接口，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。

type 接口名称  interface {

	funcName()

}


type structName struct {
}

func (st,structName) funcName(){
}
*/


type tranformer interface {
	fit()
	transform()

}

type Lr struct {
	maxIter int
	learingRate float64

}

type  Gbdt struct {

	maxDepth int
	maxTrees int
	learingRate float64
}

func (model Lr) fit() {

	fmt.Println("正在参数设置完毕")


}

func (model Lr) transform() {
	fmt.Println("Lr正在训练")
}

func (model Gbdt) fit() {
	fmt.Println("GBDT参数设置完毕")

}

func (model Gbdt) transform()  {

	fmt.Println("GBDT 正在训练")
}

func showInface(t tranformer) {

	ddt:=reflect.TypeOf(t)


	fmt.Println("我是transformer 类型",ddt)

}


func main() {

	lr:=Lr{20,0.001}
	gbdt:=Gbdt{10,2,0.001}

	lr.transform()
	gbdt.transform()


	showInface(lr)
	showInface(gbdt)

	dd:=Factorial(10)
	fmt.Println(dd)

	q1,e1:=Sqrt(64.)
	q2,e2:=Sqrt(-64.)

	fmt.Println(q1,e1)
	fmt.Println(q2,e2)


	go say("hello")

	say("weiyudang")

	time.Sleep(1000*time.Millisecond)



}

//递归

func Factorial(n uint64) (result uint64)  {

	if n>0 {
		result=n*Factorial(n-1)
		return result
	}

	return  1

}


// 错误处理
func  Sqrt(f float64) (float64 ,error)  {

	if (f<0){
		return 0,errors.New("math:square root of negative number")
	}

	return math.Sqrt(f),errors.New("no error")

}


// 并发

func say(s string)  {

	for i:=0;i<5;i++{
		time.Sleep(100*time.Millisecond)

		fmt.Println(i,s)
	}



}
