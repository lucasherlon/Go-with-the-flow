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
	var input1 string
	var input2 string
	var firstPoint []float64
	var secondPoint []float64

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input1 = scanner.Text()
	scanner.Scan()
	input2 = scanner.Text()

	vector1 := strings.Split(input1, " ")
	for _, num := range vector1 {
		f, err := strconv.ParseFloat(num, 64)
		if err != nil {
			fmt.Println("Error at parsing!")
		}
		firstPoint = append(firstPoint, f)
	}

	vector2 := strings.Split(input2, " ")
	for _, num := range vector2 {
		f, err := strconv.ParseFloat(num, 64)
		if err != nil {
			fmt.Println("Error at parsing!")
		}
		secondPoint = append(secondPoint, f)
	}

	x1 := firstPoint[0]
	y1 := firstPoint[1]
	x2 := secondPoint[0]
	y2 := secondPoint[1]

	distance := math.Sqrt(math.Pow((x2 - x1), 2) + math.Pow((y2 - y1), 2))
	fmt.Printf("%.4f\n", distance)

}
