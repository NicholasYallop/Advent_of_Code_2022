package day_six

import (
	"fmt"
	"os"
)
func Contains(r rune, arr []rune)(contained bool, index int){
    for array_index, array_char := range arr{
        if array_char == r {return true, array_index}
        if array_char == 0 {return false, array_index}
    }
    return false, -1
}

func Print(arr []rune){
    for _, c := range arr{
        if c==0 {
            fmt.Print("!")
        }else {
            fmt.Print(string(c))
        }
    }
    fmt.Println("")
}

func Repeating_Characters() (result int) {
    line := os.Args[1]

    char_buffer := make([]rune, 14)
    for line_index, line_char := range line{
        if contained, index := Contains(line_char, char_buffer); contained{
            new_buffer := make([]rune, 14)
            for i:=index+1 ; i<len(char_buffer) ; i++{
                if char_buffer[i] == 0{
                    new_buffer[i-index-1] = line_char
                    break
                }else{
                  new_buffer[i-index-1] = char_buffer[i]
                }
            }
            char_buffer = new_buffer
        }else if index!=-1 && index != len(char_buffer)-1 {
            char_buffer[index] = line_char
        }else{
            return line_index + 1
        }
    }

    return -1
}



