pub fn compute_calorie_totals(input: &str) -> Vec<i32> {
    let mut calorie_total = 0;
    let mut calorie_totals = Vec::new();
    let lines = input.lines();
    let last_index = lines.count();
    let lines = input.lines();
    for (index, line) in lines.enumerate() {
        if line.is_empty() {
            calorie_totals.push(calorie_total);
            calorie_total = 0;
            continue;
        }
        let calories: i32 = line.parse().expect("Invalid");
        calorie_total += calories;
        if index + 1 == last_index {
            calorie_totals.push(calorie_total);
        }
    }
    calorie_totals
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn empty_input() {
        let calorie_totals = compute_calorie_totals("");
        assert_eq!(calorie_totals.len(), 0)
    }

    #[test]
    fn example_input() {
        let calorie_totals = compute_calorie_totals(r#"1000
2000
3000

4000

5000
6000

7000
8000
9000

10000"#);
        let expected = vec![6000, 4000, 11000, 24000, 10000];
        assert_eq!(calorie_totals, expected)
    }
}
