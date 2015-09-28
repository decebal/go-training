package main
import (
    "fmt"
    "math/big"
)

#TODO use memoization for generating the fibonnaci sequence

func main() {
    var A, B int64
    var n int
    fmt.Scan(&A)
    if A < 0 || A > 2 {
        return
    }
    fmt.Scan(&B)
    if B < 0 || B > 2 {
        return
    }
    fmt.Scan(&n)
    if n < 3 || n > 20 {
        return
    }


    afactor := big.NewInt(A)
    bfactor := big.NewInt(B)

    for i := 2; i < n; i++ {
        multipleB := big.NewInt(1)
        afactor.Add(afactor, multipleB.Mul(bfactor, bfactor))
        afactor, bfactor = bfactor, afactor
    }

    fmt.Println(bfactor)

}




