use crate::filesystem::compute_from_command_history;

const MAX_DIRECTORY_SIZE: &&i32 = &&100000;

pub fn part_01(input: &str) -> i32 {
    let filesystem = compute_from_command_history(input);
    filesystem.values().filter(|size| size < MAX_DIRECTORY_SIZE).sum()
}
