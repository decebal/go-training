package main
import "fmt"

func absDif (x byte, y byte) byte {
    if (x < y) {
        return y - x
    }

    return x - y
}

func main() {
    var t int
    var s []byte

    fmt.Scan(&t)
    if t < 1 || t > 100 {
        return
    }

    for i:=0; i < t ; i++ {
        fmt.Scan(&s)
        var r []byte
        for j := len(s) - 1 ; j > -1; j-- {
            r = append(r, s[j])
        }
        var flag int

        for j:=1; j < len(s); j++ {
            if absDif(s[j],s[j-1]) != absDif(r[j], r[j-1]) {
                flag = 1
                fmt.Println("Not Funny")
                break
            }
        }
        if flag == 0 {
            fmt.Println("Funny")
        }
    }
}
