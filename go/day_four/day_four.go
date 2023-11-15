package day_four

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Overlapping_Assignments() (result int){
    readPath := os.Args[1]
    file, err := os.Open(readPath)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    fileScanner := bufio.NewScanner(file)

    fileScanner.Split(bufio.ScanLines)

    result = 0
    for fileScanner.Scan(){
        line := fileScanner.Text()
        elves := strings.Split(line, ",")

        player_one := strings.Split(elves[0], "-")
        start_one, err := strconv.Atoi(player_one[0])
        check(err)
        end_one, err := strconv.Atoi(player_one[1])
        check(err)

        player_two := strings.Split(elves[1], "-")
        start_two, err := strconv.Atoi(player_two[0])
        check(err)
        end_two, err := strconv.Atoi(player_two[1])
        check(err)
        
        if start_one == start_two{
            result += 1
        }else if start_one > start_two {
            if end_two >= start_one {result +=1}
        }else{
            if end_one >= start_two {result+=1}
        }
    }

    return result
}

func check (err error) {
    if err!=nil{
        panic(err)
    }
}
