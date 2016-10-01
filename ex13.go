package main

import (
  "fmt"
  "math/rand"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)

  second_main()
}

func second_main() {
  channels := make([]chan bool, 6)
  for i := range channels {
    channels[i] = make(chan bool)
  }
  go func() {
    for {
      channels[rand.Intn(6)] <- true
    }
  }()
  for i := 0; i < 36; i++ {
    var x int
    select {
    case <-channels[0]:
      x = 1
    case <-channels[1]:
      x = 2
    case <-channels[2]:
      x = 3
    case <-channels[3]:
      x = 4
    case <-channels[4]:
      x = 5
    case <-channels[5]:
      x = 6
    }
    fmt.Printf("%d ", x)
  }
  fmt.Println()
}
