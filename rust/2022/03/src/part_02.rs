use std::collections::HashSet;

pub fn part_02(input: &str) -> i32 {
    let mut sum = 0;
    let mut lines = input.lines();
    while let (Some(first_elf), Some(second_elf), Some(third_elf)) = (lines.next(), lines.next(), lines.next()) {
        let mut occurrences = HashSet::new();
        for item in first_elf.chars() {
            if occurrences.contains(&item) {
                continue;
            }
            if second_elf.contains(item) && third_elf.contains(item) {
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
