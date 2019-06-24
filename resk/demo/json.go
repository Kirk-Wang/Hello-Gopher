package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id      int    `json:"id,string"`
	Name    string `json:"username,omitempty"`
	Age     int    `json:"age,omitempty"`
	Address string `json:"-"`
}

func main() {
	u := User{
		Id:      12,
		Name:    "wendell",
		Age:     1,
		Address: "成都高新区",
	}
	data, err := json.Marshal(&u)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}
