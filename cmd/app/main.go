package main

import (
	"fmt"
	"log"

	"github.com/ladislavkosenko/lab1-tooling/internal/calc"
)

func main() {
	fmt.Println("Lab 1: Tooling setup is working!")

	result, err := calc.Divide(10.0, 2.0)
	if err != nil {
		log.Fatalf("Помилка ділення: %v", err)
	}

	fmt.Printf("Результат ділення: %f\n", result)
}
