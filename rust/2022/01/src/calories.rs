pub fn compute_calories_totals(input: &str) -> Vec<i32> {
    let mut calories_totals = Vec::new();
    for elf_calories in input.split("\n\n") {
        let calorie_total = elf_calories.lines().map(move |calorie| calorie.parse::<i32>().unwrap()).sum();
        calories_totals.push(calorie_total);
    }
    calories_totals
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn empty_input() {
        let calorie_totals = compute_calories_totals("");
        assert_eq!(calorie_totals.len(), 0)
    }

    #[test]
    fn example_input() {
        let calorie_totals = compute_calorie_totals_2(r#"1000
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
