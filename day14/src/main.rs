use regex::bytes::Regex;

struct Robot {
    x: i32,
    y: i32,
    v_x: i32,
    v_y: i32,
}

impl Robot {
    fn tick(&mut self, iterations: i32, width: i32, height: i32) {
        self.x = (self.x + self.v_x * iterations).rem_euclid(width);
        self.y = (self.y + self.v_y * iterations).rem_euclid(height);
    }
}

fn parse_input(input: String) -> Vec<Robot> {
    let mut robots = Vec::<Robot>::new();
    let num_regex = Regex::new(r"(-?\d)+").unwrap();

    let lines = input.split('\n');
    for line in lines {
        let numbers: Vec<i32> = num_regex
            .find_iter(line.as_bytes())
            .map(|str| {
                std::str::from_utf8(str.as_bytes())
                    .unwrap()
                    .parse::<i32>()
                    .unwrap()
            })
            .collect();

        if numbers.len() == 4 {
            let robot = Robot {
                x: numbers[0],
                y: numbers[1],
                v_x: numbers[2],
                v_y: numbers[3],
            };
            robots.push(robot);
        }
    }

    return robots;
}

fn get_quadrant(x: i32, y: i32, width: i32, height: i32) -> i32 {
    let center_x = width / 2;
    let center_y = height / 2;

    return if x < center_x && y < center_y {
        1
    } else if x > center_x && y < center_y {
        2
    } else if x < center_x && y > center_y {
        3
    } else if x > center_x && y > center_y {
        4
    } else {
        0
    };
}

fn construct_grid(robots: &Vec<Robot>, width: i32, height: i32) -> Vec<Vec<i32>> {
    let mut grid: Vec<Vec<i32>> = vec![vec![0; width as usize]; height as usize];

    robots.iter().for_each(|robot| {
        grid[robot.y as usize][robot.x as usize] += 1;
    });

    return grid;
}

fn print_grid(grid: &Vec<Vec<i32>>) {}

fn step_robots(robots: &mut Vec<Robot>, width: i32, height: i32) {
    let mut amt_in_row = 0;
    let mut found_tree = false;
    let mut i = 0;
    while !found_tree {
        robots.iter_mut().for_each(|robot| {
            robot.tick(1, width, height);
        });

        let grid = construct_grid(robots, width, height);
        grid.iter().for_each(|row| {
            row.iter().for_each(|n| {
                if !found_tree {
                    if *n == 0 {
                        amt_in_row = 0;
                    } else {
                        amt_in_row += 1;
                    }

                    found_tree = amt_in_row >= 7;
                }
            });
        });

        i += 1;
    }
    println!("{}", i);
}

fn simulate_robots(robots: &mut Vec<Robot>, width: i32, height: i32) -> i32 {
    let mut robot_counts: [i32; 4] = [0; 4];
    robots.iter_mut().for_each(|robot| {
        robot.tick(100, width, height);
        let quadrant = get_quadrant(robot.x, robot.y, width, height);

        if quadrant != 0 {
            robot_counts[(quadrant - 1) as usize] += 1;
        }
    });

    return robot_counts.iter().product();
}

fn main() {
    let file_contents = std::fs::read_to_string("input.txt").unwrap();
    let mut robots = parse_input(file_contents);
    step_robots(&mut robots, 101, 103);
}
