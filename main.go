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

func main() {
	initCows()
	cowchecker()
	saveData()
	//http.HandleFunc("/hello", helloHandler)
	//http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello http!")
}

func initCows() {
	var cowsNames = []string{"Лёля", "Боля", "Боря", "Моля", "Мурка", "Дурка", "Курка", "Черная", "Белая", "Фиолетовая", "Зеленая", "ERROR"}

	FieldCows = make(map[int]*Cow)
	UfoCows = make(map[int]*Cow)

	for i, name := range cowsNames {
		FieldCows[i+1] = &Cow{i + 1, name, false}
	}

}

func cowchecker() bool {

	if (len(FieldCows) + len(UfoCows)) == 0 {
		fmt.Println("коров нет нигде")
		return false
	}

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
		fmt.Println("У пришельцев нет коров, и это критическая ошибка")
		log.Fatal("logic error")
	}

	return true

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
		if userResponse() {
			if cowchecker() == false {

			}
		}
	}

	// log.Fatalf("Error load save stage, %v", err)
	json.Unmarshal(jsonData, &CurrentSave)
}

func userResponse() bool {
	var answer string
	var yesList = []string{"y", "yes", "да", "+", "1"}
	var noList = []string{"n", "no", "нет", "-", "0"}
	const maxAttempts = 10
	answerMap := make(map[string]bool)
	for _, i := range yesList {
		answerMap[i] = true
	}

	for _, i := range noList {
		answerMap[i] = false
	}

	var count int = 0

	for count < maxAttempts {
		count++
		fmt.Print("y/n:")
		fmt.Scan(&answer)
		answer = strings.ToLower(answer)

		if answerValue, ok := answerMap[answer]; ok {
			fmt.Printf("затрачено попыток:%d\nответ: %s\nзначение: %v\n", count, answer, answerValue)
			return answerValue
		} else {
			fmt.Println("неверный формат ответа, введите либо ДА либо НЕТ:")
		}

	}
	if count >= 10 {
		fmt.Println("Вероятно пользователь дурачек, засчитываем как 'нет'")
		return false
	} else {
		fmt.Println("ты не должен видеть эту строку")
		return false
	}

}
