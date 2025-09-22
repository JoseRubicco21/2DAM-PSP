package main

import (
	"fmt"
)


func isSameLength(str1 string, str2 string) bool {
	if len(str1) == len(str2) return true 
	else return false
}

func sortString(str string) string {
	slice := []rune(str)
	return string(Sort(slice))
}

func main() {

	var palabras = []string{"hola", "aloh", "peso", "gato", "toga"}


}