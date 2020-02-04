package main

import "fmt"

func main() {
  var headers map[string]string
  for k, v := range headers {
    fmt.Println(k, v)
  }
}

