use std::{
    fs::File,
    io::{self, BufRead, BufReader},
};

fn main() -> io::Result<()> {
    // Open the file and create a buffer
    let file = File::open("santa.txt").expect("Failed to open file");
    let reader = BufReader::new(file);
    let mut counter: i32 = 0;

    // Loop through lines and characters
    for line in reader.lines() {
        let line = line?;

        for c in line.chars() {
            if c == '(' {
                counter += 1;
            } else {
                counter -= 1;
            }
        }

        println!("{}", counter);
    }

    Ok(())
}
