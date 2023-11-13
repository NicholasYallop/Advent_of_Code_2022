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
    desired_results := map[string]int{
        "X": 0, // I lose
        "Y": 3, // draw
        "Z": 6, // I win
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
        desired_result := desired_results[choices[1]]

        var my_choice int;
        switch desired_result{
        case 0:
            my_choice = (their_choice+2)%3
        case 3:
            my_choice = their_choice
        case 6:
            my_choice = (their_choice+1)%3
        }

        result += my_choice + desired_result + 1
    }

    return result
}


