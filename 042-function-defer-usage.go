package main

import (
  "fmt"
)

func demoDefer (myTasks ...string) {
  fmt.Printf("<<undeferred>>  typeof(myTasks) = %T\n\n\n",myTasks);
  defer fmt.Println("Task Completed!");
  for _,s := range myTasks {
    defer fmt.Println(s);
  }

  fmt.Println("Starting...");
}

func main() {
  demoDefer(
    "look left",
    "look right",
    "again",
    "look left",
    "now",
    "cross the road!",
  )
}
