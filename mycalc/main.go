package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	value1, value2 := readInput(), readInput()
	opsReader := bufio.NewReader(os.Stdin)
	fmt.Print("Select an operation (+ - * /):")
	ops, _ := opsReader.ReadString('\n')
	ops = strings.TrimSpace(ops)
	result := Operate(value1, value2, ops)
	fmt.Println("Resulted value:", result)
}

func readInput() float64 {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter value:")
	input1, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	if err != nil {
		fmt.Println(err)

	}
	return value
}

func Operate(val1 float64, val2 float64, operation string) float64 {
	switch operation {
	case "+":
		return val1 + val2
	case "-":
		return val1 - val2
	case "*":
		return val1 * val2
	case "/":
		return val1 / val2
	default:
		panic("unsupported operation")
	}
}
