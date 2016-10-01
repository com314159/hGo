package main

import (
  "fmt"
)

func main() {
  feature_dict := make(map[string]string)
  feature_dict["11"] = "22"
  feature_dict["22"] = "33"
  feature_dict["33"] = "44"
  fmt.Println(feature_dict)

  v := feature_dict["11"]
  fmt.Println(v)

  for key, value := range(feature_dict) {
    fmt.Println(key, value)
  }
}
