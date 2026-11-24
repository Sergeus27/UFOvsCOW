package main

import (
	"encoding/json"
	// "errors"
	"fmt"
	"log"
	// "math/rand"
	"net/http"
	"os"
	"sort"
	"strings"
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

func main() {

	initCows()
	isCowInMap()

	printOrderedCows(CurrentSave.FieldCows)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello http!")
}

func initCows() {
	var cowsNames = []string{"Лёля", "Боля", "Боря", "Моля", "Мурка", "Дурка", "Курка", "Черная", "Белая", "Фиолетовая", "Зеленая", "ERROR"}

	CurrentSave.FieldCows = make(map[int]*Cow)
	CurrentSave.UfoCows = make(map[int]*Cow)

	for i, name := range cowsNames {

		id := i + 1
		// rnd := rand.Intn(12)
		// fmt.Println(rnd)
		// CurrentSave.FieldCows[rnd+1] = &Cow{rnd + 1, name, false}
		CurrentSave.FieldCows[id] = &Cow{id, name, false}
	}

}

func isCowInMap() bool {

	if (len(CurrentSave.FieldCows) + len(CurrentSave.UfoCows)) == 0 {
		fmt.Println("коров нет нигде")
		return false
	}

	if len(CurrentSave.FieldCows) > 0 {
		fmt.Println("Коровы на лугу:")

		for i, a := range CurrentSave.FieldCows {
			fmt.Printf("%d.) Корова:'%s', Забрали: ", i, a.Name)
			if a.IsAbducted {
				fmt.Println("Да")
			} else {
				fmt.Println("Нет")
			}
		}

	} else {
		fmt.Println("нет коров на лугу")

	}

	if len(CurrentSave.UfoCows) > 0 {
		fmt.Println("Коровы у пришельцев:")

		for i, a := range CurrentSave.UfoCows {
			fmt.Printf("%d.) Корова:'%s', Забрали: ", i, a.Name)
			if a.IsAbducted {
				fmt.Println("Да")
			} else {
				fmt.Println("Нет")
			}
		}

	} else {
		fmt.Println("нет коров у пришельцев")

	}

	return true

}

func saveData() {

	jsonData, err := json.MarshalIndent(CurrentSave, "", "\t")
	if err != nil {
		log.Fatalf("Error GameState marshal: %v", err)
	}
	fmt.Println("jsonData compiled successfully")
	err = os.WriteFile("cows.json", jsonData, 0644)
	if err != nil {
		log.Fatalf("Error creating file %v", err)
	}
	fmt.Println("Файл cows.json успешно создан")
	// errors.New("сука2")
}

func loadData() {
	jsonData, err := os.ReadFile("cows.json")

	if os.IsNotExist(err) {
		fmt.Println("сохранения нет, создать? \ny/n:")
		if getUserResponse() {
			initCows()
		} else {
			fmt.Println("ну и пошел на хуй тогда")
			return
		}
	}
	if err != nil {
		log.Fatalf("файл существует. Ошибка в чтении: %v", err)
	}
	// log.Fatalf("Error load save stage, %v", err)
	err = json.Unmarshal(jsonData, &CurrentSave)
	if err != nil {
		log.Fatalf("Ошибка распаковки json: %v", err)
	} else {
		fmt.Println("файл успешно загружен")
	}
}

func getUserResponse() bool {
	// errors.New("сука")
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

func mysort(cowsMap map[int]*Cow, sortFunc func(i, j int) bool) {

}

func printOrderedCows(cowsMap map[int]*Cow) {
	var cowsSlice []*Cow

	for _, cow := range cowsMap {
		cowsSlice = append(cowsSlice, cow)
	}

	sort.Slice(cowsSlice, func(i, j int) bool {
		return cowsSlice[i].Id < cowsSlice[j].Id
	})

	fmt.Println("дальше вывод отсортированных по проядку коров:")
	for i, a := range cowsSlice {
		fmt.Print(i, ".) ", a.Id, " --- ", a.Name)
		if a.IsAbducted {
			fmt.Println(", спиздили")
		} else {
			fmt.Println(", не спиздили")
		}
	}
}
