package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Cow struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	IsAbducted bool   `json:"isAbducted"`
}

type GameState struct {
	FieldCows map[int]*Cow `json:"fieldCows"`
	UfoCows   map[int]*Cow `json:"ufoCows"`
}

var CurrentSave GameState
var FieldCows map[int]*Cow
var UfoCows map[int]*Cow
var CowsNames = []string{"Лёля", "Боля", "Боря", "Моля", "Мурка", "Дурка", "Курка", "Черная", "Белая", "Фиолетовая", "Зеленая", "ERROR"}

func main() {
	initCows(CowsNames)
	cowchecker()
	saveData()
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

func saveData() {
	CurrentSave = GameState{FieldCows: FieldCows, UfoCows: UfoCows}
	jsonData, err := json.MarshalIndent(CurrentSave, "", "-->")
	if err != nil {
		log.Fatalf("Error GameState marshal: %v", err)
	}
	fmt.Println("jsonData compiled successfully")
	err = os.WriteFile("cows.json", jsonData, 0644)
	if err != nil {
		log.Fatalf("Error creating file %v", err)
	}
	fmt.Println("Файл cows.json успешно создан")
}

func loadData() {
	jsonData, err := os.ReadFile("cows.json")
	if os.IsNotExist(err) {
		fmt.Println("сохранения нет, создать? \ny/n:")
		var answer rune
		fmt.Scan(&answer)
		for answer != 'y' || answer != 'Y' || answer != 'n' || answer != 'N' {

		}
	}

	// log.Fatalf("Error load save stage, %v", err)
	json.Unmarshal(jsonData, &CurrentSave)
}

func userResponse() bool {
	var answer string
	var validAnswer bool = false
	var yes = []string{"Y", "y", "yes", "YES", "Yes", "да", "ДА", "Да", "+", "1"}
	var no = []string{"N", "n", "no", "NO", "No", "нет", "НЕТ", "Нет", "-", "0"}

	fmt.Println("y/n:")
	fmt.Scan(&answer)
	count := 0
	for count; !validAnswer || count < 10; count++ {
		if answer == yes {
			validAnswer = true
		} else if answer == no {
			validAnswer = true
			return false
		} else {
			fmt.Println("")
			validAnswer = false
		}
	}
	if count >= 10 {
		fmt.Println("Вероятно пользователь дурачек, засчитываем как 'нет'")
		return false
	}

}
