package main

import (
  "fmt"
)

func recursion(n int) int {
  if n == 1 {
    return 1
  }
  return n * recursion(n-1)
}

func main(){
  fmt.Println(recursion(3))
}

/*
    recursion(3)
      if 3 == 1  (No)
      return 3 * recursion(2)
          if 2 == 1 (no)
          return 2 * recursion(1)
            if 1 == 1 (yes)
              return 1

      3 * (2 * 1)
      = 6
*/
