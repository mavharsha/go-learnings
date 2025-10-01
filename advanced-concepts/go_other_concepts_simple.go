package main

import (
	"fmt"
)

// Go Other Essential Concepts - Simple Guide
// ==========================================

func main() {
	fmt.Println("=== Go Other Essential Concepts ===")
	
	// 1. Interfaces
	interfaces()
	
	// 2. Methods
	methods()
	
	// 3. Channels
	channels()
	
	// 4. Goroutines
	goroutines()
	
	// 5. Maps
	maps()
	
	// 6. Slices
	slices()
	
	// 7. Functions as values
	functionsAsValues()
	
	// 8. Type assertions
	typeAssertions()
	
	// 9. Error handling
	errorHandling()
}

// 1. Interfaces
// =============
func interfaces() {
	fmt.Println("\n1. INTERFACES:")
	
	// Define interface
	type Writer interface {
		Write([]byte) (int, error)
	}
	
	// Implement interface
	var writer Writer = ConsoleWriter{}
	writer.Write([]byte("Hello from interface!"))
	
	// Empty interface
	var any interface{} = 42
	fmt.Printf("   Empty interface: %v\n", any)
}

// 2. Methods
// ===========
func methods() {
	fmt.Println("\n2. METHODS:")
	
	// Create rectangle and use methods
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Printf("   Rectangle area: %f\n", rect.Area())
	
	rect.Scale(2.0)
	fmt.Printf("   After scale: %+v\n", rect)
}

// 3. Channels
// ============
func channels() {
	fmt.Println("\n3. CHANNELS:")
	
	// Unbuffered channel
	ch1 := make(chan int)
	
	// Buffered channel
	ch2 := make(chan string, 3)
	
	// Send and receive
	go func() {
		ch1 <- 42
		ch2 <- "Hello"
		ch2 <- "World"
		ch2 <- "Go"
	}()
	
	// Receive values
	val1 := <-ch1
	fmt.Printf("   Received from ch1: %d\n", val1)
	
	val2 := <-ch2
	val3 := <-ch2
	val4 := <-ch2
	fmt.Printf("   Received from ch2: %s, %s, %s\n", val2, val3, val4)
}

// 4. Goroutines
// ==============
func goroutines() {
	fmt.Println("\n4. GOROUTINES:")
	
	// Basic goroutine
	go func() {
		fmt.Printf("   Goroutine 1: Hello from goroutine!\n")
	}()
	
	// Goroutine with parameters
	go func(id int) {
		fmt.Printf("   Goroutine %d: Running\n", id)
	}(1)
	
	// Goroutine with return value
	resultCh := make(chan int)
	go func() {
		resultCh <- 42
	}()
	
	result := <-resultCh
	fmt.Printf("   Goroutine result: %d\n", result)
}

// 5. Maps
// ========
func maps() {
	fmt.Println("\n5. MAPS:")
	
	// Create maps
	m1 := make(map[string]int)
	m2 := map[string]int{
		"apple":  5,
		"banana": 3,
		"orange": 8,
	}
	
	// Add/update values
	m1["key1"] = 42
	m1["key2"] = 100
	
	// Access values
	fmt.Printf("   m1[key1]: %d\n", m1["key1"])
	fmt.Printf("   m1[key2]: %d\n", m1["key2"])
	
	// Check if key exists
	val, exists := m1["key1"]
	fmt.Printf("   m1[key1] exists: %t, value: %d\n", exists, val)
	
	// Iterate over map
	fmt.Printf("   m2 contents:\n")
	for key, value := range m2 {
		fmt.Printf("     %s: %d\n", key, value)
	}
}

// 6. Slices
// ==========
func slices() {
	fmt.Println("\n6. SLICES:")
	
	// Create slices
	slice1 := make([]int, 5)        // length 5, capacity 5
	slice2 := make([]int, 3, 10)    // length 3, capacity 10
	slice3 := []int{1, 2, 3, 4, 5} // slice literal
	
	fmt.Printf("   slice1: %v (len: %d, cap: %d)\n", slice1, len(slice1), cap(slice1))
	fmt.Printf("   slice2: %v (len: %d, cap: %d)\n", slice2, len(slice2), cap(slice2))
	fmt.Printf("   slice3: %v (len: %d, cap: %d)\n", slice3, len(slice3), cap(slice3))
	
	// Append to slice
	slice3 = append(slice3, 6, 7, 8)
	fmt.Printf("   After append: %v (len: %d, cap: %d)\n", slice3, len(slice3), cap(slice3))
	
	// Slice operations
	fmt.Printf("   slice3[2:5]: %v\n", slice3[2:5])
	fmt.Printf("   slice3[:3]: %v\n", slice3[:3])
	fmt.Printf("   slice3[3:]: %v\n", slice3[3:])
}

// 7. Functions as Values
// =======================
func functionsAsValues() {
	fmt.Println("\n7. FUNCTIONS AS VALUES:")
	
	// Function type
	type Operation func(int, int) int
	
	// Function implementations
	add := func(a, b int) int {
		return a + b
	}
	
	multiply := func(a, b int) int {
		return a * b
	}
	
	// Use function values
	var op Operation = add
	fmt.Printf("   add(5, 3) = %d\n", op(5, 3))
	
	op = multiply
	fmt.Printf("   multiply(5, 3) = %d\n", op(5, 3))
	
	// Higher-order functions
	numbers := []int{1, 2, 3, 4, 5}
	doubled := mapInts(numbers, func(x int) int { return x * 2 })
	fmt.Printf("   Doubled: %v\n", doubled)
}

// 8. Type Assertions
// ===================
func typeAssertions() {
	fmt.Println("\n8. TYPE ASSERTIONS:")
	
	// Type assertion
	var i interface{} = 42
	
	// Safe type assertion
	if val, ok := i.(int); ok {
		fmt.Printf("   i is int: %d\n", val)
	}
	
	// Type switch
	processValue(42)
	processValue("Hello")
	processValue(true)
	processValue(3.14)
}

// 9. Error Handling
// ===================
func errorHandling() {
	fmt.Println("\n9. ERROR HANDLING:")
	
	// Function that returns error
	result, err := divide(10, 2)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		fmt.Printf("   10 / 2 = %d\n", result)
	}
	
	result, err = divide(10, 0)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		fmt.Printf("   10 / 0 = %d\n", result)
	}
}

// Helper functions
// ================
func mapInts(slice []int, fn func(int) int) []int {
	result := make([]int, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

func processValue(v interface{}) {
	switch val := v.(type) {
	case int:
		fmt.Printf("   Integer: %d\n", val)
	case string:
		fmt.Printf("   String: %s\n", val)
	case bool:
		fmt.Printf("   Boolean: %t\n", val)
	default:
		fmt.Printf("   Unknown type: %T\n", val)
	}
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// Types
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	fmt.Printf("   ConsoleWriter: %s\n", string(data))
	return len(data), nil
}
