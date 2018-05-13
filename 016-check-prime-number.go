package main

import (
    "fmt"
    //"bytes"
)

func checkPrime(n int) bool {

    for i:=2;i<n;i++ {
        if (n%i) == 0 {
            return false;
        }
    }
    return true;

}

//type myString string;


func main(){
    myInput := 37;
    myString  := "";

    if checkPrime(myInput){
        //myString.WriteString("%d",myInput)
        myString = myString+" is a prime number.\n"
    } else {
        //myString.WriteString("%d",myInput)
        myString = myString+" is NOT a prime number.\n"
    }
    fmt.Printf("%d %s",myInput,myString);
}
