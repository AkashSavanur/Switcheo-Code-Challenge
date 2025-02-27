//Solution 1: Gaussian formula
func sum_to_n_b(n int) int {
    return (n*(n+1))/2
}

//Time Complexity is O(1) Constant. Involves a simple arithmetic operation.
//Space Complexity is O(1) Constant. A single value is returned as output.


//Solution 2: Iterative
func sum_to_n_a(n int) int {
    res := 0
    for i := 1; i <= n; i++ {
        res += i
    }
    return res
}

//Time Complexity is O(n) Linear. The solution involves iterating through the numbers 1 to
//n and adding them up.
// Space Complexity is O(1) Constant. The solution is stored in a single variable 'res'.


//Solution 3: Functional Programming approach
func sum_to_n_c(n int) int {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}
	return slices.Reduce(nums, func(acc, num int) int { return acc + num })
}

//Time Complexity is O(n) Linear. This is because the algorithm iterates over n elements, once for 
//initialisation of the slice and then again for reduction.
//Space Complexity is O(n) Linear: Stores number from 1 to n in a slice.

