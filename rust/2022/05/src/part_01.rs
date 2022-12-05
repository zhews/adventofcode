use regex::Regex;
use crate::parser::{parse_containers, parse_solution};

pub fn part_01(input: &str) -> String {
    let (containers_instruction, move_instructions) = input.split_once("\n\n").unwrap();
    let mut containers = parse_containers(containers_instruction);
    let move_regex = Regex::new(r"move (\d+) from (\d+) to (\d+)").unwrap();
    for line in move_instructions.lines() {
        let captures = move_regex.captures(line).unwrap();
        let amount: i32 = captures.get(1).unwrap().as_str().parse().unwrap();
        let from: i32 = captures.get(2).unwrap().as_str().parse().unwrap();
        let to: i32 = captures.get(3).unwrap().as_str().parse().unwrap();
        for _ in 0..amount {
            let container = containers.get_mut(&from).unwrap().pop().unwrap();
            containers.get_mut(&to).unwrap().push(container);
        }
    }
    parse_solution(&mut containers)
}