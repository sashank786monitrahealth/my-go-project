package main

import (
  "fmt"
)

func printer(msg string , loopCount int) {
  for i :=0; i < loopCount; i++ {
  fmt.Printf("Inside func printer. The message is: %s \n",msg);
  }
}

func main() {

   printer("Hello World.",10);

}
