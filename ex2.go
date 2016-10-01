package main

import "fmt"

func main() {
  user_name := "liuyihan"
  fmt.Println("You big name is", user_name, &user_name)

  fmt.Println("Two hour has", 2 * 60, "seconds")

  flag := true
  fflag := false
  fmt.Println("is True", flag, fflag)
}
