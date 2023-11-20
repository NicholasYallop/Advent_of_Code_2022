package day_five

import (
	"advent_of_code/helpers"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Crate_Arrangement() (result string){
    readPath := os.Args[1]
    file, err := os.Open(readPath)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    fileScanner := bufio.NewScanner(file)

    fileScanner.Split(bufio.ScanLines)

    finding_start := true 
    start_state_line_buffer := []string{}
    var start_state [][]rune
    for fileScanner.Scan(){
        line := fileScanner.Text()

        if line=="" {
            columns := 0
            splits := strings.Split(start_state_line_buffer[len(start_state_line_buffer)-1], " ")
            for i:=0 ; i<len(splits) ; i++{
                if splits[i]!="" {columns++}
            }

            start_state = make([][]rune,columns)

            for i := len(start_state_line_buffer)-2 ; i>=0 ; i-- {
                for j := 0 ; j<len(start_state_line_buffer[i]); j++{
                    if (j%4 == 1){
                        candidate := start_state_line_buffer[i][j]
                        if (candidate!=' '){
                            start_state[j/4] = append(start_state[j/4], rune(start_state_line_buffer[i][j]))
                        }
                    }
                }
            }

            finding_start=false
            continue
        }

        if (finding_start){
            start_state_line_buffer = append(start_state_line_buffer, line)
        }else{
            command := strings.Split(line, " ")
            count, err := strconv.Atoi(command[1])
            helpers.Check(err)

            from_column, err := strconv.Atoi(command[3])
            helpers.Check(err)
            from := from_column-1

            to_column, err := strconv.Atoi(command[5])
            helpers.Check(err)
            to := to_column-1
            
            start_state[to] = append(
                start_state[to], 
                start_state[from][len(start_state[from])-count:]...
            )
            start_state[from] = start_state[from][:len(start_state[from])-count]
        }

    }

    result = ""
    for i:=0 ; i<len(start_state) ; i++{
        result += string(start_state[i][len(start_state[i])-1])
    }

    return result
}

func print_state(state [][]rune){
    for i:=0 ; i<len(state) ; i++{
        fmt.Print(strconv.Itoa(i+1) + " ")
        for j:=0 ; j<len(state[i]) ; j++{
            fmt.Printf(string(state[i][j]))
            fmt.Print(" ")
        }
        fmt.Println("")
    }
    fmt.Println("")
}
