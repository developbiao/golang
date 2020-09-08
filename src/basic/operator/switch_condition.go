package main

import "fmt"

func main() {
    var grade string = "B"
    var marks int = 90

    switch marks {
       case 90: grade = "A"
       case 80: grade = "B"
       case 50, 60, 70: grade = "C"
       default: grade="D"
    }

    switch {
        case grade == "A" :
           fmt.Println("Good!")
        case grade == "B", grade == "C" :
           fmt.Println("Normal")
       case grade == "D":
           fmt.Println("Just Pass")
        default:
           fmt.Println("Low") 
    }

    fmt.Printf("Your level is %s\n", grade)


}
