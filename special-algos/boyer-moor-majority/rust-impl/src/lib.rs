pub fn majority_element(nums: Vec<i32>) -> i32 {
    if nums.len() == 0 {
        return 0;
    }

    let mut candidate = nums[0];
    let mut count = 1;

    for num in nums.iter().skip(1) {
        if count == 0 {
            candidate = *num;
        }
        if candidate == *num {
            count += 1;
        } else {
            count -= 1;
        }
    }

    candidate
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let result = majority_element(vec![2, 2, 1, 1, 1, 2, 2]);
        assert_eq!(result, 2);
    }

    #[test]
    fn it_works_with_empty_vec() {
        let result = majority_element(vec![]);
        assert_eq!(result, 0);
    }

    #[test]
    fn it_works_with_single_element() {
        let result = majority_element(vec![5]);
        assert_eq!(result, 5);
    }

    #[test]
    fn it_works_with_3_elements() {
        let result = majority_element(vec![3, 2, 3]);
        assert_eq!(result, 3);
    }
}
