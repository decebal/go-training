package main
import (
    "fmt"
    "time"
    "bufio"
    "os"
    "strconv"
)

const (
    longForm = "2 1 2006"
)

func main() {
    var lines []string
    var fine float64

    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    newFormatStart, _ := time.Parse(longForm, lines[0])
    newFormatEnd, _ := time.Parse(longForm, lines[1])

    duration := newFormatStart.Sub(newFormatEnd)

    if (duration.Hours() > 0) {
        if (newFormatStart.Year() != newFormatEnd.Year()) {
            fine = 10000
        } else if (newFormatStart.Month() != newFormatEnd.Month()) {
            delivered, err := strconv.Atoi(newFormatEnd.Format("1"))
            if (err != nil) {
                fmt.Println(err)
            }
            expected, err := strconv.Atoi(newFormatStart.Format("1"))
            if (err != nil) {
                fmt.Println(err)
            }
            fine = 500 * float64(expected - delivered)
        } else {
            fine = 15 * duration.Hours() / 24
        }
    }

    fmt.Println(fine)
}
