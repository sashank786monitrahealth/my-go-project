package main

import (
    "fmt"
    //"bytes"
)

func div(inp0, inp1 int) (int, int) {
    rem := inp0;
    quo := 0;

    for rem >= inp1 {
        quo++;
        rem = rem - inp1;
    }
    return quo,rem;
}
//type myString string;


func main(){

quotient, remainder := div(71,5);
fmt.Printf("div(71,5) -> quotient = %d , remainder = %d\n",quotient,remainder);

}
