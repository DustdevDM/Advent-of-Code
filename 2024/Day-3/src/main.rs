use regex::Regex;
use std::fs;
use std::path::Path;

fn main() {
    let mut input = fs::read_to_string(Path::new("./src/input.txt"))
        .expect("Should have been able to read the file");
    input = input.trim_start_matches('\u{feff}').to_string();

    let instructions = parse_input(&input);

    let result: i32 = instructions
        .iter()
        .map(|instruction| instruction[0] * instruction[1])
        .sum();

    println!("Result: {result}");
}

fn parse_input(input: &str) -> Vec<Vec<i32>> {
    let regex = Regex::new(r"(mul\(\d{1,4},\d{1,4}\))|do\(\)|don't\(\)").unwrap();
    let mut activated = true;
    regex
        .find_iter(input)
        .filter_map(|instruction| match instruction.as_str() {
            "do()" => {
                activated = true;
                None
            }
            "don't()" => {
                activated = false;
                None
            }
            _ if activated => Some(instruction.as_str()),
            _ => None,
        })
        .map(|instruction| { // convert mut(4,3) into Vec<4,3>
            instruction
                .split(",")
                .map(|x| x.chars().filter(|c| c.is_digit(10)).collect())
                .map(|multiplicand : String| multiplicand.parse().unwrap())
                .collect()
        })
        .collect()
}
