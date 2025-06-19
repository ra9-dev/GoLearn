package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    welcome := "Welcome to user input"
    fmt.Println(welcome)

    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Enter any input: ")

    // comma ok \\ err ok
    // input, _ := reader.ReadString('\n')
    // _, err := reader.ReadString('\n')
    input, err := reader.ReadString('\n')
    fmt.Printf("Input is of type: %T with value: %v\n", input, input)
    fmt.Printf("Err is of type: %T with value: %v\n", err, err)

}
