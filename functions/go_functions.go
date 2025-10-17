package main

import (
	"fmt"
	"strings"
)

// Go Functions - Complete Guide
// ==============================
// This file demonstrates Go functions with comprehensive examples

// Global function examples
// =========================

// Simple function with no parameters or return value
func greet() {
	fmt.Println("Hello, World!")
}

// Function with parameters
func greetPerson(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// Function with multiple parameters
func add(a int, b int) int {
	return a + b
}

// Function with parameters of same type (shorthand)
func multiply(a, b int) int {
	return a * b
}

// Function with multiple return values
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// Function with named return values
func getCoordinates() (x, y int) {
	x = 10
	y = 20
	return // Naked return
}

// Function with named return values (explicit return)
func getRectangleDimensions() (width, height int) {
	width = 100
	height = 50
	return width, height
}

func main() {
	fmt.Println("=== Go Functions ===")
	
	// 1. Basic function declaration and calling
	basicFunctions()
	
	// 2. Multiple parameters and return values
	multipleReturns()
	
	// 3. Named return values
	namedReturns()
	
	// 4. Variadic functions
	variadicFunctions()
	
	// 5. Functions as values
	functionsAsValues()
	
	// 6. Anonymous functions
	anonymousFunctions()
	
	// 7. Closures
	closures()
	
	// 8. Recursion
	recursion()
	
	// 9. Defer statements
	deferStatements()
	
	// 10. Higher-order functions
	higherOrderFunctions()
}

// 1. Basic Function Declaration and Calling
// ==========================================
func basicFunctions() {
	fmt.Println("\n1. BASIC FUNCTIONS:")
	
	// Call function with no parameters
	greet()
	
	// Call function with parameters
	greetPerson("Alice")
	greetPerson("Bob")
	
	// Call function with return value
	sum := add(5, 3)
	fmt.Printf("   5 + 3 = %d\n", sum)
	
	// Call function with shorthand parameters
	product := multiply(4, 7)
	fmt.Printf("   4 * 7 = %d\n", product)
}

// 2. Multiple Parameters and Return Values
// =========================================
func multipleReturns() {
	fmt.Println("\n2. MULTIPLE PARAMETERS AND RETURN VALUES:")
	
	// Function with multiple return values
	result, err := divide(10.0, 2.0)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		fmt.Printf("   10.0 / 2.0 = %f\n", result)
	}
	
	// Handle error case
	result2, err2 := divide(10.0, 0.0)
	if err2 != nil {
		fmt.Printf("   Error: %v\n", err2)
	} else {
		fmt.Printf("   10.0 / 0.0 = %f\n", result2)
	}
	
	// Ignore return value with blank identifier
	result3, _ := divide(20.0, 4.0)
	fmt.Printf("   20.0 / 4.0 = %f (error ignored)\n", result3)
	
	// Multiple return values in action
	min, max := getMinMax(5, 10)
	fmt.Printf("   Min: %d, Max: %d\n", min, max)
}

// 3. Named Return Values
// ======================
func namedReturns() {
	fmt.Println("\n3. NAMED RETURN VALUES:")
	
	// Call function with named returns
	x, y := getCoordinates()
	fmt.Printf("   Coordinates: x=%d, y=%d\n", x, y)
	
	// Call function with named returns (explicit)
	width, height := getRectangleDimensions()
	fmt.Printf("   Rectangle: width=%d, height=%d\n", width, height)
	
	// Named returns make code more readable
	area, perimeter := calculateRectangle(10, 5)
	fmt.Printf("   Rectangle: area=%d, perimeter=%d\n", area, perimeter)
}

// 4. Variadic Functions
// =====================
func variadicFunctions() {
	fmt.Println("\n4. VARIADIC FUNCTIONS:")
	
	// Variadic function - accepts variable number of arguments
	sum1 := sum(1, 2, 3)
	fmt.Printf("   sum(1, 2, 3) = %d\n", sum1)
	
	sum2 := sum(1, 2, 3, 4, 5)
	fmt.Printf("   sum(1, 2, 3, 4, 5) = %d\n", sum2)
	
	// Pass slice to variadic function using ...
	numbers := []int{10, 20, 30, 40}
	sum3 := sum(numbers...)
	fmt.Printf("   sum(10, 20, 30, 40) = %d\n", sum3)
	
	// Variadic with other parameters
	message := joinStrings(" - ", "Go", "is", "awesome")
	fmt.Printf("   Joined: %s\n", message)
}

// 5. Functions as Values
// ======================
func functionsAsValues() {
	fmt.Println("\n5. FUNCTIONS AS VALUES:")
	
	// Assign function to variable
	addFunc := add
	result := addFunc(10, 20)
	fmt.Printf("   addFunc(10, 20) = %d\n", result)
	
	// Function as parameter
	applyOperation(5, 3, add)
	applyOperation(5, 3, multiply)
	
	// Function as return value
	adder := makeAdder(10)
	fmt.Printf("   adder(5) = %d\n", adder(5))
	fmt.Printf("   adder(15) = %d\n", adder(15))
}

