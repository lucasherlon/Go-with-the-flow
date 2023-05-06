package main

import (
    "fmt"
    "math"
)

// hackerranck 05
func diagonalDifference(arr [][]int32) int32 {
    var firstDiag int32
    var secondDiag int32
    for i := 0; i < len(arr); i++ {
        for j := 0; j < len(arr); j++ {
            if i == j {
                firstDiag += arr[i][j]
            }
            if (i + j) == (len(arr) - 1) {
                secondDiag += arr[i][j]
            }
        }
    }
    res := int32(math.Abs(float64(firstDiag) - float64(secondDiag)))
    return res
}


func main() {
    
    matrix := [][]int32{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
    result := diagonalDifference(matrix)
    fmt.Println(result)  
}
