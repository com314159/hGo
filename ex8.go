package main

import (
  "fmt"
)

// 唯一的区别其实在于 [] 中的内容，如果有，则是数组，如果没有，则是切片
func main() {
  var a [2]string
  a[0] = "hello"
  a[1] = "world"
  fmt.Println(a[0], a[1])
  fmt.Println(a)

  // 声明数组
  cities := [...]string{"Shanghai", "Guangzhou", "Shenzhen", "Beijing"}
  fmt.Printf("%-8T %2d %q\n", cities, len(cities), cities)

  // 声明切片
  s := []string{"A", "B", "C", "D", "E", "F", "G"}
  t := s[:5]
  u := s[3:len(s)]
  fmt.Println(s, t, u)
  u[1] = "x"
  fmt.Println(s, t, u)
  // 遍历切片
  for i, letter := range s {
    fmt.Println(i, letter)
  }

  ab := make([]int, 5)
  printSlice("ab", ab)
  bb := make([]int, 0, 5)
  printSlice("bb", bb)
  c := bb[:2]
  printSlice("c", c)
  d := c[2:5]
  printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
