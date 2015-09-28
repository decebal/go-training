package main
import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
)

func add (x int, y int) int {
    return x + y
}

func main() {
    var N, sum int
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
        sum = add(sum, member)
    }

    fmt.Print(sum)
}
