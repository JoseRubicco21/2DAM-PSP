package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Transaction struct {
	TransactionID string
	Type          string
	ProductID     string
	Quantity      int
	Date          string
}


func readInventory(fileName string) (map[string]Product, error) {

	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("Error opening inventory file: %w", err)
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	// Skip first line
	scanner.Scan()
	
	for scanner.Scan() {
		text := scanner.Text()

	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Scanner error while reading lines: %w", err)
	}

	return products, nil
}

func parseProduct(line string) (Producto, error) {
	var product Producto
	values := strings.Split(strings.TrimSpace(line), ",")

	if len(values) < 5 {
		return product, errors.New("not enough values in product line")
	}

	product.ID = values[0]
	product.Name = values[1]
	product.Category = values[2]

	price, err := strconv.ParseFloat(values[3], 64)
	if err != nil {
		return product, fmt.Errorf("invalid price: %v", err)
	}
	product.Price = price

	stock, err := strconv.Atoi(values[4])
	if err != nil {
		return product, fmt.Errorf("invalid stock: %v", err)
	}
	product.Stock = stock

	return product, nil
}

func readTransactions(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening transactions file: %s\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error while reading transactions: %s\n", err)
	}
}

func displayProductInformation(str string) {
	fmt.Println(str)
}

func displayProductInMemory(products []Producto) {
	for _, p := range products {
		displayProductInformation("-----")
		displayProductInformation("ID: " + p.ID)
		displayProductInformation("Name: " + p.Name)
		displayProductInformation("Category: " + p.Category)
		displayProductInformation(fmt.Sprintf("Price: %.2f", p.Price))
		displayProductInformation(fmt.Sprintf("Stock: %d", p.Stock))
	}
}

func procesarTransaccion(productos []Producto, transacciones []Transaction) {
	// Empty for now – you can add logic to apply transactions to inventory
}

func main() {
	prods, err := readInventory("./data/inventario.txt")
	if err != nil {
		fmt.Printf("Failed to read inventory: %v\n", err)
		return
	}

	displayProductInMemory(prods)

	// To test reading transactions:
	// readTransactions("./data/transacciones.txt")
}
