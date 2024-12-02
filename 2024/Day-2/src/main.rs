use std::fs;
use std::path::Path;

fn main() {
    let mut input = fs::read_to_string(Path::new("./src/input.txt"))
        .expect("Should have been able to read the file");
    input = input.trim_start_matches('\u{feff}').to_string();
    let reports = parse_input(input);
    let save_report_count = reports.iter().filter(|&report| {
        is_report_safe(report)
    }).count();

    println!("Result 1: {}", save_report_count);

    let mut save_report_with_damper_count = 0;
    for report in reports {
        let mut result = is_report_safe(&report);
        if !result {
            for i in 0..report.len() {
                let mut reportclone = report.clone();
                reportclone.remove(i);

                let result_two = is_report_safe(&reportclone);
                if result_two {
                    result = true;
                    break;
                }
            }
        }
        if result {
            save_report_with_damper_count += 1;
        }
    }
    println!("Result 2: {}", save_report_with_damper_count);

}

fn parse_input(input: String) -> Vec<Vec<i32>> {
    input.lines().map(|report| {
        report.split_whitespace().map(|unit| unit.parse().unwrap()).collect()
    }).collect()
}

fn is_report_safe(report: &Vec<i32>) -> bool {
    let distance = report.windows(2).all(|pair| {
        let distance = pair[0].abs_diff(pair[1]);
        distance <= 3 && distance >= 1
    });

    let ascending = report.windows(2).all(|pair| pair[0] < pair[1]);
    let descending = report.windows(2).all(|pair| pair[0] > pair[1]);

    distance & (ascending || descending)
}