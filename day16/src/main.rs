use std::{collections::HashSet, fs, };

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

struct AStarCell {
    pos: (i32, i32),
    f: i32,
    g: i32,
    h: i32,
}

fn is_on_grid(grid: &Vec<Vec<char>>, pos: (i32, i32)) -> bool {
    return pos.0 >= 0 || pos.1 >= 0 || pos.1 < grid.len() as i32 || pos.0 < grid.len() as i32;
}

fn get_surrounding_positions(grid: &Vec<Vec<char>>, pos: (i32, i32)) -> Vec<(i32, i32)> {
    let mut surrounding = Vec::<(i32, i32)>::new();

    for dir in CARDINALS {
        let new_pos = add(pos, dir);
        if is_on_grid(grid, new_pos) && get_grid(grid, new_pos) != '#' {
            surrounding.push(new_pos);
        }
    }

    return surrounding;
}

fn parse_input(contents: String) -> Vec<Vec<char>> {
    return contents.split("\n").map(|s| s.chars().collect()).collect();
}

fn find_char(grid: &Vec<Vec<char>>, ch: char) -> Option<(i32, i32)> {
    for y in 0..grid.len() {
        for x in 0..grid.len() {
            if grid[y][x] == ch {
                return Some((x as i32, y as i32));
            }
        }
    }
    return None
}

fn get_lowest_f_cost(cells: &Vec<AStarCell>) -> i32 {
    let mut lowest_f = cells[0].f;
    let mut lowest_cell = 0;

    for i in 0..cells.len() {
        let cell = &cells[i];

        if cell.f < lowest_f {
            lowest_f = cell.f;
            lowest_cell = i as i32;
        }
    }

    return lowest_cell;
}

fn find_shortest_path_cost(grid: &Vec<Vec<char>>) -> i32 {
    let visited = HashSet::<(i32, i32)>::new();
    let start_pos = find_char(grid, 'S').unwrap();
    let end_pos = find_char(grid, 'E').unwrap();

    let mut current = Vec::<AStarCell>::new();
    let h = manhattan_distance(start_pos, end_pos);
    current.push(AStarCell {
        pos: start_pos,
        g: 0,
        h: h,
        f: h,
    });

    while true {
        let lowest_cost_cell_idx = get_lowest_f_cost(&current);
        let lowest_cost_cell = &current[lowest_cost_cell_idx as usize];

        let surrounding = get_surrounding_positions(grid, lowest_cost_cell.pos);
        for neighbor in surrounding {
            let g_cost = lowest_cost_cell.g + 1;        
            let h_cost = manhattan_distance(neighbor, end_pos);
            
            current.push(AStarCell {
                pos: neighbor,
                g: g_cost,
                h: h_cost,
                f: g_cost + h_cost,
            });
        }
    }

    return 0;
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