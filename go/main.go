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
    
    var fullest_bag_sum int
    for _, elf_bag := range elf_bags{
        bag_sum_buffer := 0
        for _, fooditem := range elf_bag{
            bag_sum_buffer += fooditem
        }
        fullest_bag_sum = max(fullest_bag_sum, bag_sum_buffer)
    }

    fmt.Println(fullest_bag_sum)
}
