use std::fs;

fn main() {
    let contents = fs::read_to_string("/home/ascpial/Projets/AdventOfCode/python/day5/input.txt")
        .expect("Should have been able to read the file");

    let raw_categories: Vec<&str> = contents.split("\n\n").collect();

    let raw_seeds: Vec<i64> = raw_categories[0].split(":").collect::<Vec<&str>>()[1]
        .split(" ")
        .filter(|&seed| seed.trim() != "")
        .map(|seed| seed.parse::<i64>().expect("Failed to parse int in seeds"))
        .collect();

    let mut seeds: Vec<(i64, i64)> = Vec::new();
    for index in (0..raw_seeds.len()).step_by(2) {
        seeds.push((raw_seeds[index], raw_seeds[index+1]));
    }

    let mut categories: Vec<Vec<(i64, i64, i64)>> = Vec::new();

    for category in raw_categories.iter().skip(1) {
        let mut category_data = Vec::new();
        for line in category.split("\n").into_iter().skip(1) {
            if line == "" { continue; }
            let values: Vec<&str> = line.split(" ").collect();
            let destination: i64 = values[0].parse().expect("Failed to parse int in target maps");
            let source: i64 = values[1].parse().expect("Failed to parse int in source maps");
            let length: i64 = values[2].parse().expect("Failed to parse int in length maps");
            category_data.push((destination, length, source));
        }
        category_data.sort_by_key(|interval| interval.2);
        categories.push(category_data);
    }

    categories.reverse();

    // let test = reverse(1, &categories);
    // println!("{test}")

    let mut i = 1;
    while !value_in(reverse(i, &categories), &seeds) {
        i += 1;
    }
    println!("{i}");
}

fn value_in(value: i64, source: &Vec<(i64, i64)>) -> bool {
    for (start, length) in source {
        if start <= &value && value < start + length { return true; }
    }
    false
}

fn reverse(value: i64, categories: &Vec<Vec<(i64, i64, i64)>>) -> i64 {
    let mut state = value;
    for category in categories {
        for (destination, length, source) in category {
            if destination <= &state && state < destination + length {
                state = source + state - destination;
                break;
            }
        }
    }
    state
}
