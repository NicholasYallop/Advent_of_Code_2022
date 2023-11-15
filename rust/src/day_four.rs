use std::{fs, env};

pub fn overlapping_assignments() -> usize{
    let args: Vec<String> = env::args().collect();

    let read_path: &String = &args[1];

    let contents = fs::read_to_string(read_path).expect(
        &format!("Should have been able to read file from {}", read_path)
        );

    contents.lines().into_iter().filter(|s|{
        let elves: Vec<&str> = s.split(',').collect();

        let elf_one: Vec<&str> = elves[0].split('-').collect();
        let start_one = elf_one[0].parse::<u32>().unwrap();
        let end_one = elf_one[1].parse::<u32>().unwrap();

        let elf_two: Vec<&str> = elves[1].split('-').collect();
        let start_two = elf_two[0].parse::<u32>().unwrap();
        let end_two = elf_two[1].parse::<u32>().unwrap();

        start_one == start_two
        || (start_one > start_two && end_one <= end_two)
        || (start_one < start_two && end_one >= end_two)
    }).count()
}
