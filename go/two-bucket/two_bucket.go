package twobucket

import "errors"

// TODO: This solution is not complete but is turning out to be a mind bender
// and hence parking it for a while.
func Solve(sizeBucketOne, sizeBucketTwo, goalAmount int, startBucket string) (goalBucket string, numSteps int, otherBucketLevel int, e error) {
	if goalAmount < 1 {
		e = errors.New("invalid goal amount")
		return
	}
	if sizeBucketOne < 1 || sizeBucketTwo < 1 {
		e = errors.New("invalid bucket size")
		return
	}
	if startBucket != "one" && startBucket != "two" {
		return "", 0, 0, errors.New("invalid start bucket")
	}
	if goalAmount > sizeBucketOne && goalAmount > sizeBucketTwo {
		e = errors.New("goal amount too big")
		return
	}
	if sizeBucketOne != goalAmount && sizeBucketTwo != goalAmount && (sizeBucketOne-sizeBucketTwo) != goalAmount && (sizeBucketTwo-sizeBucketOne) != goalAmount {
		e = errors.New("not possible")
		return
	}

	if startBucket == "one" {
		goalBucket, numSteps, otherBucketLevel, e = solve(sizeBucketOne, sizeBucketTwo, goalAmount)
	} else {
		goalBucket, numSteps, otherBucketLevel, e = solve(sizeBucketTwo, sizeBucketOne, goalAmount)
		if goalBucket == "one" {
			goalBucket = "two"
		} else {
			goalBucket = "one"
		}
	}
	return
}

func solve(s1, s2, goal int) (goalBucket string, steps int, otherLevel int, e error) {
	bucket1 := 0
	bucket2 := 0
	for {
		steps++
		if bucket1 == 0 {
			bucket1 = s1
			if goal == bucket1 {
				goalBucket = "one"
				otherLevel = bucket2
				return
			}
		} else if bucket2 == 0 {
			bucket2 = s2
			if goal == bucket2 {
				goalBucket = "two"
				otherLevel = bucket1
				return
			}
		}
		if steps > 10 {
			return
		}
	}
}
