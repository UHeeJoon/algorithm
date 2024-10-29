use std::io::{self, BufRead, BufReader};

fn main() {
    let reader = BufReader::new(io::stdin().lock());

    let mut score_grade_sum: f32 = 0f32;
    let mut score_sum: f32 = 0f32;

    for line in reader.lines() {
        let line = line.unwrap_or_else(|_| String::new());
        let list: Vec<&str> = line.split_ascii_whitespace().collect();

        if list.len() >= 3 && list[2] != "P" {
            if let Ok(score) = list[1].parse::<f32>() {
                score_grade_sum += score * convert_grade_to_score(list[2]);
                score_sum += score;
            }
        }
    }

    println!("{:.7}", score_grade_sum /  score_sum);
}

fn convert_grade_to_score(grade: &str) -> f32 {
    match grade {
        "A+" => 4.5,
        "A0" => 4.0,
        "B+" => 3.5,
        "B0" => 3.0,
        "C+" => 2.5,
        "C0" => 2.0,
        "D+" => 1.5,
        "D0" => 1.0,
        _ => 0.0,
    }
}