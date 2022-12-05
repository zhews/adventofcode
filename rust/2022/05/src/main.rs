use std::collections::HashMap;
use regex::Regex;
use crate::part_01::part_01;
use crate::part_02::part_02;

mod parser;
mod part_01;
mod part_02;

const INPUT_FILE_NAME: &str = "input.txt";

fn main() {
    let input = std::fs::read_to_string(INPUT_FILE_NAME).unwrap();
    println!("{}", part_01(&input));
    println!("{}", part_02(&input));
}
