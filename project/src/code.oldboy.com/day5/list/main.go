package main

import(
	"code.oldboy.com/day5/model"
	"fmt"
)

func main() {
	school := model.NewSchool
	school("peking_university","Haidianqu")
	fmt.Println("The university name is: %v",&school)
}