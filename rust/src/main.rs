use std::{env, fs, cmp};

fn main() {
    let args: Vec<String> = env::args().collect();

    let read_path: &String = &args[1];

    let contents = fs::read_to_string(read_path).expect(
        &format!("Should have been able to read file from {}", read_path)
        );

    let mut largest_elf_bag: i32 = 0;
    let mut elf_bag_buffer: i32 = 0;
    contents.lines().into_iter()
        .for_each(|s|
                  {
                      if s == ""{
                          largest_elf_bag = cmp::max(largest_elf_bag, elf_bag_buffer);
                          elf_bag_buffer = 0;
                      }else{
                          elf_bag_buffer += s.parse::<i32>().expect(
                              &format!("Should have been able to parse {}", s)
                              );
                      }
                  }
                 );
    largest_elf_bag = cmp::max(largest_elf_bag, elf_bag_buffer);

    println!("{}", largest_elf_bag);
}
