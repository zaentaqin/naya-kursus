package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id    int
	Name  string
	Email string
	Phone string
}

func main() {

	jsonString := `[
	{
	  "id": 5,
	  "name": "Zaenal",
	  "email": "zaenal@example.com",
	  "phone": "08734732323"
	},
	{
	  "id": 6,
	  "name": "alex",
	  "email": "alex@example.com",
	  "phone": "087347323546"
	},
	{
	  "id": 7,
	  "name": "zendaya",
	  "email": "zendaya@example.com",
	  "phone": "08734732654"
	}
  ]`

	var data User

	err := json.Unmarshal([]byte(jsonString), &data)

	if err != nil {
		panic(err)
	}

	fmt.Println(data)

}
