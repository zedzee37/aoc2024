use std::{fs, os::unix::fs::FileExt};

fn add(a: (i32, i32), b: (i32, i32)) -> (i32, i32) {
    return (a.0 + b.0, a.1 + b.1);
}

fn sub(a: (i32, i32), b: (i32, i32)) -> (i32, i32) {
    return (a.0 - b.0, a.1 - b.1); 
}

fn neg(a: (i32, i32)) -> (i32, i32) {
    return (-a.0, -a.1);
}

fn mul(a: (i32, i32), n: i32) -> (i32, i32) {
    return (a.0 * n, a.1 * n);
}

fn div(a: (i32, i32), n: i32) -> (i32, i32) {
    return (a.0 / n, a.1 / n);
}

fn to_usize(a: (i32, i32)) -> (usize, usize) {
    return (a.0 as usize, a.1 as usize);
}

fn manhattan_distance(a: (i32, i32), b: (i32, i32)) -> i32 {
    return (a.0 - b.0).abs() + (a.1 - b.1).abs();
}

fn get_grid(grid: &Vec<Vec<char>>, pos: (i32, i32)) -> char {
    let fixed = to_usize(pos);
    return grid[fixed.1][fixed.0];
}

fn set_grid(grid: &mut Vec<Vec<char>>, pos: (i32, i32), val: char) {
    let fixed = to_usize(pos);
    grid[fixed.1][fixed.0] = val;
}

const CARDINALS: [(i32, i32); 4] = [
    (1, 0),
    (-1, 0),
    (0, 1),
    (0, -1),
];

fn get_surrounding_positions(grid: &Vec<Vec<char>>, pos: (i32, i32)) -> Vec<(i32, i32)> {
    let mut surrounding = Vec::<(i32, i32)>::new();

    for dir in CARDINALS {
        if get_grid(grid, add(pos, dir)) != '#' {
            surrounding.push(add(pos, dir));
        }
    }

    return surrounding;
}

fn parse_input(contents: String) -> Vec<Vec<char>> {
    return contents.split("\n").map(|s| s.chars().collect()).collect();
}

fn main() {
    let file_contents = fs::read_to_string("input.txt").unwrap();
    let grid = parse_input(file_contents);
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_manhattan() {
        assert_eq!(manhattan_distance((0, 1), (3, 2)), 4);
    }
}