use crate::calories::compute_calories_totals;

pub fn part_01(input: &str) -> i32 {
    let calorie_totals = compute_calories_totals(input);
    *calorie_totals.iter().max().unwrap()
}
