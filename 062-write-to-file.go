package main;

import (
   "fmt";
   "os"
)

func main(){
   var myString string;
   myString = "This is Line1 \n This is Line2.";

   myFile , err := os.Create("myOutput.txt")
   if err != nil {
      fmt.Printf("Error encountered:%v\n",err);
   }

   fmt.Fprintf(myFile,myString);
   myFile.Close()

}
