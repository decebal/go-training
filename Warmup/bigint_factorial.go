package main
import (
    "fmt"
    "math/big"
)


func factorial(n *big.Int) (result *big.Int) {
    result = new(big.Int)

    switch n.Cmp(&big.Int{}) {
    case -1, 0:
        result.SetInt64(1)
    default:
        result.Set(n)
        var one big.Int
        one.SetInt64(1)
        result.Mul(result, factorial(n.Sub(n, &one)))
    }
    return
}

func main() {
    var n int64
    fmt.Scan(&n)
    r := big.NewInt(n)

    fmt.Println(factorial(r))
}