// 6. Anonymous Functions
// ======================
func anonymousFunctions() {
	fmt.Println("\n6. ANONYMOUS FUNCTIONS:")
	
	// Anonymous function - defined inline
	func() {
		fmt.Println("   Anonymous function called")
	}()
	
	// Anonymous function with parameters
	func(name string) {
		fmt.Printf("   Hello, %s!\n", name)
	}("Charlie")
	
	// Anonymous function assigned to variable
	square := func(x int) int {
		return x * x
	}
	fmt.Printf("   square(5) = %d\n", square(5))
	
	// Anonymous function with return value
	result := func(a, b int) int {
		return a + b
	}(10, 20)
	fmt.Printf("   Anonymous sum: %d\n", result)
}

// 7. Closures
// ===========
func closures() {
	fmt.Println("\n7. CLOSURES:")
	
	// Closure - function that references variables from outer scope
	counter := makeCounter()
	fmt.Printf("   Counter: %d\n", counter())
	fmt.Printf("   Counter: %d\n", counter())
	fmt.Printf("   Counter: %d\n", counter())
	
	// Each closure has its own scope
	counter2 := makeCounter()
	fmt.Printf("   Counter2: %d\n", counter2())
	fmt.Printf("   Counter2: %d\n", counter2())
	
	// Closure with parameters
	multiplier := makeMultiplier(5)
	fmt.Printf("   Multiplier(3) = %d\n", multiplier(3))
	fmt.Printf("   Multiplier(7) = %d\n", multiplier(7))
	
	// Closure capturing loop variable
	functions := makeFunctions()
	for i, f := range functions {
		fmt.Printf("   Function %d: %d\n", i, f())
	}
}

// 8. Recursion
// ============
func recursion() {
	fmt.Println("\n8. RECURSION:")
	
	// Factorial using recursion
	fmt.Printf("   Factorial(5) = %d\n", factorial(5))
	fmt.Printf("   Factorial(7) = %d\n", factorial(7))
	
	// Fibonacci using recursion
	fmt.Printf("   Fibonacci(10) = %d\n", fibonacci(10))
	fmt.Printf("   Fibonacci(15) = %d\n", fibonacci(15))
	
	// Sum of array using recursion
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("   Sum of %v = %d\n", numbers, sumArray(numbers))
}

// 9. Defer Statements
// ===================
func deferStatements() {
	fmt.Println("\n9. DEFER STATEMENTS:")
	
	// Defer executes at end of function
	fmt.Println("   Start")
	defer fmt.Println("   Deferred 1")
	defer fmt.Println("   Deferred 2")
	defer fmt.Println("   Deferred 3")
	fmt.Println("   End")
	// Note: Deferred functions execute in LIFO order (Last In, First Out)
	
	// Defer with function calls
	deferExample()
	
	// Defer with parameters
	deferWithParameters()
}

// 10. Higher-Order Functions
// ===========================
func higherOrderFunctions() {
	fmt.Println("\n10. HIGHER-ORDER FUNCTIONS:")
	
	// Map function
	numbers := []int{1, 2, 3, 4, 5}
	squared := mapInts(numbers, func(x int) int {
		return x * x
	})
	fmt.Printf("   Squared: %v\n", squared)
	
	// Filter function
	evenNumbers := filterInts(numbers, func(x int) bool {
		return x%2 == 0
	})
	fmt.Printf("   Even numbers: %v\n", evenNumbers)
	
	// Reduce function
	sum := reduceInts(numbers, 0, func(acc, x int) int {
		return acc + x
	})
	fmt.Printf("   Sum: %d\n", sum)
	
	// Chain operations
	result := reduceInts(
		filterInts(
			mapInts(numbers, func(x int) int { return x * 2 }),
			func(x int) bool { return x > 5 },
		),
		0,
		func(acc, x int) int { return acc + x },
	)
	fmt.Printf("   Chain result: %d\n", result)
}

// Helper functions
// ================

func getMinMax(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

func calculateRectangle(width, height int) (area, perimeter int) {
	area = width * height
	perimeter = 2 * (width + height)
	return
}

func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func joinStrings(separator string, parts ...string) string {
	return strings.Join(parts, separator)
}

func applyOperation(a, b int, operation func(int, int) int) {
	result := operation(a, b)
	fmt.Printf("   Operation result: %d\n", result)
}

func makeAdder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func makeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func makeFunctions() []func() int {
	var functions []func() int
	for i := 0; i < 5; i++ {
		// Capture loop variable properly
		i := i // Create new variable in each iteration
		functions = append(functions, func() int {
			return i * i
		})
	}
	return functions
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func sumArray(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	return numbers[0] + sumArray(numbers[1:])
}

func deferExample() {
	fmt.Println("   Defer example start")
	defer fmt.Println("   Defer example end")
	fmt.Println("   Defer example middle")
}

func deferWithParameters() {
	x := 10
	defer fmt.Printf("   Deferred value: %d\n", x)
	x = 20
	fmt.Printf("   Current value: %d\n", x)
	// Deferred function uses x=10 (value when defer was called)
}

func mapInts(numbers []int, fn func(int) int) []int {
	result := make([]int, len(numbers))
	for i, num := range numbers {
		result[i] = fn(num)
	}
	return result
}

func filterInts(numbers []int, fn func(int) bool) []int {
	var result []int
	for _, num := range numbers {
		if fn(num) {
			result = append(result, num)
		}
	}
	return result
}

func reduceInts(numbers []int, initial int, fn func(int, int) int) int {
	result := initial
	for _, num := range numbers {
		result = fn(result, num)
	}
	return result
}

