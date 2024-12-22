use std::fs;

use regex::Regex;

#[derive(Debug)]
enum Instruction {
    Adv(i32),
    Bxl(i32),
    Bst(i32),
    Jnz(i32),
    Bxc(i32),
    Out(i32),
    Bdv(i32),
    Cdv(i32),
}

struct Program {
    a: i32,
    b: i32,
    c: i32,
    instructions: Vec<Instruction>,
    output: String,
}

impl Program {
    fn parse_input(input: String) -> Result<Self, String> {
        let num_regex = Regex::new(r"\d+").unwrap();  
        let numbers: Vec<i32> = num_regex
            .find_iter(&input)
            .map(|v| v.as_str().parse::<i32>().unwrap())
            .collect();
        
        let a = numbers[0];
        let b = numbers[1];
        let c = numbers[2];

        let mut instructions = Vec::<Instruction>::new();
        for i in (3..numbers.len()).step_by(2) {
            let num = numbers[i];
            let next = numbers[i + 1];

            instructions.push(match num {
                0 => Instruction::Adv(next),
                1 => Instruction::Bxl(next),
                2 => Instruction::Bst(next),
                3 => Instruction::Jnz(next),
                4 => Instruction::Bxc(next),
                5 => Instruction::Out(next),
                6 => Instruction::Bdv(next),
                7 => Instruction::Cdv(next),
                _ => return Err(String::from("Encountered unkown opcode.")),
            });
        }

        Ok(Self {
            a,
            b,
            c,
            instructions,
            output: String::new(),
        })
    }

    fn get_combo_value(&self, combo: i32) -> i32 {
        return match combo {
            4 => self.a,
            5 => self.b,
            6 => self.c,
            _ => combo
        }
    }

    fn output(&mut self, val: i32) {
        self.output += &val.to_string(); 
        self.output.push(',');
    }

    fn run(&mut self) {
        let mut instruction_pointer: usize = 0; 
        
        while instruction_pointer < self.instructions.len() {
            let instruction = &self.instructions[instruction_pointer];
            let mut jumped = false;

            match instruction {
                Instruction::Adv(combo) => {
                    let v = 2_i32.pow(self.get_combo_value(*combo) as u32);
                    self.a /= v;
                },
                Instruction::Bxl(literal) => self.b ^= *literal,
                Instruction::Bst(combo) => {
                    let v = self.get_combo_value(*combo);
                    self.b = v.rem_euclid(8);
                },
                Instruction::Jnz(literal) => {
                    if self.a != 0 {
                        jumped = true; 
                        instruction_pointer = (*literal / 2) as usize;
                    }
                },
                Instruction::Bxc(_) => self.b ^= self.c,
                Instruction::Out(combo) => {
                    let v = self.get_combo_value(*combo).rem_euclid(8);
                    self.output(v);
                },
                Instruction::Bdv(combo) => {
                    let v = 2_i32.pow(self.get_combo_value(*combo) as u32);
                    self.b = self.a / v;
                },
                Instruction::Cdv(combo) => {
                    let v = 2_i32.pow(self.get_combo_value(*combo) as u32);
                    self.c = self.a / v;
                },
            }

            if !jumped {
                instruction_pointer += 1;
            }
        }
    }
}

fn main() {
    let input = fs::read_to_string("input.txt").unwrap();
    let mut program = Program::parse_input(input).unwrap();
    program.run();
    println!("{}", program.output);
}
