package main

import (
	"fmt"
	"reflect"
)

// Go Other Essential Concepts
// ==========================
// This file covers interfaces, methods, channels, goroutines, and more

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
	
	// 8. Type assertions and type switches
	typeAssertions()
	
	// 9. Reflection
	reflection()
	
	// 10. Error handling
	errorHandling()
}

// 1. Interfaces
// =============
func interfaces() {
	fmt.Println("\n1. INTERFACES:")
	
	// Define interfaces
	type Writer interface {
		Write([]byte) (int, error)
	}
	
	type Reader interface {
		Read([]byte) (int, error)
	}
	
	type ReadWriter interface {
		Reader
		Writer
	}
	
	// Implement interfaces
	type ConsoleWriter struct{}
	
	func (cw ConsoleWriter) Write(data []byte) (int, error) {
		fmt.Printf("   ConsoleWriter: %s\n", string(data))
		return len(data), nil
	}
	
	type FileWriter struct {
		Filename string
	}
	
	func (fw FileWriter) Write(data []byte) (int, error) {
		fmt.Printf("   FileWriter (%s): %s\n", fw.Filename, string(data))
		return len(data), nil
	}
	
	// Use interfaces
	var writer Writer = ConsoleWriter{}
	writer.Write([]byte("Hello from ConsoleWriter"))
	
	writer = FileWriter{Filename: "output.txt"}
	writer.Write([]byte("Hello from FileWriter"))
	
	// Interface composition
	type StringWriter interface {
		WriteString(string) (int, error)
	}
	
	// Empty interface
	var any interface{} = 42
	fmt.Printf("   Empty interface: %v\n", any)
	
	any = "Hello"
	fmt.Printf("   Empty interface: %v\n", any)
	
	any = []int{1, 2, 3}
	fmt.Printf("   Empty interface: %v\n", any)
}

// 2. Methods
// ===========
func methods() {
	fmt.Println("\n2. METHODS:")
	
	// Method on struct
	type Rectangle struct {
		Width  float64
		Height float64
	}
	
	// Value receiver method
	func (r Rectangle) Area() float64 {
		return r.Width * r.Height
	}
	
	// Pointer receiver method
	func (r *Rectangle) Scale(factor float64) {
		r.Width *= factor
		r.Height *= factor
	}
	
	// Method on custom type
	type Counter int
	
	func (c *Counter) Increment() {
		*c++
	}
	
	func (c Counter) Value() int {
		return int(c)
	}
	
	// Use methods
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Printf("   Rectangle area: %f\n", rect.Area())
	
	rect.Scale(2.0)
	fmt.Printf("   After scale: %+v\n", rect)
	
	counter := Counter(0)
	counter.Increment()
	counter.Increment()
	fmt.Printf("   Counter value: %d\n", counter.Value())
	
	// Method sets
	var rectPtr *Rectangle = &rect
	fmt.Printf("   Area via pointer: %f\n", rectPtr.Area()) // Go automatically dereferences
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
	
	// Channel operations
	ch3 := make(chan int, 2)
	ch3 <- 1
	ch3 <- 2
	
	// Close channel
	close(ch3)
	
	// Receive from closed channel
	val5, ok := <-ch3
	fmt.Printf("   Received: %d, ok: %t\n", val5, ok)
	
	val6, ok := <-ch3
	fmt.Printf("   Received: %d, ok: %t\n", val6, ok)
	
	val7, ok := <-ch3
	fmt.Printf("   Received: %d, ok: %t\n", val7, ok)
	
	// Select statement
	ch4 := make(chan int)
	ch5 := make(chan int)
	
	go func() {
		ch4 <- 1
	}()
	
	go func() {
		ch5 <- 2
	}()
	
	select {
	case val := <-ch4:
		fmt.Printf("   Received from ch4: %d\n", val)
	case val := <-ch5:
		fmt.Printf("   Received from ch5: %d\n", val)
	default:
		fmt.Printf("   No value ready\n")
	}
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
	
	// Multiple goroutines
	done := make(chan bool)
	
	for i := 0; i < 3; i++ {
		go func(id int) {
			fmt.Printf("   Worker %d: Processing\n", id)
			done <- true
		}(i)
	}
	
	// Wait for all goroutines
	for i := 0; i < 3; i++ {
		<-done
	}
	
	fmt.Printf("   All goroutines completed\n")
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
	
	val, exists = m1["nonexistent"]
	fmt.Printf("   m1[nonexistent] exists: %t, value: %d\n", exists, val)
	
	// Delete keys
	delete(m1, "key1")
	fmt.Printf("   After delete: %v\n", m1)
	
	// Iterate over map
	fmt.Printf("   m2 contents:\n")
	for key, value := range m2 {
		fmt.Printf("     %s: %d\n", key, value)
	}
	
	// Map of structs
	type Person struct {
		Name string
		Age  int
	}
	
	people := map[string]Person{
		"alice": {Name: "Alice", Age: 30},
		"bob":   {Name: "Bob", Age: 25},
	}
	
	fmt.Printf("   People map: %+v\n", people)
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
	
	// Copy slices
	slice4 := make([]int, len(slice3))
	copy(slice4, slice3)
	fmt.Printf("   Copied slice: %v\n", slice4)
	
	// 2D slices
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Printf("   2D slice: %v\n", matrix)
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
	
	// Function that takes function as parameter
	func calculate(a, b int, op Operation) int {
		return op(a, b)
	}
	
	fmt.Printf("   calculate(10, 4, add) = %d\n", calculate(10, 4, add))
	fmt.Printf("   calculate(10, 4, multiply) = %d\n", calculate(10, 4, multiply))
	
	// Higher-order functions
	func mapInts(slice []int, fn func(int) int) []int {
		result := make([]int, len(slice))
		for i, v := range slice {
			result[i] = fn(v)
		}
		return result
	}
	
	numbers := []int{1, 2, 3, 4, 5}
	doubled := mapInts(numbers, func(x int) int { return x * 2 })
	fmt.Printf("   Doubled: %v\n", doubled)
	
	// Closure
	func createCounter() func() int {
		count := 0
		return func() int {
			count++
			return count
		}
	}
	
	counter := createCounter()
	fmt.Printf("   Counter: %d\n", counter())
	fmt.Printf("   Counter: %d\n", counter())
	fmt.Printf("   Counter: %d\n", counter())
}

