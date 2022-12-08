pub fn parse_map(input: &str) -> Vec<Vec<u32>> {
    let mut map = Vec::new();
    for line in input.lines() {
        let mut row = Vec::new();
        for tree in line.chars() {
            let height = tree.to_digit(10).unwrap();
            row.push(height);
        }
        map.push(row);
    }
    map
}
