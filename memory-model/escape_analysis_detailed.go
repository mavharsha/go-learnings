package main

import (
	"fmt"
	"runtime"
	"time"
	"unsafe"
)

// Detailed Escape Analysis Examples
// =================================
// This file shows specific scenarios where Go's escape analysis
// determines stack vs heap allocation with detailed explanations.

func main() {
	fmt.Println("=== Detailed Escape Analysis Examples ===")
	
	// Scenario 1: Basic variable allocation
	basicVariableAllocation()
	
	// Scenario 2: Function return patterns
	functionReturnPatterns()
	
	// Scenario 3: Struct field access patterns
	structFieldPatterns()
	
	// Scenario 4: Interface and method dispatch
	interfaceMethodPatterns()
	
	// Scenario 5: Slice and array patterns
	sliceArrayPatterns()
	
	// Scenario 6: Closure capture patterns
	closureCapturePatterns()
	
	// Scenario 7: Goroutine and concurrency patterns
	goroutinePatterns()
	
	// Scenario 8: Large object allocation patterns
	largeObjectPatterns()
	
	// Scenario 9: Memory alignment and padding
	memoryAlignmentPatterns()
	
	// Scenario 10: Performance implications
	performanceImplications()
}

// Scenario 1: Basic Variable Allocation
// ====================================
func basicVariableAllocation() {
	fmt.Println("\n1. BASIC VARIABLE ALLOCATION:")
	
	// STACK: Simple local variables
	var a int = 42
	var b float64 = 3.14
	var c string = "Hello"
	var d bool = true
	
	fmt.Printf("   Local variables: a=%d, b=%f, c=%s, d=%t\n", a, b, c, d)
	fmt.Println("   ✓ STACK: Simple local variables")
	
	// STACK: Arrays with known size
	var arr [10]int
	for i := range arr {
		arr[i] = i * 2
	}
	fmt.Printf("   Array: %v\n", arr)
	fmt.Println("   ✓ STACK: Small arrays with known size")
	
	// STACK: Struct with simple fields
	type Point struct {
		X, Y int
	}
	var p Point = Point{X: 10, Y: 20}
	fmt.Printf("   Struct: %+v\n", p)
	fmt.Println("   ✓ STACK: Simple structs")
	
	// HEAP: Taking address of local variable
	ptr := &a
	fmt.Printf("   Address of local: %p\n", ptr)
	fmt.Println("   ✗ HEAP: Taking address of local variable")
}

// Scenario 2: Function Return Patterns
// ===================================
func functionReturnPatterns() {
	fmt.Println("\n2. FUNCTION RETURN PATTERNS:")
	
	// STACK: Return by value
	result1 := returnValue()
	fmt.Printf("   Return value: %d\n", result1)
	fmt.Println("   ✓ STACK: Return by value")
	
	// HEAP: Return by pointer
	result2 := returnPointer()
	fmt.Printf("   Return pointer: %d\n", *result2)
	fmt.Println("   ✗ HEAP: Return by pointer")
	
	// STACK: Return struct by value
	point := returnStruct()
	fmt.Printf("   Return struct: %+v\n", point)
	fmt.Println("   ✓ STACK: Return struct by value")
	
	// HEAP: Return struct by pointer
	pointPtr := returnStructPointer()
	fmt.Printf("   Return struct pointer: %+v\n", *pointPtr)
	fmt.Println("   ✗ HEAP: Return struct by pointer")
}

// Scenario 3: Struct Field Access Patterns
// ========================================
func structFieldPatterns() {
	fmt.Println("\n3. STRUCT FIELD ACCESS PATTERNS:")
	
	// STACK: Struct on stack, field access
	person := Person{Name: "Alice", Age: 30}
	person.Age++
	fmt.Printf("   Stack struct field: %+v\n", person)
	fmt.Println("   ✓ STACK: Struct on stack, field access")
	
	// HEAP: Struct on heap, field access
	personPtr := &Person{Name: "Bob", Age: 25}
	personPtr.Age++
	fmt.Printf("   Heap struct field: %+v\n", *personPtr)
	fmt.Println("   ✗ HEAP: Struct on heap, field access")
	
	// HEAP: Taking address of struct field
	agePtr := &person.Age
	*agePtr = 35
	fmt.Printf("   Address of field: %d\n", *agePtr)
	fmt.Println("   ✗ HEAP: Taking address of struct field")
}

