use std::env;

fn main() {
    if let Ok(current_dir) = env::current_dir(){
        println!("{:?}", current_dir);
    }
    println!("Hello, world!");
}
