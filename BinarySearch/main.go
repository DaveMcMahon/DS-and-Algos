package main

import (
  "fmt"
)

// O(log n)

func main(){

  l := []int {1,3,5,7,9};
  
  i, ok := bSearch(l, 45)
  if !ok {
    fmt.Println("Item not in the list")
    return
  }

  fmt.Println(i)
}

func bSearch(arr []int, e int) (int, bool) {
  low := 0    // Set the low index 
  high := len(arr) - 1 // Set the high index

  for low <= high {           // Keep dividing up the upper or lower portion of the array
    mid := (low + high) / 2   // Split the resulting array in half
    guess := arr[mid]         // Use that value as the index to the array

    if guess == e {           // If the value equals the item, returrn
      return mid, true
    }

    if guess > e {            // Otherwise, check if we need to move into the upper portion,
      high = mid - 1          // or the lower portion of the halves
    } else {
      low = mid + 1
    }
  }
  return 0, false                 // Return 0 and false if the item isnt found
