package main
import (
    "fmt"
    "strings"
    "bufio"
    "os"
)

func main() {
    var s string

    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        s = strings.ToLower(scanner.Text())
    }

    pan := createCandidate(s)
    fmt.Println(pan.isPangram())
}

type PangramCandidate struct {
    lettersRemaining map[rune]int
}

func (pan PangramCandidate) isPangram() string {
    if (len(pan.lettersRemaining) == 0) {
        return "pangram"
    }
    return "not pangram"
}

func createCandidate(phrase string) PangramCandidate {
    var pan PangramCandidate
    pan.lettersRemaining = make(map[rune]int, 26)
    for ch := 'a'; ch <= 'z'; ch++ {
        pan.lettersRemaining[ch] = 1;
    }
    phrase = strings.ToLower(phrase);
    for _, runeValue := range phrase {
        delete(pan.lettersRemaining, runeValue)
    }

    return pan
}
