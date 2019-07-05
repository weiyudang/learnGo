package main

import "fmt"

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


	const(
		C1=1<<iota
		C2
		C3

	)




	fmt.Println(C1,C2,C3)
}


//aa
func Max(a,b int) int {
	if(a>b){
		return a
	}else{
		return b
	}
}

func GetSeq() func() int{
	i:=1

	return func() int {
		i+=1
		return i
	}

}



