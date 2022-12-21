pub struct Point {
    pub x: i32,
    pub y: i32,
}

impl Point {
    pub fn touches(&self, knot: &Point) -> bool {
        if self.x == knot.x && self.y == knot.y {
            return true;
        }
        if self.x == knot.x && self.y + 1 == knot.y {
            return true;
        }
        if self.x == knot.x && self.y - 1 == knot.y {
            return true;
        }
        if self.x + 1 == knot.x && self.y == knot.y {
            return true;
        }
        if self.x - 1 == knot.x && self.y == knot.y {
            return true;
        }
        if self.x + 1 == knot.x && self.y + 1 == knot.y {
            return true;
        }
        if self.x - 1 == knot.x && self.y + 1 == knot.y {
            return true;
        }
        if self.x + 1 == knot.x && self.y - 1 == knot.y {
            return true;
        }
        if self.x - 1 == knot.x && self.y - 1 == knot.y {
            return true;
        }
        false
    }
}
