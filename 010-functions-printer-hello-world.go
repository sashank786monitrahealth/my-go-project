package main

import (
  "fmt"
)

func printer(msg string) {
  fmt.Printf("Inside func printer. The message is: %s \n",msg);
}

func main() {

   printer("Hello World.");

}
