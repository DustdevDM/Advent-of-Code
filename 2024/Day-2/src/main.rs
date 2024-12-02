use std::fs;
use std::path::Path;

fn main() {
    let input = fs::read_to_string(Path::new("./src/input.txt"))
        .expect("Should have been able to read the file");
    let input = input.trim_start_matches('\u{feff}').to_string();

}

fn parse_input(input: String) {

}