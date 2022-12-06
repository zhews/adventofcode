use crate::detector::detect_distinct_character_marker;

const START_OF_MESSAGE_MARKER_AMOUNT: usize = 14;

pub fn part_02(input: &str) -> i32 {
    detect_distinct_character_marker(input, START_OF_MESSAGE_MARKER_AMOUNT)
}
