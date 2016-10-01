package main

import (
  "fmt"
  "os"
  "bufio"
)

func main() {
  reader := bufio.NewReader(os.Stdin)
  fmt.Printf("Type some things: ")
  if line, err := reader.ReadString('\n'); err != nil {
    fmt.Println("Some things wrong")
  } else {
    fmt.Println("You input", line)
  }
}
