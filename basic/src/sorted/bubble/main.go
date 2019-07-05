package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	arr := []int{1, 6, 2, 4, 9, 0, 5, 3, 7, 8}

	arrRand:=randArry(10)

	//bubbleSort(arr)

	//selectSort(arr)

	//InsertSort(arr)

	fmt.Println(arrRand)

	ShellSort(arr)


	//fmt.Println(arr)
	//fmt.Println(11/2)


}

/*
每次轮训比较相邻的，如果i>i+1 则交换，第一轮最大的在最后一个 问题规模n 轮训n次 复杂度为O^2
冒泡排序的思想接近正常思维，要对西瓜重量从小到大排序，首先每次只比较相邻的两个西瓜，加入某西瓜的重量大于下一个则进行交换
这样所有西瓜都被交换一次后，最大的西瓜就会被放在最后面，这样轮训n-1次后，西瓜被排序好
这样的问题很明显：会存在重复比较或者无效比较的问题，[1,2,3,4],依然要比较
这种问题在选择排序/插入排序中同样存在
*/
func bubbleSort(arr []int) {

	for j := 0; j < len(arr)-2; j++ {
		for i := 0; i < len(arr)-j-1; i++ {
			if arr[i] > arr[i+1] {

				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
			//fmt.Println(arr)



		}


	}

}

/*
选择排序 每次轮训找到最小/最大的，放在与第一个位置进行交换,近似动态规划 O(n^2)
https://blog.csdn.net/hellozhxy/article/details/79911867
选择排序，指挥官下达的命令首先找到最小，最小的与第一个换位置
与冒泡类似复杂度较高
最佳情况：T(n) = O(n2)  最差情况：T(n) = O(n2)  平均情况：T(n) = O(n2)

*/

func selectSort(arr []int) {

	for i := 0; i < len(arr); i++ {
		minIndex := i
		for j := i; j < len(arr); j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}

		}
		arr[minIndex], arr[i] = arr[i], arr[minIndex]
	}

}

/*
对已有序列进行插入，从后往前进行轮训
插入排序的的思想，认为前面的都是排序好的，只需要从大到小进行比较，如果该西瓜的重量大于等于某位置，就插入，否则继续比较
可想而知最好的情况是[1,2,3,4,5]，时间复杂度O(n) 一般情况下还是O(n^2)

最佳情况：T(n) = O(n)   最坏情况：T(n) = O(n2)   平均情况：T(n) = O(n2)

 */
func InsertSort(arr []int) {

	for i := 0; i < len(arr)-1; i++ {
		current := arr[i+1]
		preindex := i

		//for {
		//	// 先决条件 preindex 在前
		//	if ((preindex >= 0) && current <= arr[preindex]) {
		//		arr[preindex+1] = arr[preindex]
		//		preindex--
		//
		//	} else {
		//
		//		break
		//	}
		//
		//}

		for ; ((preindex >= 0) && current <= arr[preindex]); preindex-- {
			arr[preindex+1] = arr[preindex]

		}
		arr[preindex+1] = current

	}

}
/*
插入排序每次只对一个元素进行调节，这样对调整的速度太慢，shellsort 是对它的一种改进，先从宏观上在微观调节，以西瓜重量排序
为例子，指挥官指令为：假设有10个西瓜
- 位置是5的倍的进行排序，调整后，宏观上看小的会跑到前面，大的会跑到后面
- 位置为2的背书的进行排序，慢慢的进行微观上的调整，最终将西瓜重量排序完成
最佳情况：T(n) = O(nlog2 n)  最坏情况：T(n) = O(nlog2 n)  平均情况：T(n) =O(nlog2n)

*/
func ShellSort(arr []int)  {

	length:=len(arr)

	gap:=length/2
	for ; gap>0;gap--{


		for i:=0;i+gap<len(arr);i++{
			if arr[i+gap]<arr[i]{
				arr[i+gap],arr[i]=arr[i],arr[i+gap]
			}
		}

	}


	
}



func quickSort(arr []int) []int {

	if len(arr)<=1{
		return  arr
	}

	median:=arr[rand.Intn(len(arr))]

	lowPart:=make([]int,0,len(arr))
	highPart:=make([]int,0,len(arr))

	middlePart:=make([]int,0,len(arr))

	for _,item:=range arr{
		switch  {
		case item<median:
			lowPart = append(lowPart,item)
		case item==median:
			middlePart=append(middlePart,item)

		case item>median:
			highPart=append(highPart,item)


		}
	}

	lowPart=quickSort(lowPart)
	highPart=quickSort(highPart)

	lowPart=append(lowPart,middlePart...)
	lowPart=append(lowPart,highPart...)
	return lowPart

}

func randArry(n int) []int{


	arr:=make([]int,n)

	// go会默认随机种子，这样每次就会一样，所有需要自己设置随机种子
	rand.Seed(time.Now().UnixNano())

	for i:=0;i<n;i++ {
		arr[i]=rand.Intn(100)
	}

	return  arr
}