// Scenario 4: Interface and Method Dispatch
// ========================================
func interfaceMethodPatterns() {
	fmt.Println("\n4. INTERFACE AND METHOD DISPATCH:")
	
	// HEAP: Interface variable
	var writer Writer = &ConsoleWriter{}
	writer.Write("Hello")
	fmt.Println("   ✗ HEAP: Interface variables")
	
	// HEAP: Method calls on interfaces
	var reader Reader = &StringReader{data: "Hello World"}
	content := reader.Read()
	fmt.Printf("   Interface method: %s\n", content)
	fmt.Println("   ✗ HEAP: Method calls on interfaces")
	
	// HEAP: Empty interface
	var any interface{} = 42
	fmt.Printf("   Empty interface: %v\n", any)
	fmt.Println("   ✗ HEAP: Empty interface")
	
	// HEAP: Type assertion
	if val, ok := any.(int); ok {
		fmt.Printf("   Type assertion: %d\n", val)
		fmt.Println("   ✗ HEAP: Type assertion")
	}
}

// Scenario 5: Slice and Array Patterns
// ===================================
func sliceArrayPatterns() {
	fmt.Println("\n5. SLICE AND ARRAY PATTERNS:")
	
	// STACK: Small array
	var arr [5]int
	for i := range arr {
		arr[i] = i
	}
	fmt.Printf("   Small array: %v\n", arr)
	fmt.Println("   ✓ STACK: Small arrays")
	
	// HEAP: Large array (might escape)
	var largeArr [1000]int
	fmt.Printf("   Large array: %d elements\n", len(largeArr))
	fmt.Println("   ? HEAP: Large arrays might escape")
	
	// HEAP: Slice
	slice := make([]int, 5)
	for i := range slice {
		slice[i] = i * 2
	}
	fmt.Printf("   Slice: %v\n", slice)
	fmt.Println("   ✗ HEAP: Slices always on heap")
	
	// HEAP: Slice literal
	slice2 := []int{1, 2, 3, 4, 5}
	fmt.Printf("   Slice literal: %v\n", slice2)
	fmt.Println("   ✗ HEAP: Slice literals")
	
	// HEAP: Slice of structs
	people := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
	}
	fmt.Printf("   Slice of structs: %+v\n", people)
	fmt.Println("   ✗ HEAP: Slice of structs")
}

// Scenario 6: Closure Capture Patterns
// ===================================
func closureCapturePatterns() {
	fmt.Println("\n6. CLOSURE CAPTURE PATTERNS:")
	
	// HEAP: Closure capturing local variable
	counter := createCounter()
	fmt.Printf("   Counter: %d\n", counter())
	fmt.Println("   ✗ HEAP: Closure capturing local variable")
	
	// HEAP: Closure capturing parameter
	multiplier := createMultiplier(5)
	fmt.Printf("   Multiplier: %d\n", multiplier(3))
	fmt.Println("   ✗ HEAP: Closure capturing parameter")
	
	// HEAP: Closure capturing struct field
	person := Person{Name: "Alice", Age: 30}
	ageGetter := func() int {
		return person.Age  // Captures person
	}
	fmt.Printf("   Age getter: %d\n", ageGetter())
	fmt.Println("   ✗ HEAP: Closure capturing struct field")
	
	// HEAP: Closure capturing slice
	numbers := []int{1, 2, 3, 4, 5}
	summer := func() int {
		sum := 0
		for _, n := range numbers {
			sum += n
		}
		return sum
	}
	fmt.Printf("   Summer: %d\n", summer())
	fmt.Println("   ✗ HEAP: Closure capturing slice")
}

