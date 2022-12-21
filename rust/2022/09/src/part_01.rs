use crate::point::Point;
use std::collections::HashSet;

pub fn part_01(input: &str) -> i32 {
    let mut head = Point { x: 0, y: 0 };
    let mut tail = Point { x: 0, y: 0 };
    let mut tail_positions = HashSet::new();
    tail_positions.insert(format!("{}x{}", tail.x, tail.y));
    for line in input.lines() {
        let (direction, amount_str) = line.split_once(" ").unwrap();
        let amount: i32 = amount_str.parse().unwrap();
        for _ in 0..amount {
            let next_point: Point;
            match direction {
                "U" => {
                    next_point = Point {
                        x: head.x,
                        y: head.y + 1,
                    }
                }
                "D" => {
                    next_point = Point {
                        x: head.x,
                        y: head.y - 1,
                    }
                }
                "R" => {
                    next_point = Point {
                        x: head.x + 1,
                        y: head.y,
                    }
                }
                "L" => {
                    next_point = Point {
                        x: head.x - 1,
                        y: head.y,
                    }
                }
                _ => panic!("invalid direction! {}", direction),
            }
            if !tail.touches(&next_point) {
                tail.x = head.x;
                tail.y = head.y;
            }
            head = next_point;
            tail_positions.insert(format!("{}x{}", tail.x, tail.y));
        }
    }
    tail_positions.len() as i32
}
