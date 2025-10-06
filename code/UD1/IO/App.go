package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInventory(fileName string) {

	file, err := os.Open(fileName)

	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		displayProductInformation(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error while reading lines. %s", err)
	}

}

func readTransactions(filename string) {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Printf("Error %s", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error while reading lines. %s", err)
	}
}

func displayProductInformation(str string) {
	fmt.Printf("%s\n", str)
}

func main() {
	readInventory("./data/inventario.txt")
	readTransactions("./data/transacciones.txt")
}
