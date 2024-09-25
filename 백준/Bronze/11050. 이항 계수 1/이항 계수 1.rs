use std::io::{self, BufRead, BufReader};

fn main() {
    let mut reader = BufReader::new(io::stdin().lock());
    let mut input: String = String::new();
    reader.read_line(&mut input).unwrap();

    let nums: Vec<usize> = input.trim().split_whitespace().flat_map(str::parse).collect();
    println!("{}", binomial_coefficient(nums[0], nums[1]));
}
const MOD: usize = 1_000_000_007;

fn mod_pow(mut a: usize, mut b: usize, m: usize) -> usize {
    let mut res = 1;
    while b > 0 {
        if b % 2 == 1 {
            res = res * a % m;
        }
        a = a * a % m;
        b /= 2;
    }
    res
}

fn binomial_coefficient(n: usize, k: usize) -> usize {
    if k > n {
        return 0;
    }

    let mut fact = vec![1; n + 1];

    for i in 2..=n {
        fact[i] = fact[i - 1] * i % MOD;
    }

    let numerator = fact[n];
    let denominator = fact[k] * fact[n - k] % MOD;

    numerator * mod_pow(denominator, MOD - 2, MOD) % MOD
}
