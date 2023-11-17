use std::{env, fs, string::String};

pub fn crate_arrangement() -> String{
    let args: Vec<String> = env::args().collect();

    let read_path: &String = &args[1];

    let contents = fs::read_to_string(read_path).expect(
        &format!("Should have been able to read file from {}", read_path)
        );
    
    let mut finding_start = true;
    let mut start_lines: Vec<&str> = Vec::new();
    let mut crates: Vec<Vec<char>> = Vec::new();
    contents.lines().into_iter().for_each(|s|{
        if s == ""{
            while crates.len() < start_lines[start_lines.len()-1].split(" ")
                                    .filter(|string| string.len()>0).count()
            {
                crates.push(Vec::new());
            }

            start_lines.iter().rev().skip(1).for_each(|line|{
                for (pos, c) in line.chars().enumerate(){
                    if pos%4==1 && c!=' '{
                        crates[pos/4].push(c);
                    }
                }
            });

            finding_start = false;
        }else if finding_start{
            start_lines.push(s);
        }else{
            let command: Vec<&str> = s.split(" ").collect();
            let count = command[1].parse::<usize>().unwrap();
            let from = command[3].parse::<usize>().unwrap()-1;
            let to = command[5].parse::<usize>().unwrap()-1;
            
            let mut source = crates[from][crates[from].len()-count..].to_vec();
            source.reverse();
            crates[to].extend(source.iter().cloned());
            let length = crates[from].len()-count;
            crates[from].truncate(length);
        }
    });
    
    let mut result: String = String::new();
    crates.iter().for_each(|c|{
        result.push(c.last().unwrap().clone());
    });
    return result;
}

#[allow(dead_code)]
fn print_crates(crates: &Vec<Vec<char>>){
    let mut i = 1;
    crates.iter().for_each(|crate_line|{
        let string: String = crate_line.iter().collect();
        println!("{} {}", i, string);
        i+=1;
    });
    println!("");
}
