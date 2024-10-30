fn main() {
    let mut input: String = String::new();
    std::io::stdin().read_line(&mut input).unwrap();
    let mut list: Vec<u32> = input.split_whitespace().flat_map(str::parse::<u32>).collect::<Vec<u32>>();
    let chars: &[u8; 36] = b"0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ";
    let mut sum: Vec<char> = Vec::new();
    while list[0] > 0 {
        let remainder: usize = (list[0] % list[1]) as usize;
        sum.push(chars[remainder] as char);
        list[0] /= list[1];
    }
    let result: String = sum.iter().rev().collect();
    println!("{}", result);
}
