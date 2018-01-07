package main

import (
  "fmt" 
  "strconv"
  "os"
)
func main() {
  if b, berr := strconv.Atoi(os.Args[2]); berr == nil {
    if x, err := strconv.ParseUint(os.Args[1],b,64); err == nil {
      sx := strconv.FormatUint(x, 9)
      dx := string(sx[len(sx)-1:])
      fmt.Printf("%T %s\n", dx, dx)
    } else {
      panic(err)
    }
  } else {
    panic(berr)
  }
}
