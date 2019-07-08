package exercises

func countPathsRecursive(steps int) int {
	if steps < 0 {
		return 0
	} else if steps == 0 {
		return 1
	}

	return countPathsRecursive(steps-1) + countPathsRecursive(steps-2) + countPathsRecursive(steps-3)
}

func countPathsRecursiveMemo(steps int) int {
	memo := make([]int, steps+1)
	return countPathsRecursiveM(steps, memo)
}

func countPathsRecursiveM(steps int, memo []int) int {
	if steps < 0 {
		return 0
	} else if steps == 0 {
		return 1
	}

	if memo[steps] == 0 {
		memo[steps] = countPathsRecursiveM(steps-1, memo) + countPathsRecursiveM(steps-2, memo) + countPathsRecursiveM(steps-3, memo)
	}
	return memo[steps]
}

func countPathsDP(steps int) int {
	if steps < 0 {
		return 0
	} else if steps <= 1 {
		return 1
	}
	paths := make([]int, steps+1)

	paths[0] = 1
	paths[1] = 1
	paths[2] = 2

	for i := 3; i <= steps; i++ {
		paths[i] = paths[i-1] + paths[i-2] + paths[i-3]
	}
	return paths[steps]
}

func countPathsIterative(steps int) int {
	if steps < 0 {
		return 0
	} else if steps <= 1 {
		return 1
	}
	paths := [...]int{1, 1, 2}

	for i := 3; i <= steps; i++ {
		count := paths[0] + paths[1] + paths[2]
		paths[0] = paths[1]
		paths[1] = paths[2]
		paths[2] = count
	}
	return paths[2]
}
