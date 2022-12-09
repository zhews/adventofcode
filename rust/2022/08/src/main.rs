use part_01::part_01;

mod parser;
mod part_01;

const INPUT_FILE_NAME: &str = "input.txt";

fn main() {
    let input = std::fs::read_to_string(INPUT_FILE_NAME).unwrap();
    println!("{}", part_01(&input))
}
