use std::env;

pub fn repeating_characters()->i32{
    let args: Vec<String> = env::args().collect();
    let input: &String = &args[1];
    
    let mut ans: i32 = 0;
    let mut count: i32 = 0;
    let mut char_buffer: Vec<char> = Vec::new();
    for input_char in input.chars(){
        count+=1;
        if char_buffer.contains(&input_char){
            char_buffer = char_buffer[char_buffer.iter().position(|&r| r==input_char).unwrap()+1..].to_vec();
            char_buffer.push(input_char);
        }else if char_buffer.len() == 13{
            ans = count;
            break;
        }else{
            char_buffer.push(input_char);
        }
    }

    return ans;
}
