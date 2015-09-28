package main

import "fmt"
import "math"

func main() {
	var n int

	fmt.Scan(&n)

	if n < 1 || n > 100 {
		return
	}

	matrix := make([][]int, n)

	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)

		for j := 0; j < n; j++ {
			fmt.Scan(&matrix[i][j])

			if matrix[i][j] < -100 || matrix[i][j] > 100 {
				return
			}
		}
	}

	i := 0
	j := 0

	diagA := 0
	for i < n {
		diagA += matrix[i][j]
		i++
		j++
	}

	i = n - 1
	j = 0

	diagB := 0
	for j < n {
		diagB += matrix[i][j]
		i--
		j++
	}

	diff := float64(diagA - diagB)

	fmt.Println(math.Abs(diff))
}
