pub fn compute_calorie_totals(input: &str) -> Vec<i32> {
    let mut calorie_total = 0;
    let mut calorie_totals = Vec::new();
    for line in input.lines() {
        if line.is_empty() {
            calorie_totals.push(calorie_total);
            calorie_total = 0;
            continue;
        }
        let calories: i32 = line.parse().expect("Invalid");
        calorie_total += calories;
    }
    calorie_totals
}
