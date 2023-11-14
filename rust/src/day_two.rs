use std::{env, fs};

#[allow(dead_code)]
pub fn rock_paper_scissors() -> i32{
    let args: Vec<String> = env::args().collect();

    let read_path: &String = &args[1];

    let contents = fs::read_to_string(read_path).expect(
        &format!("Should have been able to read file from {}", read_path)
        );

    let mut result: i32 = 0;
    contents.lines().into_iter().for_each(|s|{
        let choices: Vec<&str> = s.split(" ").collect();
        
        let their_choice: i32 = match choices[0].chars().nth(0).unwrap(){
            'A' => 0,
            'B' => 1,
            'C' => 2,
            _ => panic!("tried to turn {} into A, B, C", choices[0].chars().nth(0).unwrap())
        };
        let desired_outcome: i32 = match choices[1].chars().nth(0).unwrap(){
            'X' => 0,
            'Y' => 3,
            'Z' => 6,
            _ => panic!("tried to turn {} into X, Y, Z", choices[0].chars().nth(0).unwrap())
        };

        let my_choice: i32; // 0: rock | 1: paper | 2: scissors
        match desired_outcome{
            0 => my_choice = (their_choice+2)%3,
            3 => my_choice = their_choice,
            6 => my_choice = (their_choice+1)%3,
            _ => my_choice = 0,
        }

        result += my_choice + desired_outcome + 1;
    });

    return result;
}
