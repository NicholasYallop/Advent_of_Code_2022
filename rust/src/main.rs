use std::{env, fs};

fn main() {
    let args: Vec<String> = env::args().collect();

    let read_path: &String = &args[1];

    let contents = fs::read_to_string(read_path).expect(
        &format!("Should have been able to read file from {}", read_path)
        );

    let mut elf_bag_buffer: i32 = 0;
    let mut largest_elf_bags = [elf_bag_buffer; 3];
    contents.lines().into_iter()
        .for_each(|s|
                  {
                      if s == ""{
                                for index in 0..(largest_elf_bags.len()){
                                    if elf_bag_buffer > largest_elf_bags[index]{
                                        let mut subsection = largest_elf_bags[index..].to_vec();
                                        subsection.rotate_right(1);
                                        largest_elf_bags[index..].copy_from_slice(&subsection);
                                        largest_elf_bags[index] = elf_bag_buffer;
                                        break;
                                    }
                                }
                          elf_bag_buffer = 0;
                      }else{
                          elf_bag_buffer += s.parse::<i32>().expect(
                              &format!("Should have been able to parse {}", s)
                              );
                      }
                  }
                 );
    if elf_bag_buffer!=0{
        for index in 0..(largest_elf_bags.len()){
            if elf_bag_buffer > largest_elf_bags[index]{
                let mut subsection = largest_elf_bags[index..].to_vec();
                subsection.rotate_right(1);
                largest_elf_bags[index..].copy_from_slice(&subsection);
                largest_elf_bags[index] = elf_bag_buffer;
                break;
            }
        }
    }
    
    let mut result = 0;
    largest_elf_bags.iter().for_each(
        |i| {result += i}
    );
    println!("{}", result);
}
