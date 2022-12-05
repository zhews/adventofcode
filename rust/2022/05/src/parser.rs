use std::collections::HashMap;

pub fn parse_solution(containers: &mut HashMap<i32, Vec<char>>) -> String {
    let mut result = String::new();
    for row in 1..=containers.len() {
        result += &containers.get_mut(&(row as i32)).unwrap().pop().unwrap().to_string();
    }
    result
}

pub fn parse_containers(input: &str) -> HashMap<i32, Vec<char>> {
    let mut containers: HashMap<i32, Vec<char>> = HashMap::new();
    for line in input.lines().rev().skip(1) {
        let mut row = 1;
        for container in (2..line.len()).step_by(4) {
            let container_name = line.chars().nth(container-1).unwrap();
            if container_name != ' ' {
                match containers.get_mut(&row) {
                    Some(stack) => {
                        stack.push(container_name);
                    }
                    None => {
                        let mut stack = Vec::new();
                        stack.push(container_name);
                        containers.insert(row, stack);
                    }
                }
            }
            row  += 1;
        }
    }
    containers
}
