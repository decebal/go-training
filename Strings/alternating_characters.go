package main
import (
    "fmt"
)

func main() {
    var t int
    var testString string

    fmt.Scan(&t)

    for i:=0; i<t; i++ {
        fmt.Scanln(&testString)
        is := createIdeal(testString)
        fmt.Println(is.counter)
    }
}

type idealStringCandidate struct {
    counter int
}

func (is *idealStringCandidate) deleteChar() {
    is.counter ++;
}

func createIdeal(word string) idealStringCandidate {
    var is idealStringCandidate
    current := make(map[int]rune, len(word))
    index := 0
    for _, cRune := range word {
        current[index] = cRune
        index ++
    }
    index = 0
    for key, cRune := range word {
        if (key == 0) {
            continue
        }

        if (cRune == current[index]) {
            is.deleteChar()
        } else {
            index = key
        }
    }
    return is
}
