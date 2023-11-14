package day_three

import (
    "bufio"
    "errors"
    "os"
)

func Rucksack_Compartments() (result int){
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
        length := len(line)
        compartment_one := line[0:(length/2)]
        compartment_two := line[(length/2):]

        have_tested := make(map[rune]int)
        for _, char_one := range compartment_one{
            // if we haven't already evaluated 
            if _, ok := have_tested[char_one]; !ok {
                have_tested[char_one] = 1;

                for _, char_two := range compartment_two{
                    if char_one == char_two{
                        i, err := priority(char_one)
                        if err!=nil{
                            panic(err)    
                        }
                        result += i
                        break
                    }
                }    
            }
        }
    }

    return result
}

func priority(r rune) (priority int, err error){
    i := int(r)
    if (i > 65){
        if (i > 97){
            return i-96, nil
        }
        return i-38, nil
    }
    return -1, errors.New("priority cast error")
}
