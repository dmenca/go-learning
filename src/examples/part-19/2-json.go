package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	person := Person{
		Name:  "John",
		Age:   30,
		Email: "john@qq.com",
	}
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("JSON marshal failed:", err)
	}
	fmt.Printf("Json data:%s\n", string(jsonData))

	var person1 Person
	err = json.Unmarshal(jsonData, &person1)
	if err != nil {
		fmt.Println("JSON unmarshal failed:", err)
	}
	fmt.Printf("Person info: {Name: %s, Age: %d, Email: %s}\n",
		person1.Name, person1.Age, person1.Email)

}
