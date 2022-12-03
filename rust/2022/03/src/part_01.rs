use std::collections::HashSet;

pub fn part_01(input: &str) -> i32 {
    let mut sum = 0;
    for line in input.lines() {
        let midpoint = line.len() / 2;
        let first_compartment: String = line.chars().take(midpoint).collect();
        let second_compartment: String = line.chars().skip(midpoint).collect();
        let mut occurrences = HashSet::new();
        for item in first_compartment.chars() {
            if occurrences.contains(&item) {
                continue;
            }
            if second_compartment.contains(item) {
                if item.is_lowercase() {
                    sum += -96 + item as i32;
                }
                if item.is_uppercase() {
                    sum += -38 + item as i32;
                }
            }
            occurrences.insert(item);
        }
    }
    sum
}
