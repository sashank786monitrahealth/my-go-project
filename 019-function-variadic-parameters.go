package main

import (
    "fmt"
)

func avg(myNumArray ...float64) float64{
    n := len(myNumArray);
    t := 0.0;

    for _,v := range myNumArray {
        t += v;
    }

    return t/float64(n);
}

func sum(myNumArray ...float64) float64 {
    var sum float64;
    for _, v := range myNumArray {
        sum += v;
    }
    return sum;
}

func main(){

fmt.Printf("avg([1.0,2.1,3.2,4.3,5.4,6.5,7.6]) = %.2f\n",avg(1.0,2.1,3.2,4.3,5.4,6.5,7.6));
myFloatingPointArray := []float64{1.0,2.1,3.2,4.3,5.4,6.5,7.6,8.7,9.8};
fmt.Printf("sum(%v) = %.2f\n",myFloatingPointArray,sum(myFloatingPointArray...));

}
