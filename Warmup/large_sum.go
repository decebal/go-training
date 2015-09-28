package main
import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
    "math/big"
)


func main() {
    var N int
    var lines []string
    sum := big.NewInt(0)

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
        member := big.NewInt(0)

        _, ok := member.SetString(element, 10)
        if !ok {
            fmt.Printf("Error: %s\n", err)
        }
        sum.Add(sum, member)
    }

    fmt.Print(sum)
}
