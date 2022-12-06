use crate::detector::detect_distinct_character_marker;

const START_OF_PACKET_MARKER_AMOUNT: usize = 4;

pub fn part_01(input: &str) -> i32 {
    detect_distinct_character_marker(input, START_OF_PACKET_MARKER_AMOUNT)
}
