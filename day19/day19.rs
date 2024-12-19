use std::collections::HashMap;
use std::fs;
use std::io::{self, Read};
use std::time::Instant;

fn main() {
    let start = Instant::now();
    let filename = "input.txt";
    match read_input_from_file(filename) {
        Ok(input) => {
            let (partials, patterns) = parse(&input);
            let mut cache: HashMap<String, i64> = HashMap::new();

            let mut res = 0;
            let mut perms = 0;

            for line in patterns {
                let temp = dp(&partials, &line, &mut cache);

                if temp > 0 {
                    res += 1;
                }

                perms += temp;
            }

            println!("Part 1\n{}", res);
            println!("Part 2 \n{}", perms);
            println!("{:?}", start.elapsed())
        }
        Err(e) => eprintln!("Error reading file: {}", e),
    }
}

fn dp(partials: &[String], line: &str, cache: &mut HashMap<String, i64>) -> i64 {
    if let Some(&result) = cache.get(line) {
        return result;
    }

    if line.is_empty() {
        return 1;
    }

    let mut res = 0;

    for k in partials {
        if line.starts_with(k) {
            res += dp(partials, &line[k.len()..], cache);
        }
    }

    cache.insert(line.to_string(), res);
    res
}

fn parse(input: &str) -> (Vec<String>, Vec<String>) {
    let parts: Vec<&str> = input.split("\n\n").collect();
    let partials: Vec<String> = parts[0].split(", ").map(|s| s.to_string()).collect();
    let patterns: Vec<String> = parts[1]
        .lines()
        .map(|line| line.trim().to_string())
        .collect();

    (partials, patterns)
}

fn read_input_from_file(filename: &str) -> io::Result<String> {
    let mut file = fs::File::open(filename)?;
    let mut content = String::new();
    file.read_to_string(&mut content)?;
    Ok(content)
}

