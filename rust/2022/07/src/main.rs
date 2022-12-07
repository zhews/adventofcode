use std::collections::HashMap;
use std::string::ToString;

const INPUT_FILE_NAME: &str = "input.txt";

fn main() {
    let input = std::fs::read_to_string(INPUT_FILE_NAME).unwrap();
    println!("{}", part_01(&input));
    println!("{}", part_02(&input));
}

fn part_01(input: &str) -> i32 {
    let mut current_path = Vec::new();
    let mut file_sizes = HashMap::new();
    for line in input.lines() {
        if line.starts_with("$") {
            process_command(line, &mut current_path);
        } else {
            if line.starts_with("dir") {
                continue;
            }
            let (size_str, _) = line.split_once(" ").unwrap();
            let size: i32 = size_str.parse().unwrap();
            for path_length in 0..current_path.len() {
                let path_part = current_path[0..=path_length].join("");
                let current_size = file_sizes.get(&path_part).or(Some(&0)).unwrap();
                file_sizes.insert(path_part, current_size + size);
            }
        }
    }
    file_sizes.values().filter(|size| size < &&100000).sum()
}

fn process_command(command: &str, current_path: &mut Vec<String>) {
    if command == "$ cd /" {
        current_path.clear();
        current_path.push("/".to_string());
    } else if command.starts_with("$ cd") {
        let path: String = command.split_whitespace().rev().take(1).collect();
        match path.as_str() {
            ".." => drop(current_path.pop()),
            _ => current_path.push(path + &"/".to_string())
        }
    }
}

fn part_02(input: &str) -> i32 {
    let mut current_path = Vec::new();
    let mut file_sizes = HashMap::new();
    for line in input.lines() {
        if line.starts_with("$") {
            process_command(line, &mut current_path);
        } else {
            if line.starts_with("dir") {
                continue;
            }
            let (size_str, _) = line.split_once(" ").unwrap();
            let size: i32 = size_str.parse().unwrap();
            for path_length in 0..current_path.len() {
                let path_part = current_path[0..=path_length].join("");
                let current_size = file_sizes.get(&path_part).or(Some(&0)).unwrap();
                file_sizes.insert(path_part, current_size + size);
            }
        }
    }
    let current_disk_size = file_sizes.get("/").unwrap();
    let mut min_dir_size: i32 = 70000000;
    for (i, dir_size) in file_sizes.values().enumerate()  {
        if i == 0 || (current_disk_size - dir_size) <= 40000000 {
            if &min_dir_size > dir_size {
                min_dir_size = *dir_size;
            }
        }
    }
    min_dir_size
}
