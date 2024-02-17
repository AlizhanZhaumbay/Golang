package main

func main() {
	// First Task
	slice := []int{31, 213, 53, 2, 554, 12, 15, 6, 7, 14}
	print("Before:")
	printSlice(slice)
	sortSlice(slice)
	print("After:")
	printSlice(slice)

	// Second Task
	functionsMixed := appendFunc(function1, function2, function3)
	functionsMixed(slice)
}

func incrementOdd(slice []int) {
	for index := range slice {
		if index&1 == 1 {
			slice[index] += 1
		}
	}
}

func printSlice(slice []int) {
	for _, val := range slice {
		print(val, " ")
	}
	println()
}

func reverseSlice(slice []int) {
	startIndex := 0
	endIndex := len(slice) - startIndex - 1

	for startIndex < endIndex {
		swapValues(slice, startIndex, endIndex)

		startIndex++
		endIndex--
	}
}

func sortSlice(slice []int) {
	for i := range slice[:len(slice)-1] {

		j, minIdx := i+1, i
		for j < len(slice) {
			if slice[j] < slice[minIdx] {
				minIdx = j
			}
			j++
		}

		swapValues(slice, i, minIdx)
	}
}

func swapValues(slice []int, a int, b int) {
	temp := slice[a]
	slice[a] = slice[b]
	slice[b] = temp
}

// Second Task
func appendFunc(dst func([]int), src ...func([]int)) func([]int) {
	return func(slice []int) {
		dst(slice)

		for _, function := range src {
			function(slice)
		}
	}
}

func function1(nums []int) {
	println("Function 1 invoked!")
}

func function2(nums []int) {
	println("Function 2 invoked!")
}

func function3(nums []int) {
	println("Function 3 invoked!")
}
