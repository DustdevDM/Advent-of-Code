use std::fs;
use std::path::Path;
use regex::Regex;

fn main() {
    let mut input = fs::read_to_string(Path::new("./src/input.txt"))
        .expect("Should have been able to read the file");
    input = input.trim_start_matches('\u{feff}').to_string();

    let instructions = parse_input(&input);

    let result : i32 = instructions.iter().map(|instruction| instruction[0] * instruction[1]).sum();

    println!("Result 1: {result}")
}

fn parse_input(input: &str) -> Vec<Vec<i32>> {
    let regex = Regex::new(r"mul\(\d{1,4},\d{1,4}\)").unwrap();
    regex.find_iter(input).map(|instruction| {
        instruction
            .as_str()
            .replace("mul(", "")
            .replace(")", "")
            .split(",")
            .map(|multiplicand| multiplicand.parse().unwrap())
            .collect()
    }).collect()
}