use std::collections::{HashSet, LinkedList};

use crate::point::Point;

const CHAIN_LENGTH: i32 = 10;

#[derive(Clone, Copy, Debug)]
struct Change {
    old: Point,
    new: Point,
}

pub fn part_02(input: &str) -> i32 {
    let mut tail_positions = HashSet::new();
    let mut chain: LinkedList<Point> = LinkedList::new();
    for _ in 0..CHAIN_LENGTH {
        chain.push_back(Point { x: 0, y: 0 });
    }
    for line in input.lines() {
        let (direction, amount_str) = line.split_once(" ").unwrap();
        let amount: i32 = amount_str.parse().unwrap();
        for _ in 0..amount {
            let head = chain.back_mut().unwrap();
            let mut change = Change {
                old: head.clone(),
                new: Point { x: 0, y: 0 },
            };
            match direction {
                "U" => head.y += 1,
                "D" => head.y -= 1,
                "R" => head.x += 1,
                "L" => head.x -= 1,
                _ => panic!("Invalid instruction! {}", direction),
            }
            change.new.x = head.x;
            change.new.y = head.y;
            let mut latest_change = Some(change);
            for point in chain.iter_mut().rev().skip(1) {
                match latest_change {
                    Some(c) => {
                        if point.touches(&c.new) {
                            continue;
                        }
                        let mut new_change = Change {
                            old: point.clone(),
                            new: Point { x: 0, y: 0 },
                        };
                        point.x = c.old.x;
                        point.y = c.old.y;
                        new_change.new.x = c.old.x;
                        new_change.new.y = c.old.y;
                        latest_change = Some(new_change);
                    }
                    None => break,
                }
            }
            let tail = chain.front().unwrap();
            tail_positions.insert(format!("{}x{}", tail.x, tail.y));
        }
    }
    println!("{:?}", chain);
    tail_positions.len() as i32
}
