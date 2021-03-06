package jsontest

import (
	"encoding/json"
	"fmt"
	"log"
)

// Movie
// json:"color,omitempty" 指定Json 格式化的格式 (field tag)
type Movie struct {
	Title  string
	Year   int
	Color  bool `json:"color,omitempty"`
	Actors []string
}

// MovieToJsonTest print json
func MovieToJsonTest() {

	var movies = []Movie{
		{Title: "Shi Er Sheng Xiao", Year: 2020, Color: true, Actors: []string{"Cheng Long", "Sun Quan"}},
		{Title: "Si Da Ming Bu", Year: 2008, Color: false, Actors: []string{"Rookie", "Jack Love"}},
	}

	j1, err := json.Marshal(movies)

	if err != nil {
		log.Fatal("Error", err)
	}
	fmt.Println()
	fmt.Printf("未缩进显示:  \n %s ", j1)

	fmt.Println()
	fmt.Println()

	// 缩进打印 json
	j2, err := json.MarshalIndent(movies, "", " ")

	if nil != err {
		log.Fatalf("Error :%s", err)

	}
	fmt.Printf("缩进显示: \n %s ", j2)
}
