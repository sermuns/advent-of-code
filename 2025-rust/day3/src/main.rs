use std::{error::Error, fs};

fn main() -> Result<(), Box<dyn Error>> {
    let input = fs::read_to_string("input")?;
    let total_joltage = input.lines().map(|line| {
        let mut joltage = String::new();

        let mut sorted_batteries: Vec<char> = line.chars().collect();
        sorted_batteries.sort();

        for c in line.chars() {
            sorted_batteries.pop
        }
    });
    Ok(())
}
