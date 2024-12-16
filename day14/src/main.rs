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
    let result = simulate_robots(&mut robots, 101, 103);
    println!("{}", result)
}
