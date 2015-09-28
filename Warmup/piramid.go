package main
import (
    "fmt"
    "strings"
)

func main() {
    var n int
    fmt.Scan(&n)
    for i:=n-1; i>=0; i-- {
        line := []string{strings.Repeat(" ", i), strings.Repeat("#", (n-i))}
        fmt.Println(strings.Join(line, ""))
    }
}
