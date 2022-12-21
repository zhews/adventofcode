use part_01::part_01;

mod point;
mod part_01;

const INPUT_FILE_NAME: &str = "input.txt";

fn main() {
    let input = std::fs::read_to_string(INPUT_FILE_NAME).unwrap();
    println!("Part 1: {}", part_01(&input));
}

