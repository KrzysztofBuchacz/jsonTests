package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
)

type _Data struct {
	Volume float64
}

func main() {
	var data _Data
	data.Volume = math.Inf(1)
	substr, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
	err = ioutil.WriteFile("myjsonfile.json", []byte(substr), 0644)
	if err != nil {
		fmt.Println(err)
	}

}
