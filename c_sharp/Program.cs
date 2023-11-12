using System;
using System.Linq;

namespace MyApp
{
    internal class Program
    {
        static void Main(string[] args)
        {
            var readPath = args[0];

            var lines = File.ReadLines(readPath);
            
            int[] largest_elf_bags = {0,0,0};
            var elf_bag_buffer = 0;

            foreach (var line in lines){
                if (line == ""){
                    for (var i=0 ; i<largest_elf_bags.Length ; i++){
                        if (elf_bag_buffer > largest_elf_bags[i]){
                            for (var j=largest_elf_bags.Length-2 ; j>=i ; j--){
                                largest_elf_bags[j+1] = largest_elf_bags[j];
                            }
                            largest_elf_bags[i] = elf_bag_buffer;
                            break;
                        }
                    }
                    elf_bag_buffer = 0;
                    continue;
                }
                if (int.TryParse(line, out var fooditem)){
                        elf_bag_buffer += fooditem;
                }
            }
            if (elf_bag_buffer!=0){
                for (var i=0 ; i<largest_elf_bags.Length ; i++){
                    if (elf_bag_buffer > largest_elf_bags[i]){
                        for (var j=largest_elf_bags.Length-2 ; j>=i ; j--){
                            largest_elf_bags[j+1] = largest_elf_bags[j];
                        }
                        largest_elf_bags[i] = elf_bag_buffer;
                    }
                }
            }

            Console.WriteLine(largest_elf_bags.Sum());
        }
    }
}
