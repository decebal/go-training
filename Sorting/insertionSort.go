package main
import (
    "fmt"
    "bytes"
)

func insertionSort(limit int, a []int) {
    for i := 1; i < limit ; i++ {
        value := a[i]
        j := i - 1
        for j >= 0 && a[j] > value {
            a[j+1] = a[j]
            j = j - 1
            fmt.Println(strSlice(a))
        }
        a[j+1] = value
    }
    fmt.Println(strSlice(a))
}

func strSlice(fs []int) string {
    var b bytes.Buffer
    for _, f := range fs {
        fmt.Fprintf(&b, "%d ", f)
    }
    return b.String()
}

func main() {
    var n int
    fmt.Scan(&n)
    vector := make([]int, n)
    for i := 0 ; i < n ; i++ {
        fmt.Scan(&vector[i])
    }
    insertionSort(n, vector)
}
