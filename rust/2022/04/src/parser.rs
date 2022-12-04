use std::ops::RangeInclusive;

pub fn parse_ranges(input: &str) -> Vec<(RangeInclusive<i32>, RangeInclusive<i32>)> {
    input.lines()
        .map(|line| line.split_once(",").unwrap())
        .map(|(first_range_instruction, second_range_instruction)| {
            let (first_start_str, first_end_str) = first_range_instruction.split_once("-").unwrap();
            let (second_start_str, second_end_str) = second_range_instruction.split_once("-").unwrap();
            let first_start: i32 = first_start_str.parse().unwrap();
            let first_end: i32 = first_end_str.parse().unwrap();
            let second_start: i32 = second_start_str.parse().unwrap();
            let second_end: i32 = second_end_str.parse().unwrap();
            ((first_start..=first_end), (second_start..=second_end))
        })
        .collect()
}
