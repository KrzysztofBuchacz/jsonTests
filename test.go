package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	bytes, err := ioutil.ReadFile("badInf.json")
	if err != nil {
		fmt.Println(err)
	}
	var data interface{}
	err = json.Unmarshal([]byte(bytes), &data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Data:")
	fmt.Println(data)
}
