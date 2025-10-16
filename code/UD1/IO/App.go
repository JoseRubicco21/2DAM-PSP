package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Product struct {
  ID string
	Name string
	Category string
	Stock int
  Price float64
}

type Transaction struct {
	Type          string
	ProductID     string
	Quantity      int32
	Date          string
}


func readInventory(fileName string) (map[string]Product, error) {

	var products = make(map[string]Product)

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
		pr, err := parseProduct(text)
	  if err != nil {
			return nil, fmt.Errorf("Error creating map. Cannot insert nil value")
		} else {
		products[text] = pr
		}
		
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Scanner error while reading lines: %w", err)
	}

	return products, nil
}

func parseProduct(line string) (Product, error) {
	var product Product
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


// Maybe we can use the csv lib here ? 
func parseTransaction(line string) (Transaction, error){
	var transaction Transaction
	values := strings.Split(strings.TrimSpace(line), ",")

	transaction.Type = values[0]
	transaction.ProductID = values[1]
	

	quantity, err := strconv.ParseInt(values[2],10, 32)
	if err != nil {
		return transaction, fmt.Errorf("Invalid quantity in transaction registry")
	}
	transaction.Quantity = int32(quantity)
	
	transaction.Date = values[3]

	return transaction, nil

}

//TODO: Make it return a map[string]Transaction where str is the transaction ID.
func readTransactions(filename string) ([]Transaction, error) {
 
	var transactions =  make([]Transaction,0)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: opening transactions file: %s\n", err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		tr, err := parseTransaction(text)
		if err != nil {
			errtxt := fmt.Errorf("Error parsing transaction")
		fmt.Println(errtxt)
		} else {
		transactions = append(transactions, tr)		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error while reading transactions: %s\n", err)
	}

	return transactions, nil
}



func displayProductInformation(str string) {
	fmt.Println(str)
}

func displayProductInMemory(products []Product) {
	for _, p := range products {
		displayProductInformation("-----")
		displayProductInformation("ID: " + p.ID)
		displayProductInformation("Name: " + p.Name)
		displayProductInformation("Category: " + p.Category)
		displayProductInformation(fmt.Sprintf("Price: %.2f", p.Price))
		displayProductInformation(fmt.Sprintf("Stock: %d", p.Stock))
		
}




func main() {
	_, err := readInventory("./data/inventario.txt")
	if err != nil {
		fmt.Printf("Failed to read inventory: %v\n", err)
		return
	}

	// Create prod rep and print them wihth built-in
	
	//fmt.Printf("prods: %v\n", prods)


	transactions, err := readTransactions("./data/transacciones.txt")
 
	if err != nil {
		fmt.Print("Error reading the transaction file")
		return
	}

	
	fmt.Printf("transactions: %v\n", transactions)
	//displayProductInMemory(prods)

	// To test reading transactions:
	// readTransactions("./data/transacciones.txt")
}


