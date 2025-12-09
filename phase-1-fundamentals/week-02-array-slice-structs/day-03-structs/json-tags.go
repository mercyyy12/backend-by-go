package main

type Employee struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Role string `json:"-"`
	Emp_address
}

type Emp_address struct {
	Permanent_address string `json:"address"`
	Phone             int    `json:"phone"`
}

func Json_tags() {
	
}