// 8. Type Assertions and Type Switches
// ======================================
func typeAssertions() {
	fmt.Println("\n8. TYPE ASSERTIONS AND TYPE SWITCHES:")
	
	// Type assertion
	var i interface{} = 42
	
	// Safe type assertion
	if val, ok := i.(int); ok {
		fmt.Printf("   i is int: %d\n", val)
	}
	
	// Type switch
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
	
	processValue(42)
	processValue("Hello")
	processValue(true)
	processValue(3.14)
	
	// Interface type assertion
	type Stringer interface {
		String() string
	}
	
	type MyInt int
	
	func (mi MyInt) String() string {
		return fmt.Sprintf("MyInt(%d)", int(mi))
	}
	
	var s Stringer = MyInt(42)
	if str, ok := s.(Stringer); ok {
		fmt.Printf("   Stringer: %s\n", str.String())
	}
}

// 9. Reflection
// ==============
func reflection() {
	fmt.Println("\n9. REFLECTION:")
	
	// Basic reflection
	var x int = 42
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	
	fmt.Printf("   Type of x: %s\n", t)
	fmt.Printf("   Value of x: %v\n", v)
	fmt.Printf("   Kind of x: %s\n", t.Kind())
	
	// Reflect on struct
	type Person struct {
		Name string
		Age  int
	}
	
	p := Person{Name: "Alice", Age: 30}
	
	// Get struct type
	structType := reflect.TypeOf(p)
	fmt.Printf("   Struct type: %s\n", structType)
	
	// Get struct value
	structValue := reflect.ValueOf(p)
	fmt.Printf("   Struct value: %v\n", structValue)
	
	// Iterate over struct fields
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		value := structValue.Field(i)
		fmt.Printf("   Field %s: %v\n", field.Name, value)
	}
	
	// Modify struct through reflection
	pPtr := &p
	structValuePtr := reflect.ValueOf(pPtr).Elem()
	ageField := structValuePtr.FieldByName("Age")
	if ageField.CanSet() {
		ageField.SetInt(35)
		fmt.Printf("   Modified age: %d\n", p.Age)
	}
}

// 10. Error Handling
// ====================
func errorHandling() {
	fmt.Println("\n10. ERROR HANDLING:")
	
	// Custom error type
	type CustomError struct {
		Code    int
		Message string
	}
	
	func (e CustomError) Error() string {
		return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
	}
	
	// Function that returns error
	func divide(a, b int) (int, error) {
		if b == 0 {
			return 0, CustomError{Code: 400, Message: "division by zero"}
		}
		return a / b, nil
	}
	
	// Handle error
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
	
	// Multiple return values
	func getValues() (int, string, bool) {
		return 42, "Hello", true
	}
	
	val1, val2, val3 := getValues()
	fmt.Printf("   Multiple values: %d, %s, %t\n", val1, val2, val3)
	
	// Ignoring return values
	_, _, _ = getValues() // Ignore all
	val1, _, _ = getValues() // Ignore second and third
}
