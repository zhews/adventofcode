use crate::parser::parse_ranges;

pub fn part_02(input: &str) -> i32 {
    parse_ranges(input)
        .iter()
        .filter(|(first, second)| first.contains(&second.start()) || first.contains(&second.end()) || second.contains(&first.start()) || second.contains(&first.end()))
        .count() as i32
}
