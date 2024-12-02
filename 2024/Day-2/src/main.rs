use std::fs;
use std::path::Path;

fn main() {
    let mut input = fs::read_to_string(Path::new("./src/input.txt"))
        .expect("Should have been able to read the file");
    input = input.trim_start_matches('\u{feff}').to_string();
    let reports = parse_input(input);
    let save_report_count = reports.iter().filter(|&report| is_report_safe(report)).count();

    println!("Result 1: {}", save_report_count)
}

fn parse_input(input: String) -> Vec<Vec<i32>> {
    let mut reports = Vec::new();
    for level in input.split("\n") {
        println!("{level}");
        let mut parsed_level = Vec::new();
        for unit in level.split_whitespace() {
            parsed_level.push(unit.parse().unwrap())
        }
        reports.push(parsed_level);
    }
    reports
}

fn is_report_safe(report: &Vec<i32>) -> bool {
    let distance = report.windows(2).all(|pair| {
        let distance = pair[0].abs_diff(pair[1]);
        distance <= 3 && distance >= 1
    });

    let ascending = report.windows(2).all(|pair| pair[0] < pair[1]);
    let descending = report.windows(2).all(|pair| pair[0] > pair[1]);

    println!("{0}\nDistance: {distance}\nAscending: {ascending}\nDescending: {descending}\n---", format!("{:?}", report));
    distance & (ascending || descending)
}