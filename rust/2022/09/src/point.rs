#[derive(Clone, Copy, Debug)]
pub struct Point {
    pub x: i32,
    pub y: i32,
}

impl Point {
    pub fn touches(&self, point: &Point) -> bool {
        if self.x == point.x && self.y == point.y {
            return true;
        }
        if self.x == point.x && self.y + 1 == point.y {
            return true;
        }
        if self.x == point.x && self.y - 1 == point.y {
            return true;
        }
        if self.x + 1 == point.x && self.y == point.y {
            return true;
        }
        if self.x - 1 == point.x && self.y == point.y {
            return true;
        }
        if self.x + 1 == point.x && self.y + 1 == point.y {
            return true;
        }
        if self.x - 1 == point.x && self.y + 1 == point.y {
            return true;
        }
        if self.x + 1 == point.x && self.y - 1 == point.y {
            return true;
        }
        if self.x - 1 == point.x && self.y - 1 == point.y {
            return true;
        }
        false
    }
}
