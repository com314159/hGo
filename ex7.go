package main

import (
  "fmt"
)

// 数组在 Go 中是按值传递的

func main() {
  i := 9
  j := 5
  product := 0
  fmt.Println("Origin Data:", i, j)
  swapAndProduct1(&i, &j, &product)
  fmt.Println("After swap:", i, j, product)

  i = 9
  j = 5
  i, j, product = swapAndProduct2(i, j)
  fmt.Println("Swapped2:", i, j, product)
}

func swapAndProduct1(x, y, product *int) {
  if *x > *y {
    *x, *y = *y, *x
  }
  *product = *x * *y
}

func swapAndProduct2(x, y int) (int, int, int) {
  if x > y {
    x, y = y, x
  }
  return x, y, x * y
}
