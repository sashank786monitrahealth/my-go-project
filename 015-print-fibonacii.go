package main

import (
    "fmt"
)

func myFibonacii(n int){
    fmt.Printf("fib(%d): [",n);

    var p0,p1 uint64 = 0,1;
    fmt.Printf("%d %d ",p0,p1);
    for i := 2; i<=n; i++ {
        p0,p1 = p1, p0+p1;
        fmt.Printf("%d ",p1);
    }

    fmt.Println("]");

}


func main(){
  myFibonacii(12);
}
