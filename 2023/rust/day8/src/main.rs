use regex::Regex;
use std::collections::HashMap;
use std::fs;

fn main() {
    let re = Regex::new(r"([A-Z1-9]{3}) = \(([A-Z1-9]{3}), ([A-Z1-9]{3})\)").unwrap();
    let contents = fs::read_to_string("/home/ascpial/Projets/AdventOfCode/python/day8/input.txt")
        .expect("Should have been able to read the file");

    let contents: Vec<&str> = contents.split("\n").collect();

    let instructions: Vec<i64> = contents[0]
        .chars()
        .map(|direction| if direction == 'L' { 0 } else { 1 })
        .collect();

    let nodes: HashMap<_, (_, _)> = contents
        .into_iter()
        .skip(2)
        .filter(|node| node != &"")
        .map(|node| {
            let Some((_, [node, left, right])) = re.captures(node).map(|caps| caps.extract())
            else {
                println!("{node}");
                panic!("C'est pété");
            };
            (node, (left, right))
        })
        .collect();

    println!("{:?}", nodes);

    let mut positions: Vec<_> = nodes
        .keys()
        .filter(|node| node.chars().last().unwrap() == 'A')
        .collect();

    let mut steps = 0;
    let mut instruction_index = 0;

    while !positions_valids(&positions) {
        steps += 1;
        let direction = instructions.get(instruction_index).unwrap();
        positions = positions
            .into_iter()
            .map(|node| {
                if direction == &0 {
                    &nodes.get(node).unwrap().0
                } else {
                    &nodes.get(node).unwrap().1
                }
            })
            .collect();
        instruction_index += 1;
        if instruction_index >= instructions.len() {
            instruction_index = 0;
        }
    }
    println!("{:?}", positions);
    println!("{steps}");
}

fn positions_valids(positions: &Vec<&&str>) -> bool {
    for position in positions {
        if position.chars().last().unwrap() != 'Z' {
            return false;
        }
    }
    return true;
}
