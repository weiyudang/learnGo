package main

import (
	_ "./util"
	"fmt"
)

type Student  int



/*
1. 数组的定义 var  varName [Size] varType  数组相当与数据结构中的List
	忽略长度 [...]
* 两种遍历方式
- for i:=0;i<len(varName);i++
- for i,v:=range varName


切片（slice）是建立在数组之上的更方便，更灵活，更强大的数据结构。切片并不存储任何元素而只是对现有数组的引用。
元素类型为 T 的切片表示为： []T


*/
func main() {


	var  names [3] string

	arr2 := [3] int{12}

	name:="weiyudang"
	fmt.Println(names[0],arr2,name)

	var project =[...]string{"wee"}
	//project[2]="dd" 初始化时决定序列长度 能改变
	fmt.Println(project,len(project))
	a := [...]float64{67.7, 89.8, 21, 78}


	for i:=0;i<len(a);i++ {
		fmt.Printf("第%d的数字时%f\n",i,a[i])

	}

	for i,v:=range a{
		fmt.Println(i,v)
	}


	var b []float64=a[1:4]
	fmt.Println(b)

	c:=[]int{1,2,3}
	fmt.Println(c)


	fmt.Printf("变量的地址%x\n",&a[0])

	var ip *int
	ip=&c[0]
	fmt.Println(*ip)

	var dd Student
	dd=60
	fmt.Println(dd)









}
