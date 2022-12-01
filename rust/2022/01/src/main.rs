use crate::part_01::part_01;
use crate::part_02::part_02;

mod calories;
mod part_01;
mod part_02;

const INPUT_FILE_NAME: &str = "input.txt";

fn main() {
    let input = std::fs::read_to_string(INPUT_FILE_NAME).expect("No input file found!");
    let input_2 = "1000
2000
3000

4000

5000
6000

7000
8000
9000

10000";
    let elves: Vec<&str> = input_2.split("\n\n").collect();
    println!("{}", elves.join(" - "));
    println!("Part01: {}", part_01(&input));
    println!("Part02: {}", part_02(&input));
}
