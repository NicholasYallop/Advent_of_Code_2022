use std::{fs, env, collections::HashMap};

pub fn rucksack_compartments() -> u32{
    let args: Vec<String> = env::args().collect();

    let read_path: &String = &args[1];

    let contents = fs::read_to_string(read_path).expect(
        &format!("Should have been able to read file from {}", read_path)
        );

    let mut result: u32 = 0;
    contents.lines().into_iter().for_each(|s|{
        let length = s.len();
        let compartment_one = &s[..(length/2)];
        let compartment_two = &s[(length/2)..];
        
        let mut tested: HashMap<char, ()> = HashMap::new();
        compartment_one.chars().into_iter().for_each(|c|{
            if !tested.contains_key(&c){
                tested.insert(c, ());

                for index in 0..length/2{
                    if c == compartment_two.chars().nth(index).unwrap(){
                        result += priority(c);
                        break;
                    }
                }
            }
        });
    });
    return result;
}

fn priority(c: char) -> u32{
    let i = c as u32;
    let a = 'a' as u32 - 1;
    if i > a{
        return i - a;
    }
    return i - 'A' as u32 + 27;
}
