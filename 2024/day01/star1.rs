use std::fs;

fn main() {
    println!("Hello, world!");
    let contents = fs::read_to_string("input.txt")
      .expect("If this fails, think about downloading the input");

    for line in contents.split("\n") {
        println!("line: {line}")
        let numbers = line.split();
    }
}
