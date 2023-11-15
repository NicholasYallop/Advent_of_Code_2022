use std::{fs, env, collections::HashMap};

#[allow(dead_code)]
pub fn rucksack_compartments() -> u32{
    let args: Vec<String> = env::args().collect();

    let read_path: &String = &args[1];

    let contents = fs::read_to_string(read_path).expect(
        &format!("Should have been able to read file from {}", read_path)
        );

    let mut result: u32 = 0;
    let mut chars: HashMap<char, ()> = HashMap::new();
    let mut wrapped_elf_counter = -1;
    contents.lines().into_iter().for_each(|s|{
        wrapped_elf_counter = (wrapped_elf_counter+1)%3;
        match wrapped_elf_counter{
            0 => {
                chars = HashMap::new();
                s.chars().into_iter().for_each(|c|{
                    chars.insert(c, ());
                });
            },
            1 => {
                let mut temp_chars: HashMap<char, ()> = HashMap::new();
                s.chars().into_iter().for_each(|c|{
                    if chars.contains_key(&c){
                        temp_chars.insert(c, ());
                    }
                });
                chars = temp_chars;
            },
            2 => {
                let mut temp_chars: HashMap<char, ()> = HashMap::new();
                s.chars().into_iter().for_each(|c|{
                    if chars.contains_key(&c){
                        temp_chars.insert(c, ());
                    }
                });
                let char = temp_chars.iter().nth(0).unwrap();
                println!("{}", *char.0);
                result += priority(*char.0);
            },
            _ => println!("whoops"),
        }
    });
    return result;
}

#[allow(dead_code)]
fn priority(c: char) -> u32{
    let i = c as u32;
    let a = 'a' as u32 - 1;
    if i > a{
        return i - a;
    }
    return i - 'A' as u32 + 27;
}
