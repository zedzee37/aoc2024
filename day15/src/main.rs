use std::fs;

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

fn parse_wider_grid(contents: &str) -> Vec<Vec<char>> {
    return contents.split("\n").map(|str| {
            let mut chars = Vec::<char>::new();

            for ch in str.chars() {
                match ch {
                    'O' => {
                        chars.push('[');
                        chars.push(']');
                    },
                    '@' => {
                        chars.push('@');
                        chars.push('.');
                    },
                    _ => {
                        chars.push(ch);
                        chars.push(ch);
                    },
                }
            }

            return chars;
    }).collect();
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

fn is_off_grid(grid: &Vec<Vec<char>>, pos: (i32, i32)) -> bool {
    return pos.0 < 0 || pos.1 < 0 || pos.1 >= grid.len() as i32 || pos.0 >= grid[pos.1 as usize].len() as i32;
}

fn simulate_robot(grid: &mut Vec<Vec<char>>, movements: &Vec<Movement>) -> usize {
    let mut pos = find_robot(grid).unwrap();

    for movement in movements {
        let dir = movement.get_dir();
        let next_pos = (pos.0 as i32 + dir.0, pos.1 as i32 + dir.1);

        if is_off_grid(grid, next_pos) {
            continue;
        }

        let ch = grid[next_pos.1 as usize][next_pos.0 as usize];
        
        match ch {
            '.' => {
                grid[pos.1][pos.0] = '.';
                pos = (next_pos.0 as usize, next_pos.1 as usize);
                grid[pos.1][pos.0] = '@';
            },
            '#' => {},
            'O' => {
                // Find the last box pos, or dont move at all.
                let mut last_box_pos = (next_pos.0 + dir.0, next_pos.1 + dir.1);
                let mut can_move = true;

                loop {
                    let cur_ch = grid[last_box_pos.1 as usize][last_box_pos.0 as usize]; 

                    match cur_ch {
                        '.' => break,
                        '#' => {
                            can_move = false;
                            break;
                        },
                        _ => {}
                    }

                    last_box_pos = (last_box_pos.0 + dir.0, last_box_pos.1 + dir.1);
                }

                if can_move {
                    grid[pos.1][pos.0] = '.';
                    pos = (next_pos.0 as usize, next_pos.1 as usize);
                    grid[pos.1][pos.0] = '@';
                    grid[last_box_pos.1 as usize][last_box_pos.0 as usize] = 'O'
                }
            },
            _ => {} 
        }
    }

    let mut total = 0;

    for y in 0..grid.len() {
        for x in 0..grid[y].len() {
            let ch = grid[y][x];

            if ch == 'O' {
                let gps = y * 100 + x;
                total += gps;
            }
        }
    }

    return total;
}

fn main() {
    let file_contents = fs::read_to_string("input.txt").unwrap();
    let split: Vec<&str> = file_contents.split("\n\n").collect();
    let mut grid = parse_grid(split[0]);
    let movements = parse_movements(split[1]);
    let result = simulate_robot(&mut grid, &movements);
    println!("{}", result);

    let wider_grid = parse_wider_grid(split[0]);
    for y in 0..wider_grid.len() {
        for x in 0..wider_grid[y].len() {
            print!("{}", wider_grid[y][x]);
        }   
        println!()
    }
}
