package day_three

import (
	"bufio"
	"errors"
	"fmt"
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

    wrapped_elf_count := -1
    var chars map[rune]int
    for fileScanner.Scan(){
        line := fileScanner.Text()

        wrapped_elf_count = (wrapped_elf_count+1)%3
        switch wrapped_elf_count{
        case 0:
            chars = make(map[rune]int)
            for _, char := range line{
                chars[char]=1
            }

        case 1:
            temp_chars := make(map[rune]int)
            for _, char := range line{
                if _, ok := chars[char]; ok{
                    temp_chars[char]=1
                }
            }
            chars = temp_chars

        case 2:
            temp_chars := make(map[rune]int)
            for _, char := range line{
                _, ok := chars[char]; 
                _, temp_ok := temp_chars[char];
                if ok && !temp_ok{
                    temp_chars[char]=1
                }
            }
            for c := range temp_chars {
                fmt.Println(string(c));
                p, err := priority(c)
                if err!=nil {
                    panic(err)
                }
            
                result += p
                break
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
