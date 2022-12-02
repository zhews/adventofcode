use crate::scores::{SCORE_DRAW, SCORE_LOST, SCORE_PAPER, SCORE_ROCK, SCORE_SCISSOR, SCORE_WON};

pub fn part_01(input: &str) -> i32 {
    let mut score = 0;
    for line in input.lines() {
        match line {
            "A X" => score += SCORE_ROCK + SCORE_DRAW,
            "A Y" => score += SCORE_PAPER + SCORE_WON,
            "A Z" => score += SCORE_SCISSOR + SCORE_LOST,
            "B X" => score += SCORE_ROCK + SCORE_LOST,
            "B Y" => score += SCORE_PAPER + SCORE_DRAW,
            "B Z" => score += SCORE_SCISSOR + SCORE_WON,
            "C X" => score += SCORE_ROCK + SCORE_WON,
            "C Y" => score += SCORE_PAPER + SCORE_LOST,
            "C Z" => score += SCORE_SCISSOR + SCORE_DRAW,
            _ => panic!("Invalid input!")
        }
    }
    score
}