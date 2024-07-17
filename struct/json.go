package main

import (
	"encoding/json"
	"fmt"
)

type Result struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type User struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func ResultSuccess(data interface{}) *Result {
	return &Result{
		Data: data,
	}
}

func main() {
	user := []User{{
		Name:    "H",
		Age:     20,
		Address: "earth",
	}}
	marshal, err := json.Marshal(ResultSuccess(user))
	if err != nil {
		return
	}
	res := string(marshal)
	fmt.Println(res)

	var user1 User
	jsonstr := `{"name":"黄","age":20,"address":"earth"}`
	err = json.Unmarshal([]byte(jsonstr), &user1)
	if err != nil {
		return
	}
	fmt.Println(user1.Name)
}
