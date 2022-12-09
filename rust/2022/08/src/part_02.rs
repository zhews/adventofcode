use crate::parser::parse_map;

pub fn part_02(input: &str) -> usize {
    let map = parse_map(&input);
    let mut scenic_scores = Vec::new();
    for (row_index, row) in map.iter().enumerate() {
        for (tree_index, tree) in row.iter().enumerate() {
            let mut left = 0;
            let mut right = 0;
            let mut up = 0;
            let mut down = 0;
            // to left
            for tree_index_left in (0..tree_index).rev() {
                let tree_left = row.get(tree_index_left).unwrap();
                left += 1;
                if tree_left >= tree {
                    break
                }
            }
            // to right
            for tree_index_right in tree_index+1..row.len() {
                let tree_right = row.get(tree_index_right).unwrap();
                right += 1;
                if tree_right >= tree {
                    break
                }
            }
            // above
            for tree_index_up in (0..row_index).rev() {
                let tree_up = map.get(tree_index_up).unwrap().get(tree_index).unwrap();
                up += 1;
                if tree_up >= tree {
                    break
                }
            }
            // below
            for tree_index_down in row_index+1..map.len() {
                let tree_down = map.get(tree_index_down).unwrap().get(tree_index).unwrap();
                down += 1;
                if tree_down >= tree {
                    break
                }
            }
            scenic_scores.push(left * right * up * down);
        }
    }
    *scenic_scores.iter().max().unwrap() as usize
}
