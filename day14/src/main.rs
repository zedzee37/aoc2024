use regex::bytes::Regex;

struct Robot {
    x: i32,
    y: i32,
    v_x: i32,
    v_y: i32,
}

fn parse_input(input: String) -> Vec<Robot> {
    let mut robots = Vec::<Robot>::new();
    let num_regex = Regex::new(r"(-?\d)+").unwrap();

    let lines = input.split('\n');
    for line in lines {
        let numbers: Vec<i32> = num_regex
            .find_iter(line.as_bytes())
            .map(|str| 
                std::str::from_utf8(str.as_bytes()).unwrap().parse::<i32>().unwrap()
            ).collect();

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

    robots
}

fn main() {
    let file_contents = std::fs::read_to_string("input.txt").unwrap();
    let robots = parse_input(file_contents);
}
