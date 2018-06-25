package main;

import (
    "fmt";
    "strings";
)

func main() {
    c := strings.Split("Hyderabad,Hanoi", ",");
    city1, city2 := c[0], c[1];
    fmt.Printf("city1 = %v , city2 = %v\n",city1, city2);
}
