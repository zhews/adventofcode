use std::collections::HashSet;

pub fn detect_distinct_character_marker(input: &str, n: usize) -> i32 {
    for stream_location in 0..input.len() - n - 1 {
        let data = &input[stream_location..stream_location + n];
        let mut set: HashSet<char> = HashSet::new();
        for c in data.chars() {
            set.insert(c);
        }
        if set.len() == n {
            return stream_location as i32 + n as i32;
        }
    }
    -1
}
