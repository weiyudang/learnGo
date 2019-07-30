package main

import (
	"encoding/json"
	"fmt"
)

type  Bird struct {
	Speices string
	Description string
}

func main() {
	//Structured data
	birdJson := `{"species": "pigeon","description": "likes to perch on rocks"}`
	var bird Bird
	err:=json.Unmarshal([]byte(birdJson),&bird)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(bird.Description)
}
