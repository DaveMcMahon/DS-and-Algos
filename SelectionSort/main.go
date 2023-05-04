package main

import (
  "fmt"
)

func findSmallest(a []int) int {
  s := a[0]
  sIdx := 0

  for i, e := range a {
    if e < s {
      s = e
      sIdx = i
    }
  }
  return sIdx
}

func selectionSort(items []int){

  for i := 0; i < len(items); i++ {
    var minIdx = i

    for j := i; j < len(items); j++ {
      if items[j] < items[minIdx]{
        minIdx = j
      }
    }

    items[i], items[minIdx] = items[minIdx], items[i]
  } 
}

func main(){
  
  myArr := []int {5, 3, 6, -1, 2, 11, 1}
  selectionSort(myArr)
  fmt.Println(myArr)
}
