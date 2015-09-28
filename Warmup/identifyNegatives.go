package main
import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
    "math"
)

func countNegatives (x int, y int) int {
    if y < 0 {
        return x + 1
    }

    return x
}

func countPositive (x int, y int) int {
    if y > 0 {
        return x + 1
    }

    return x
}

func countZeros (x int, y int) int {
    if y == 0 {
        return x + 1
    }

    return x
}

func Round(val float64, roundOn float64, places int ) (newVal float64) {
    var round float64
    pow := math.Pow(10, float64(places))
    digit := pow * val
    _, div := math.Modf(digit)
    if div >= roundOn {
        round = math.Ceil(digit)
    } else {
        round = math.Floor(digit)
    }
    newVal = round / pow
    return
}

func main() {
    var N, negatives, positives, zeroes int
    var lines []string

    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    N,err := strconv.Atoi(lines[0])
    if err != nil {
        fmt.Printf("Error: %s\n", err)
    }

    if N <= 0 {
        fmt.Print(0)
        os.Exit(200)
    }

    array := strings.Fields(lines[1])

    for _,element := range array {

        member, err := strconv.Atoi(element)
        if err != nil {
            fmt.Printf("Error: %s\n", err)
        }
        negatives = countNegatives(negatives, member)
        positives = countPositive(positives, member)
        zeroes = countZeros(zeroes, member)
    }


    PosPercent := (float64(positives) / float64(N))
    NegPercent := (float64(negatives) / float64(N))
    ZeroPercent := (float64(zeroes) / float64(N))
    fmt.Print(Round(PosPercent, .5, 3))
    fmt.Print("\n")
    fmt.Print(Round(NegPercent, .5, 3))
    fmt.Print("\n")
    fmt.Print(Round(ZeroPercent, .5, 3))
}

