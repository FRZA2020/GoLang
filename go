package main

import (
    "fmt"
    "strconv"
    "strings"
    "os"
   )

func main() {
   /* test string */
   var test_str = "01:10:31"

   /* calling the my_minutes function */
   my_result := my_minutes(test_str)
   fmt.Printf( "Total of minutes (rounded) are : %d\n", my_result )
}

/* function returning the minutes, with seconds rounded to the minute */
func my_minutes(my_str string) int {
   var my_var [3]int
   var err error
	
   temp := strings.Split(my_str, ":")
   
   for i := 0; i < 3; i++ {
	my_var[i], err = strconv.Atoi(temp[i]) 			 
   }

   if err != nil {
        fmt.Println(err)
        os.Exit(2)
   }

   /* rounding of seconds */
   if (my_var[2] > 30){
     my_var[1] = my_var[1] + 1
   } 
   result := my_var[0] * 60 + my_var[1]
   return result 
}
