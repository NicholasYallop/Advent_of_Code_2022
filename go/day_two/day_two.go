package day_two

import (
    "bufio"
    "os"
    "strings"
)


func RockPaperScissors() (result int){
    theirchoices := map[string]int{
        "A": 0, // rock
        "B": 1, // paper
        "C": 2, // scissor
    }
    mychoices := map[string]int{
        "X": 0, // rock
        "Y": 1, // paper
        "Z": 2, // scissor
    }

    readPath := os.Args[1]
    file, err := os.Open(readPath)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    fileScanner := bufio.NewScanner(file)

    fileScanner.Split(bufio.ScanLines)

    for fileScanner.Scan(){
        line := fileScanner.Text()

        choices := strings.Split(line, " ")
        their_choice := theirchoices[choices[0]]
        my_choice := mychoices[choices[1]]

        var result_score int
        if (their_choice == my_choice){
            //draw
            result_score = 3
        }else if (their_choice + 1) % 3 == my_choice{
            // I win
            result_score = 6
        }else{
            //they win
            result_score = 0
        }

        result += result_score + my_choice + 1
    }

    return result
}


