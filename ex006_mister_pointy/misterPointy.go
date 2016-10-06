package misterPointy

func add(a, b int) int {
	return a + b
}

func addPointers(a, b *int) int {
	return *a + *b
}

func addReturnPointer(a, b int) *int{
	result := a + b
	return &result
}

func addPointersReturnPointer(a, b *int) *int {
	result := *a + *b
	return &result
}

func addModifyArgument(a int, b *int) *int {
	*b = a + *b
	return b
}

func addPointerToPointer(a, b **int) int {
	return **a + **b
}
