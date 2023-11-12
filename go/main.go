package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main(){
    readPath := os.Args[1]
    fmt.Println(readPath)
    file, err := os.Open(readPath)
    if err != nil{
        panic(err)
    }
    defer file.Close()

    fileScanner := bufio.NewScanner(file)

    fileScanner.Split(bufio.ScanLines)

    var elf_bags [][]int
    var elf_bag_buffer []int
    for fileScanner.Scan() {
        line := fileScanner.Text()

        if (line == ""){
            if len(elf_bag_buffer) > 0 {
                elf_bags = append(elf_bags, elf_bag_buffer)
                elf_bag_buffer = []int{}
            }
        }else{
            data, err := strconv.Atoi((line))
            if err != nil{
                panic(err)
            }

            elf_bag_buffer = append(elf_bag_buffer, data)
        }
    }
    if len(elf_bag_buffer) > 0 {
        elf_bags = append(elf_bags, elf_bag_buffer)
    }

    fullest_bags := []int{0,0,0}
    for _, elf_bag := range elf_bags{
        bag_sum_buffer := 0
        for _, fooditem := range elf_bag{
            bag_sum_buffer += fooditem
        }

        for index := range fullest_bags{
            if (bag_sum_buffer > fullest_bags[index]){
                for i:=len(fullest_bags)-2 ; i>=index ; i--{
                    fullest_bags[i+1] = fullest_bags[i]
                }
                fullest_bags[index] = bag_sum_buffer
                break
            }
        }
    }

    result := 0
    for _, fullest_bag := range fullest_bags{
        println(fullest_bag)
        result += fullest_bag
    }

    fmt.Println(result)
}
