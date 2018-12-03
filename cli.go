
package main

import (
  "flag"
  "fmt"
)

var name = flag.String("name", "world", "a name to say hello to.")

var spanish bool

func init() {
  flag.BoolVar(&spanish, "spanish", false, "use spanish lang")
  flag.BoolVar(&spanish, "s", false, "use spanish lang")
}

func main() {
  flag.Parse()
  if spanish == true {
    fmt.Printf("hola %s!\n", *name)
  } else {
    fmt.Printf("hello %s!\n", *name)
  }
}
