use crate::scores::{SCORE_DRAW, SCORE_LOST, SCORE_PAPER, SCORE_ROCK, SCORE_SCISSOR, SCORE_WON};

pub fn part_02(input: &str) -> i32 {
    let mut score = 0;
    for line in input.lines() {
        match line {
            "A X" => score += SCORE_SCISSOR + SCORE_LOST,
            "A Y" => score += SCORE_ROCK + SCORE_DRAW,
            "A Z" => score += SCORE_PAPER + SCORE_WON,
            "B X" => score += SCORE_ROCK + SCORE_LOST,
            "B Y" => score += SCORE_PAPER + SCORE_DRAW,
            "B Z" => score += SCORE_SCISSOR + SCORE_WON,
            "C X" => score += SCORE_PAPER + SCORE_LOST,
            "C Y" => score += SCORE_SCISSOR + SCORE_DRAW,
            "C Z" => score += SCORE_ROCK + SCORE_WON,
            _ => panic!("Invalid input!")
        }
    }
    score
}

