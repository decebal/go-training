package main
import "fmt"

func SliceIndex(limit int, predicate func(i int) bool) int {
    for i := 0; i < limit; i++ {
        if predicate(i) {
            return i
        }
    }
    return -1
}

func main() {
    var V, n int
    fmt.Scan(&V)
    if V < -1000 || V > 1000 {
        return
    }
    fmt.Scan(&n)
    if n < 1 || n > 1000 {
        return
    }
    
    values := make([]int, n)
    
    for i := 0; i < n; i++ {
        fmt.Scan(&values[i])
    }
    
    fmt.Print(SliceIndex(n, func(i int) bool { return values[i] == V }))
    
}
