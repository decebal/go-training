package main
import "fmt"

func main() {
    var n,m int
    fmt.Scan(&n, &m)
    limit := m - n
    max := -1
    for i := 0; i <= limit; i++ {
        for j := 0; j <= limit; j++ {
            xor := (n+i)^(n+j)
            if max < xor {
                max = xor
            }
        }
    }
    fmt.Println(max)
}
