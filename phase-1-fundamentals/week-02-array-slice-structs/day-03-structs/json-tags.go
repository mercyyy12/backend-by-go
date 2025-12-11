package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Role string `json:"-"`
	Emp_address
}

type Emp_address struct {
	Permanent_address string `json:"address"`
	Phone             int    `json:"phone,omitempty"`
}

func Json_tags() {
	e := Employee{"Mercyyy", 19, "sales", Emp_address{"kathmandu", 9841}}

	// marshalling
	data, _ := json.MarshalIndent(e, "", "")
	fmt.Println(string(data))

	// unmarshalling
	var emp Employee
	json.Unmarshal([]byte(data), &emp)
	fmt.Println("Unmarshalled Struct:", emp)
}
