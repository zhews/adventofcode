use crate::parser::parse_map;

mod parser;

const INPUT_FILE_NAME: &str = "input.txt";

fn main() {
    let input = std::fs::read_to_string(INPUT_FILE_NAME).unwrap();
    println!("{:?}", parse_map(&input))
}
