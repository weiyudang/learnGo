package main

import (
	"encoding/json"
	"fmt"
)

/*
结构体用户的自定义的类型，它代表若干哥字段的集合
*/
type bigInt int

type address struct {
	province string
	city     string
}

type family struct {
	father string
	mother string
}

type person struct {
	name string

	age int


	homeAddress  address
	familymember family
}

// 结构体可以看作python中的类，基于类对象的函数就是成员方法
func (p person) showInfo() {
	fmt.Printf("我的名字叫%q ，我今年%d\n", p.name, p.age)

}

func main() {

	add1:=address{"山东省","菏泽市"}

	fm1:=family{"f1","m2"}

	p1 := person{"weiyudang", 12,
		add1,fm1}

	fmt.Println(p1)

	p1.showInfo()

	p2 := &person{"wangjiaqi", 26,add1,family{"whb","wxx"}}

	fmt.Println((*p2).age)

	bs,_:=json.Marshal(p1)
	fmt.Println(bs)
	fmt.Println(string(bs))

}