// Scenario 7: Goroutine and Concurrency Patterns
// =============================================
func goroutinePatterns() {
	fmt.Println("\n7. GOROUTINE AND CONCURRENCY PATTERNS:")
	
	// HEAP: Goroutine capturing local variable
	x := 42
	go func() {
		fmt.Printf("   Goroutine with captured variable: %d\n", x)
	}()
	fmt.Println("   ✗ HEAP: Goroutine capturing local variable")
	
	// HEAP: Goroutine with closure
	counter := 0
	go func() {
		for i := 0; i < 5; i++ {
			counter++
			fmt.Printf("   Goroutine counter: %d\n", counter)
		}
	}()
	fmt.Println("   ✗ HEAP: Goroutine with closure")
	
	// HEAP: Goroutine with channel
	ch := make(chan int, 1)
	go func() {
		ch <- 42
	}()
	value := <-ch
	fmt.Printf("   Goroutine with channel: %d\n", value)
	fmt.Println("   ✗ HEAP: Goroutine with channel")
}

// Scenario 8: Large Object Allocation Patterns
// ============================================
func largeObjectPatterns() {
	fmt.Println("\n8. LARGE OBJECT ALLOCATION PATTERNS:")
	
	// HEAP: Large struct
	type LargeStruct struct {
		Data [1000]int
	}
	var large LargeStruct
	fmt.Printf("   Large struct size: %d bytes\n", unsafe.Sizeof(large))
	fmt.Println("   ✗ HEAP: Large structs")
	
	// HEAP: Large slice
	largeSlice := make([]int, 10000)
	fmt.Printf("   Large slice: %d elements\n", len(largeSlice))
	fmt.Println("   ✗ HEAP: Large slices")
	
	// HEAP: Dynamic allocation
	size := 5000
	dynamicSlice := make([]int, size)
	fmt.Printf("   Dynamic slice: %d elements\n", len(dynamicSlice))
	fmt.Println("   ✗ HEAP: Dynamic allocation")
}

// Scenario 9: Memory Alignment and Padding
// ========================================
func memoryAlignmentPatterns() {
	fmt.Println("\n9. MEMORY ALIGNMENT AND PADDING:")
	
	// Show struct alignment
	type AlignedStruct struct {
		Field1 int32  // 4 bytes
		Field2 int64  // 8 bytes
		Field3 int32  // 4 bytes
	}
	
	var s AlignedStruct
	fmt.Printf("   Struct size: %d bytes\n", unsafe.Sizeof(s))
	fmt.Printf("   Field1 offset: %d\n", unsafe.Offsetof(s.Field1))
	fmt.Printf("   Field2 offset: %d\n", unsafe.Offsetof(s.Field2))
	fmt.Printf("   Field3 offset: %d\n", unsafe.Offsetof(s.Field3))
	fmt.Println("   ✓ Go automatically handles alignment")
}

// Scenario 10: Performance Implications
// ====================================
func performanceImplications() {
	fmt.Println("\n10. PERFORMANCE IMPLICATIONS:")
	
	// Stack allocation performance
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		var x int = i
		_ = x
	}
	stackTime := time.Since(start)
	fmt.Printf("   Stack allocation: %v\n", stackTime)
	
	// Heap allocation performance
	start = time.Now()
	for i := 0; i < 1000000; i++ {
		x := new(int)
		*x = i
		_ = x
	}
	heapTime := time.Since(start)
	fmt.Printf("   Heap allocation: %v\n", heapTime)
	
	// Show memory stats
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("   Heap size: %d KB\n", m.HeapAlloc/1024)
	fmt.Printf("   GC cycles: %d\n", m.NumGC)
}

// Helper functions
// ===============

// Stack allocation functions
func returnValue() int {
	x := 42
	return x
}

func returnStruct() Point {
	return Point{X: 10, Y: 20}
}

// Heap allocation functions
func returnPointer() *int {
	x := 42
	return &x
}

func returnStructPointer() *Point {
	return &Point{X: 10, Y: 20}
}

// Interface definitions
type Writer interface {
	Write(string)
}

type Reader interface {
	Read() string
}

type ConsoleWriter struct{}

func (cw *ConsoleWriter) Write(s string) {
	fmt.Printf("   ConsoleWriter: %s\n", s)
}

type StringReader struct {
	data string
}

func (sr *StringReader) Read() string {
	return sr.data
}

// Closure functions
func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// Types
type Point struct {
	X, Y int
}

type Person struct {
	Name string
	Age  int
}
