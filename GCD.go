package main
import (
    "os"
    "bufio"
    "strconv"
    "strings"
    "fmt"
)

// Please do not remove package declarations because these are used by the autograder. If you need additional packages, then you may declare them above.

func DividesAll(a []int, d int) bool {
    
    
    if d == 0 {
        return false
    }
    
    
        
    for _, value := range a {
        if value % d != 0 {
            return false
        }
    }
    return true
}

func MaxIntegerArray(list []int) int {
    first_value := list[0]
    for i := 1; i < len(list); i++ {
        if list[i] > first_value {
            first_value = list[i]   
        }
    }
    return first_value

}
// Insert your GCDArray() function here, along with any subroutines that you need.
func GCDArray(a []int) int {
    current_GCD := 1
    max_numb := MaxIntegerArray(a)
    for i := 1; i <= max_numb; i++ {
        if DividesAll(a, i) {
            current_GCD = i
        }
    }
        return current_GCD

}
