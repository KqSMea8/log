package main

import (
	"fmt"
	"reflect"
)

type class struct {
	peoples []people
}

type people struct {
	name string
	age int
	other map[string]string
}

func main() {

	sliceMap1 := make([]map[interface{}]interface{},0)

	sliceMap2 := make([]map[interface{}]interface{},0)


	map1 := make(map[interface{}]interface{},0)
	map2 := make(map[interface{}]interface{},0)
	map3 := make(map[interface{}]interface{},0)
	map4 := make(map[interface{}]interface{},0)


	map1["step"] = 120
	map1["values"] = []byte{1,2,3,4}

	map2["step"] = 60
	map2["values"] = []byte{2,2,3,4}


	map3["step"] = 120
	map3["values"] = []byte{1,2,3,4}

	map4["step"] = 60
	map4["values"] = []byte{2,2,3,4}


	sliceMap1 = append(sliceMap1,map1)
	sliceMap1 = append(sliceMap1,map2)

	sliceMap2 = append(sliceMap1,map3)
	sliceMap2 = append(sliceMap1,map4)


	fmt.Println(reflect.DeepEqual(sliceMap1,sliceMap2))


}
