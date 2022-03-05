package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Value 1:")
	input1, _ := reader.ReadString('\n')

	value1, err1 := strconv.ParseFloat(strings.TrimSpace(input1), 64)

	if err1 != nil {
		fmt.Println(err1)
		panic("Value 1 must be a number")
	}

	fmt.Print("Value 2:")
	input2, _ := reader.ReadString('\n')

	value2, err2 := strconv.ParseFloat(strings.TrimSpace(input2), 64)

	if err2 != nil {
		fmt.Println(err2)
		panic("Value 2 must be a number")
	}

	sum := value1 + value2
	sumRounded := math.Round(sum*100) / 100
	fmt.Printf("The sum of %v and %v is %v\n", value1, value2, sumRounded)

}
