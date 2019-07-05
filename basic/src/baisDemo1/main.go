package main


import (
	"./config"
	"fmt"
)



func main() {

	fmt.Println("hello world")
	var age int
	age=2
	fmt.Println(age)

	var a,b int
	a=20
	b=23
	fmt.Println(Max(a,b))

	fmt.Println(GetSeq()())

	// 常量的设置
	const(
		C1=1<<iota
		C2
		C3

	)

	fmt.Println(C1,C2,C3)


	//binary
	fmt.Printf("%d - %b\n",42,42)

	var sum=0
	for i:=0;i<100;i++{
		sum+=i

	}

	fmt.Printf("100的加和是:%d\n",sum)

	// package  call
	fmt.Println(config.Name)


	config.ShowInfo()



}


//aa
func Max(a,b int) int {
	if(a>b){
		return a
	}else{
		return b
	}
}

//闭包，内部函数可以访问外部函数的变量

func GetSeq() func() int{
	i:=1

	return func() int {
		i+=1
		return i
	}

}



