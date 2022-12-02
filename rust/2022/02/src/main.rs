use part_01::part_01;
use part_02::part_02;

const INPUT_FILE_NAME: &str = "input.txt";

mod scores;
mod part_01;
mod part_02;

fn main() {
    let input = std::fs::read_to_string(INPUT_FILE_NAME).unwrap();
    println!("Part01: {}", part_01(&input));
    println!("Part02: {}", part_02(&input));
}
