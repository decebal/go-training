package main
import (
    "fmt"
    "bufio"
    "os"
    "strconv"
)

func add (x int, y string) int {
    second, err := strconv.Atoi(y)
    if err != nil {
        fmt.Print("Error: %s \n", err)
    }
    return x + second
}

func abs (x int) int {
    if (x < 0) {
        return -x
    }

    return x
}

func main() {
    var N int
    var lines []string
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    N,err := strconv.Atoi(lines[0])
    if err != nil {
        fmt.Print("Error: %s \n", err)
    }

    if N <= 0 {
        fmt.Print(0)
        os.Exit(200)
    }

    elements := lines[1:]

    var sumA, sumB int
    for i := 0; i < N; i++ {
        sumA = add(sumA, elements[N*i+i])
        sumB = add(sumB, elements[N*(i+1)-i-1])

    }
    fmt.Print(abs(sumA-sumB))
}

