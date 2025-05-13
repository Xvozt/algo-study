package is_bad_version

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	left, right := 1, n

	for left <= right {
		mid := left + (right-left)/2

		if isBadVersion(mid) == true && isBadVersion(mid-1) == false {
			return mid
		}

		if isBadVersion(mid) == true {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}

func isBadVersion(n int) bool {
	return n >= 4
}
