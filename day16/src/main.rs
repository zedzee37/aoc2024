use std::fs;

struct AStarCell {
    pos: (i32, i32),
    g: i32,
    h: i32,
    f: i32,
    from: Box<AStarCell>
}

fn parse_input(input: String) -> Vec<Vec<char>> {
    input.split('\n').map(|s| s.chars().collect()).collect()
}

fn find_shortest_path(grid: &Vec<Vec<char>>) -> Vec<(i32, i32)> {
    let path = Vec::<(i32, i32)>::new();

    

    path
}

fn main() {
    let file_contents = fs::read_to_string("input.txt").unwrap();
    let grid = parse_input(file_contents);
    println!("Hello, world!");
}
