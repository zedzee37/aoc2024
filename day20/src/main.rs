use std::fs;

struct Grid {
    cells: Vec<GridCell>
}

struct GridCell {
    dist_to_start: u32,
    next: usize
}

fn parse_input(file_name: &str) {
    let file_contents = fs::read_to_string(file_name);
}

fn main() {
    parse_input("input.txt");
    println!("Hello, world!");
}
