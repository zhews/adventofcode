use crate::filesystem::compute_from_command_history;

const REQUIRED_FREE_SPACE: i32 = 40000000;

pub fn part_02(input: &str) -> i32 {
    let filesystem = compute_from_command_history(&input);
    let filesystem_size = filesystem.get("/").unwrap();
    let mut smallest_folder_to_delete = i32::MAX;
    for directory_size in filesystem.values() {
        if filesystem_size - directory_size <= REQUIRED_FREE_SPACE && directory_size < &smallest_folder_to_delete {
            smallest_folder_to_delete = *directory_size;
        }
    }
    smallest_folder_to_delete
}
