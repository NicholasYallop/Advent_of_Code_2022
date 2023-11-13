use std::{env, fs};

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
        let my_choice: i32 = match choices[1].chars().nth(0).unwrap(){
            'X' => 0,
            'Y' => 1,
            'Z' => 2,
            _ => panic!("tried to turn {} into X, Y, Z", choices[0].chars().nth(0).unwrap())
        };

        let result_score: i32;
        if my_choice == their_choice{
            result_score = 3;
        }else if (their_choice+1)%3 == my_choice{
            result_score = 6;
        }else{
            result_score = 0;
        }

        result += result_score + my_choice + 1;
    });

    return result;
}
