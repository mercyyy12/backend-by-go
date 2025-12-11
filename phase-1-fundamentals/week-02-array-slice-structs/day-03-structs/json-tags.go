package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	Name string `json:"name"` //basic json tag
	Age  int    `json:"age"`
	Role string `json:"-"` // ignoring a field

	// embeded structure
	Emp_address
}

type Emp_address struct {
	Permanent_address string `json:"address"`
	Phone             int    `json:"phone,omitempty"` // omit a field if empty
}

func Json_tags() {
	e := Employee{"Mercyyy", 19, "sales", Emp_address{"kathmandu", 9841}}

	// marshalling (go struct to json)
	data, _ := json.MarshalIndent(e, "", "")
	fmt.Println(string(data))

	// unmarshalling (json to go struct)
	var emp Employee
	json.Unmarshal([]byte(data), &emp)
	fmt.Println("Unmarshalled Struct:", emp)
}
