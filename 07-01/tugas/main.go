package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type User struct {
	Data     []Data     `json:"data"`
	Included []Included `json:"included"`
}

type Data struct {
	Type         string       `json:"type"`
	Id           string       `json:"id"`
	Attributtes  Attributtes  `json:"attributes"`
	Relationship Relationship `json:"relationship"`
}
type Attributtes struct {
	Title   string    `json:"title"`
	Body    string    `json:"body"`
	Created time.Time `json:"created"`
	Update  time.Time `json:"update"`
}
type Relationship struct {
	Author Author `json:"author"`
}
type Author struct {
	Data DataAuthor `json:"data"`
}
type DataAuthor struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type Included struct {
	Type        string          `json:"type"`
	Id          string          `json:"id"`
	Attributtes AttributtesIncl `json:"attributes"`
}
type AttributtesIncl struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

func main() {
	jsonString := `{
		"data": [
		  {
			"type": "articles",
			"id": "1",
			"attributes": {
			  "title": "JSON:API paints my bikeshed!",
			  "body": "The shortest article. Ever.",
			  "created": "2015-05-22T14:56:29.000Z",
			  "updated": "2015-05-22T14:56:28.000Z"
			},
			"relationships": {
			  "author": {"data": {"id": "42","type": "people"}}
			}
		  }
		],
		"included": [
		  {
			"type": "people",
			"id": "42",
			"attributes": {
				
			  "name": "John",
			  "age": 80,
			  "gender": "male"
			}
		  }
		]
	}`

	var data User

	err := json.Unmarshal([]byte(jsonString), &data)

	if err != nil {
		panic(err)
	}

	fmt.Println(data)

}
