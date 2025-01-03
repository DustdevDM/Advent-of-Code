use std::fs;
use std::path::Path;

fn main() {
    let mut input = fs::read_to_string(Path::new("./src/input.txt"))
        .expect("Should have been able to read the file");
    input = input.trim_start_matches('\u{feff}').to_string();
    let word_search = parse_input(input);

    let search_word_1: Vec<char> = "XMAS".chars().collect();
    let mut all_xmas_occurrences = 0;
    let mut all_mas_occurrences = 0;

    word_search.iter().enumerate().for_each(|(x, line)| {
        line.iter()
            .enumerate()
            .filter(|(_, &char)| char == search_word_1[0])
            .for_each(|(y, _)| {
                let right_okay = y + search_word_1.len() <= line.len();
                let left_okay = y + 1 >= search_word_1.len();
                let top_okay = x + 1 >= search_word_1.len();
                let bottom_okay = x + search_word_1.len() <= word_search.len();

                // search to right
                if right_okay {
                    if search_word_1
                        .iter()
                        .enumerate()
                        .all(|(z, &char)| word_search[x][y + z] == char)
                    {
                        all_xmas_occurrences += 1;
                    }
                }

                //search to left
                if left_okay {
                    if search_word_1
                        .iter()
                        .enumerate()
                        .all(|(z, &char)| word_search[x][y - z] == char)
                    {
                        all_xmas_occurrences += 1;
                    }
                }

                //search to bottom
                if bottom_okay {
                    if search_word_1
                        .iter()
                        .enumerate()
                        .all(|(z, &char)| word_search[x + z][y] == char)
                    {
                        all_xmas_occurrences += 1;
                    }
                }

                //search to top
                if top_okay {
                    if search_word_1
                        .iter()
                        .enumerate()
                        .all(|(z, &char)| word_search[x - z][y] == char)
                    {
                        all_xmas_occurrences += 1;
                    }
                }

                //search to top-left
                if top_okay && left_okay {
                    if search_word_1
                        .iter()
                        .enumerate()
                        .all(|(z, &char)| word_search[x - z][y - z] == char)
                    {
                        all_xmas_occurrences += 1;
                    }
                }
                //search to down-left
                if bottom_okay && left_okay {
                    if search_word_1
                        .iter()
                        .enumerate()
                        .all(|(z, &char)| word_search[x + z][y - z] == char)
                    {
                        all_xmas_occurrences += 1;
                    }
                }

                //search to top-right
                if top_okay && right_okay {
                    if search_word_1
                        .iter()
                        .enumerate()
                        .all(|(z, &char)| word_search[x - z][y + z] == char)
                    {
                        all_xmas_occurrences += 1;
                    }
                }

                //search to down-right
                if bottom_okay && right_okay {
                    if search_word_1
                        .iter()
                        .enumerate()
                        .all(|(z, &char)| word_search[x + z][y + z] == char)
                    {
                        all_xmas_occurrences += 1;
                    }
                }
            })
    });

    word_search.iter().enumerate().for_each(|(x, line)| {
        line.iter()
            .enumerate()
            .filter(|(_, &char)| char == 'A')
            .for_each(|(y, _)| {
                let right_okay = y + 1 <= line.len();
                let left_okay = y >= 1;
                let top_okay = x  >= 1;
                let bottom_okay = x + 1 <= word_search.len();

                //get x chars

                let mut xchars : Vec<char> = Vec::new();
                if top_okay && left_okay {
                    xchars.push(word_search[x - 1][y - 1])
                }
                if top_okay && right_okay {
                    xchars.push(word_search[x - 1][y + 1])
                }
                if bottom_okay && left_okay {
                    xchars.push(word_search[x + 1][y - 1])
                }
                if bottom_okay && right_okay {
                    xchars.push(word_search[x + 1][y + 1])
                }
                let mut xchars : Vec<&char> = xchars.iter().filter_map(|xchar| match xchar {
                    'M' => Some(xchar),
                    'S' => Some(xchar),
                    _ => None
                }).collect();
                if xchars.len() != 4 {
                    return;
                }

                if xchars.windows(3).all(|xchar| xchar[0] == xchar[2]) {
                    all_mas_occurrences += 1
                }


            })
    });



    println!("Result 1: {all_xmas_occurrences}");
    println!("Result 2: {all_mas_occurrences}");
}

fn parse_input(input: String) -> Vec<Vec<char>> {
    input.lines().map(|line| line.chars().collect()).collect()
}
