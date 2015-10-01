package main
import "fmt"

func main() {
    var T, N, K int

    fmt.Scan(&T)

    if T < 1 || T > 10 {
        return
    }

    for i := 0 ; i < T; i++ {
        fmt.Scan(&N, &K)
        if N < 1 || N > 1000 || K < 1 || K > N {
            return
        }
        presence := make([]int, N)
        var counter int
        for j := 0; j < N; j++ {
            fmt.Scan(&presence[j])
            if presence[j] <= 0 {
                counter ++
            }
        }
        if counter < K {
            fmt.Println("YES")
        } else {
            fmt.Println("NO")
        }
    }

}
