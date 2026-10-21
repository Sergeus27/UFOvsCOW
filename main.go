package main

import (
	"fmt"
	"net/http"
)

type Cow struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	IsAbducted bool   `json:"isAbducted"`
}

var FieldCows map[int]*Cow
var UfoCows map[int]*Cow
var CowsNames = []string{"Лёля", "Боля", "Боря", "Моля", "Мурка", "Дурка", "Курка", "Черная", "Белая", "Фиолетовая", "Зеленая", "ERROR"}

func main() {
	initCows(CowsNames)
	cowchecker()
	//http.HandleFunc("/hello", helloHandler)
	//http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello http!")
}

func initCows(names []string) {
	FieldCows = make(map[int]*Cow)
	UfoCows = make(map[int]*Cow)

	for i, name := range names {
		FieldCows[i+1] = &Cow{i + 1, name, false}
	}

}

func cowchecker() {
	if len(FieldCows) > 0 {
		fmt.Println("Коровы на лугу:")
		for i := 1; i <= len(FieldCows); i++ {
			fmt.Println(i, ":   ", FieldCows[i])
		}
	} else {
		fmt.Println("нет коров на лугу-")
	}

	if len(UfoCows) > 0 {
		fmt.Println("Коровы у пришельцев:")
		for i := 1; i <= len(UfoCows); i++ {
			fmt.Println(i, ":   ", UfoCows[i])
		}
	} else {
		fmt.Println("У пришельцев нет коров")
	}
}
