use core::error;
use std::{fmt::Error, fs, task::Context};

#[derive(Debug)]
enum Movement {
    Left,
    Right,
    Up,
    Down,
}

impl Movement {
    fn get_dir(&self) -> (i32, i32) {
        return match self {
            Self::Left => (-1, 0),
            Self::Right => (1, 0),
            Self::Up => (0, -1),
            Self::Down => (0, 1),
        } 
    }
}

fn parse_grid(contents: &str) -> Vec<Vec<char>> {
    return contents.split("\n").map(|str| str.chars().collect()).collect();
}

fn parse_movements(contents: &str) -> Vec<Movement> {
    return contents
        .chars()
        .filter_map(|ch| match ch {
            '^' => Some(Movement::Up),
            '>' => Some(Movement::Right),
            'v' => Some(Movement::Down),
            '<' => Some(Movement::Left),
            _ => None
        })
        .collect();
}

fn find_robot(grid: &Vec<Vec<char>>) -> Option<(usize, usize)> {
    for y in 0..grid.len() {
        for x in 0..grid.len() {
            if grid[y][x] == '@' {
                return Some((x, y));
            }
        }
    }
    return None;
}

fn simulate_robot(grid: &mut Vec<Vec<char>>, movements: &Vec<Movement>) -> Result<usize, Error> {
    let mut pos = find_robot(grid).unwrap();

    for movement in movements {
        let dir = movement.get_dir();
    }

    return Ok(10)
}

fn main() {
    let file_contents = fs::read_to_string("input.txt").unwrap();
    let split: Vec<&str> = file_contents.split("\n\n").collect();
    let mut grid = parse_grid(split[0]);
    let movements = parse_movements(split[1]);
    let result = simulate_robot(&mut grid, &movements).unwrap();
    println!("{}", result);
}
