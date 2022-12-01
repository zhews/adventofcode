use crate::calories::compute_calories_totals;

pub fn part_02(input: &str) -> i32 {
    let mut calorie_totals = compute_calories_totals(input);
    calorie_totals.sort();
    calorie_totals.iter().rev().take(3).sum()
}
