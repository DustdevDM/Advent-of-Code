use std::fs;
use std::path::Path;

fn main() {
    let input = fs::read_to_string(Path::new("./src/input.txt"))
        .expect("Should have been able to read the file");
    let input = input.trim_start_matches('\u{feff}').to_string();

    let lists = parse_input(input);

    let mut left_list = lists.0;
    left_list.sort();
    let mut right_list : Vec<i32> = lists.1;
    right_list.sort();

    let mut distance_list: Vec<i32> = vec![];

    for i in 0..left_list.len() {
        if left_list[i] > right_list[i] {
            distance_list.push(left_list[i] - right_list[i]);
        }
        else {
            distance_list.push(right_list[i] - left_list[i])
        }
    }

    let mut sum : i32 = 0;
    for i in 0..distance_list.len() {
        sum += distance_list[i];
    }

    println!("Result 1: {sum}");

    let mut similarity_score : i32 = 0;

    for i in 0..left_list.len() {
        let mut found = 0;
        for y in 0..right_list.len() {
            if left_list[i] == right_list[y] {
                found+= 1;
            }
        }
        similarity_score += left_list[i] * found;
    }

    println!("Result 2: {similarity_score}");

}

fn parse_input(input: String) -> (Vec<i32>, Vec<i32>) {
    let mut first_list : Vec<i32> = vec![];
    let mut second_list: Vec<i32> = vec![];

    for line in input.lines(){
        let line_split = line.split_whitespace().collect::<Vec<&str>>();
        first_list.push(line_split[0].trim().parse().unwrap());
        second_list.push(line_split[1].trim().parse().unwrap());
    }

    return (first_list, second_list);
}