use std::collections::HashSet;

use crate::parser::parse_map;

pub fn part_01(input: &str) -> usize {
    let map = parse_map(&input);
    let mut visible = HashSet::new();
    // left to right
    for (row_index, row) in map.iter().enumerate() {
        let mut max_tree_size = -1;
        for (tree_index, tree) in row.iter().enumerate() {
            if max_tree_size < *tree as i32 {
                max_tree_size = *tree as i32;
                visible.insert(format!("{}x{}", row_index, tree_index));
            }
        }
    }
    // right to left
    for (row_index, row) in map.iter().enumerate() {
        let mut max_tree_size = -1;
        for (tree_index, tree) in row.iter().rev().enumerate() {
            if max_tree_size < *tree as i32 {
                max_tree_size = *tree as i32;
                visible.insert(format!("{}x{}", row_index, row.len() - 1 - tree_index));
            }
        }
    }
    // top to bottom
    for tree_index in 0..map.get(0).unwrap().len() {
        let mut max_tree_size: i32 = -1;
        for row_index in 0..map.len() {
            let tree = map.get(row_index).unwrap().get(tree_index).unwrap();
            if max_tree_size < *tree as i32 {
                max_tree_size = *tree as i32;
                visible.insert(format!("{}x{}", row_index, tree_index));
            }
        }
    }
    // bottom to top
    for tree_index in 0..map.get(0).unwrap().len() {
        let mut max_tree_size: i32 = -1;
        for row_index in (0..map.len()).rev() {
            let tree = map.get(row_index).unwrap().get(tree_index).unwrap();
            if max_tree_size < *tree as i32 {
                max_tree_size = *tree as i32;
                visible.insert(format!("{}x{}", row_index, tree_index));
            }
        }
    }
    visible.len()
}
