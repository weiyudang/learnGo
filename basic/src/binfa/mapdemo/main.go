package main

import "fmt"



func main() {
	var modelMap map[string]string
	modelMap=make(map[string]string)
	modelMap["Lr"]="logist regression"
	modelMap["gbdt"]="gradient boost decision trees"
	modelMap["dnn"]="deep neural network"

	for model :=range modelMap {
		fmt.Println(model,"全称：",modelMap[model])

	}




}
