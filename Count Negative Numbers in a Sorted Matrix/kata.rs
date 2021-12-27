package kata

impl Solution {
    pub fn count_negatives(grid: Vec<Vec<i32>>) -> i32 {
        let mut counter: i32 = 0;
        let bound: i32 = 0; 
        for vector in grid.iter() {
                for num in vector.iter() {
                    if num < &bound {
                        counter += 1
                    }
                }
            }
        return counter
    }
}
