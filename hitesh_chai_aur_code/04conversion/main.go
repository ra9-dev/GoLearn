package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    fmt.Println("Welcome to our pizza app")
    fmt.Println("Please rate our pizza from 1-5")
    reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString(' ')
    fmt.Println("Thanks for rating: ", input)
    numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
    if err != nil {
        fmt.Println("Err: ", err)
    } else {
        fmt.Println("Output: ", numRating+1)
    }
}
