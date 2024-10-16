fn main() {
    let mut input: String = String::new();
    std::io::stdin().read_line(&mut input).unwrap();
    let list: Vec<&str> = input.split_whitespace().collect::<Vec<&str>>();
    let b: &str = list[0];
    let n: u32 = list[1].parse::<u32>().unwrap();
    match u32::from_str_radix(b, n) {
        Ok(value) => println!("{}", value),
        _ => {},
    }
}
