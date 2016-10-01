package main

import (
  "fmt"
  "os"
  "strings"
)

func main() {
  name := "liuyihan"
  if len(os.Args) > 1{
    // os.Args[0] == file_name
    name = strings.Join(os.Args[1:], " ")
  }
  fmt.Println("Name: ", name)
}
