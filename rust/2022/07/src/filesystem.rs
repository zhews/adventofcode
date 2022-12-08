use std::collections::HashMap;
use std::fmt::format;

pub fn compute_from_command_history(input: &str) -> HashMap<String, i32> {
    let mut current_path: Vec<String> = Vec::new();
    let mut filesystem = HashMap::new();
    for line in input.lines() {
        match &line[0..4] {
            "$ cd" => {
                let directory = parse_change_directory(&line);
                match directory {
                    ".." => drop(current_path.pop()),
                    "/" => {
                        current_path.clear();
                        current_path.push("/".to_string());
                    }
                    _ => current_path.push(format!("/{}", directory)),
                }
            }
            "$ ls" | "dir " => continue,
            _ => {
                let (file_size_output, _) = line.split_once(" ").unwrap();
                let file_size: i32 = file_size_output.parse().unwrap();
                for paths in 0..current_path.len() {
                    let path = current_path[0..=paths].join("");
                    let current_size = filesystem.get(&path);
                    match current_size {
                        Some(size) => drop(filesystem.insert(path, size + file_size)),
                        None => drop(filesystem.insert(path, file_size)),
                    }
                }
            }
        }
    }
    filesystem
}

fn parse_change_directory(command: &str) -> &str {
    &command[5..]
}
