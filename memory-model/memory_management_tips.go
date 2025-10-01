package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Memory Management Tips and Best Practices
// =======================================

func main() {
	fmt.Println("=== Memory Management Tips ===")
	
	// General memory management principles
	generalPrinciples()
	
	// Stack optimization techniques
	stackOptimization()
	
	// Heap optimization techniques
	heapOptimization()
	
	// Memory profiling and debugging
	memoryProfiling()
	
	// Common memory pitfalls
	commonPitfalls()
	
	// Advanced memory management
	advancedTechniques()
}

// General Memory Management Principles
// ===================================
func generalPrinciples() {
	fmt.Println("\n1. GENERAL PRINCIPLES:")
	
	fmt.Println("   ✓ Prefer stack allocation over heap allocation")
	fmt.Println("   ✓ Minimize the number of heap allocations")
	fmt.Println("   ✓ Reuse objects when possible")
	fmt.Println("   ✓ Use appropriate data structures for your use case")
	fmt.Println("   ✓ Monitor memory usage and profile regularly")
	fmt.Println("   ✓ Understand escape analysis rules")
	fmt.Println("   ✓ Use object pools for frequently allocated objects")
}

// Stack Optimization Techniques
// ============================
func stackOptimization() {
	fmt.Println("\n2. STACK OPTIMIZATION TECHNIQUES:")
	
	// Use value types when possible
	valueTypes()
	
	// Avoid unnecessary pointers
	avoidUnnecessaryPointers()
	
	// Use value receivers
	valueReceivers()
	
	// Local variable optimization
	localVariableOptimization()
}

func valueTypes() {
	fmt.Println("   Use Value Types When Possible:")
	
	// GOOD: Value type
	type Point struct {
		X, Y int
	}
	
	func() {
		p := Point{X: 10, Y: 20}  // Stack allocation
		p.X += 5
		fmt.Printf("     Point: %+v (stack)\n", p)
	}()
	
	// BAD: Unnecessary pointer
	func() {
		p := &Point{X: 10, Y: 20}  // Heap allocation
		p.X += 5
		fmt.Printf("     Point: %+v (heap)\n", *p)
	}()
}

func avoidUnnecessaryPointers() {
	fmt.Println("   Avoid Unnecessary Pointers:")
	
	// GOOD: Pass by value for small structs
	func() {
		type SmallStruct struct {
			Value int
		}
		
		s := SmallStruct{Value: 42}
		result := processSmallStruct(s)  // Stack allocation
		fmt.Printf("     Small struct: %d (stack)\n", result)
	}()
	
	// BAD: Unnecessary pointer for small struct
	func() {
		type SmallStruct struct {
			Value int
		}
		
		s := &SmallStruct{Value: 42}
		result := processSmallStructPointer(s)  // Heap allocation
		fmt.Printf("     Small struct: %d (heap)\n", result)
	}()
}

func valueReceivers() {
	fmt.Println("   Use Value Receivers:")
	
	// GOOD: Value receiver for small structs
	type Point struct {
		X, Y int
	}
	
	func (p Point) Distance() float64 {
		return float64(p.X*p.X + p.Y*p.Y)
	}
	
	p := Point{X: 3, Y: 4}
	distance := p.Distance()
	fmt.Printf("     Distance: %f (value receiver)\n", distance)
}

func localVariableOptimization() {
	fmt.Println("   Local Variable Optimization:")
	
	// GOOD: Declare variables close to where they're used
	func() {
		// Process data
		data := make([]int, 100)
		for i := range data {
			data[i] = i
		}
		
		// Use data immediately
		sum := 0
		for _, v := range data {
			sum += v
		}
		fmt.Printf("     Sum: %d\n", sum)
	}()
}

// Heap Optimization Techniques
// ===========================
func heapOptimization() {
	fmt.Println("\n3. HEAP OPTIMIZATION TECHNIQUES:")
	
	// Pre-allocate slices
	preAllocateSlices()
	
	// Use object pools
	useObjectPools()
	
	// Minimize allocations in loops
	minimizeAllocationsInLoops()
	
	// Use appropriate data structures
	appropriateDataStructures()
}

func preAllocateSlices() {
	fmt.Println("   Pre-allocate Slices:")
	
	// BAD: Growing slice
	func() {
		var slice []int
		for i := 0; i < 1000; i++ {
			slice = append(slice, i)  // Multiple reallocations
		}
		fmt.Printf("     Growing slice: len=%d\n", len(slice))
	}()
	
	// GOOD: Pre-allocate with known capacity
	func() {
		slice := make([]int, 0, 1000)  // Pre-allocate capacity
		for i := 0; i < 1000; i++ {
			slice = append(slice, i)  // No reallocations
		}
		fmt.Printf("     Pre-allocated slice: len=%d, cap=%d\n", len(slice), cap(slice))
	}()
}

func useObjectPools() {
	fmt.Println("   Use Object Pools:")
	
	// Object pool for frequently allocated objects
	pool := make(chan *Person, 10)
	
	// Get from pool
	person := getFromPool(pool)
	person.Name = "Alice"
	person.Age = 30
	
	// Return to pool
	returnToPool(pool, person)
	fmt.Println("     Object pool reduces allocation overhead")
}

func minimizeAllocationsInLoops() {
	fmt.Println("   Minimize Allocations in Loops:")
	
	// BAD: Allocating in loop
	func() {
		var results []string
		for i := 0; i < 100; i++ {
			// BAD: String concatenation creates new strings
			results = append(results, fmt.Sprintf("Item %d", i))
		}
		fmt.Printf("     Bad loop: %d items\n", len(results))
	}()
	
	// GOOD: Pre-allocate and use builder pattern
	func() {
		results := make([]string, 0, 100)  // Pre-allocate
		for i := 0; i < 100; i++ {
			results = append(results, fmt.Sprintf("Item %d", i))
		}
		fmt.Printf("     Good loop: %d items\n", len(results))
	}()
}

func appropriateDataStructures() {
	fmt.Println("   Use Appropriate Data Structures:")
	
	// For small, fixed-size data: arrays
	func() {
		var coords [3]float64 = [3]float64{1.0, 2.0, 3.0}
		fmt.Printf("     Array: %v\n", coords)
	}()
	
	// For dynamic data: slices
	func() {
		coords := make([]float64, 0, 10)
		coords = append(coords, 1.0, 2.0, 3.0)
		fmt.Printf("     Slice: %v\n", coords)
	}()
	
	// For key-value pairs: maps
	func() {
		config := make(map[string]string)
		config["host"] = "localhost"
		config["port"] = "8080"
		fmt.Printf("     Map: %v\n", config)
	}()
}

// Memory Profiling and Debugging
// ===============================
func memoryProfiling() {
	fmt.Println("\n4. MEMORY PROFILING AND DEBUGGING:")
	
	// Check escape analysis
	checkEscapeAnalysis()
	
	// Monitor memory usage
	monitorMemoryUsage()
	
	// Use profiling tools
	profilingTools()
}

func checkEscapeAnalysis() {
	fmt.Println("   Check Escape Analysis:")
	fmt.Println("     go build -gcflags='-m' your_file.go")
	fmt.Println("     Look for 'escapes to heap' messages")
}

func monitorMemoryUsage() {
	fmt.Println("   Monitor Memory Usage:")
	
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	
	fmt.Printf("     Heap size: %d KB\n", m.HeapAlloc/1024)
	fmt.Printf("     Stack size: %d KB\n", m.StackInuse/1024)
	fmt.Printf("     GC cycles: %d\n", m.NumGC)
	fmt.Printf("     GC time: %v\n", time.Duration(m.PauseTotalNs))
}

func profilingTools() {
	fmt.Println("   Use Profiling Tools:")
	fmt.Println("     go tool pprof -http=:8080 profile.out")
	fmt.Println("     go tool trace trace.out")
	fmt.Println("     runtime/pprof package for programmatic profiling")
}

// Common Memory Pitfalls
// =====================
func commonPitfalls() {
	fmt.Println("\n5. COMMON MEMORY PITFALLS:")
	
	// Memory leaks
	memoryLeaks()
	
	// Unnecessary allocations
	unnecessaryAllocations()
	
	// Large object allocation
	largeObjectAllocation()
	
	// Goroutine memory issues
	goroutineMemoryIssues()
}

func memoryLeaks() {
	fmt.Println("   Memory Leaks:")
	
	// BAD: Global slice that grows indefinitely
	// var globalSlice []int
	// func addToGlobal(item int) {
	//     globalSlice = append(globalSlice, item)  // Never cleaned up
	// }
	
	// GOOD: Use local slices or clean up periodically
	func() {
		localSlice := make([]int, 0, 100)
		for i := 0; i < 10; i++ {
			localSlice = append(localSlice, i)
		}
		// localSlice is automatically cleaned up
		fmt.Printf("     Local slice: %v (auto cleanup)\n", localSlice)
	}()
}

func unnecessaryAllocations() {
	fmt.Println("   Unnecessary Allocations:")
	
	// BAD: String concatenation in loop
	func() {
		var result string
		for i := 0; i < 100; i++ {
			result += fmt.Sprintf("%d ", i)  // Creates new string each time
		}
		fmt.Printf("     String concatenation: %d chars\n", len(result))
	}()
	
	// GOOD: Use strings.Builder
	func() {
		var builder strings.Builder
		for i := 0; i < 100; i++ {
			builder.WriteString(fmt.Sprintf("%d ", i))
		}
		result := builder.String()
		fmt.Printf("     strings.Builder: %d chars\n", len(result))
	}()
}

func largeObjectAllocation() {
	fmt.Println("   Large Object Allocation:")
	
	// Large objects might escape to heap
	func() {
		large := make([]int, 1000000)  // 1M integers
		fmt.Printf("     Large slice: %d elements\n", len(large))
		// This will likely escape to heap
	}()
}

func goroutineMemoryIssues() {
	fmt.Println("   Goroutine Memory Issues:")
	
	// Each goroutine has its own stack
	var wg sync.WaitGroup
	
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// Each goroutine uses stack memory
			var local [1000]int
			for j := range local {
				local[j] = id * 1000 + j
			}
			fmt.Printf("     Goroutine %d: %d elements\n", id, len(local))
		}(i)
	}
	
	wg.Wait()
}

// Advanced Memory Management
// ==========================
func advancedTechniques() {
	fmt.Println("\n6. ADVANCED TECHNIQUES:")
	
	// Memory alignment
	memoryAlignment()
	
	// Custom memory management
	customMemoryManagement()
	
	// Memory-mapped files
	memoryMappedFiles()
	
	// Zero-copy techniques
	zeroCopyTechniques()
}

func memoryAlignment() {
	fmt.Println("   Memory Alignment:")
	
	type AlignedStruct struct {
		Field1 int32  // 4 bytes
		Field2 int64  // 8 bytes
		Field3 int32  // 4 bytes
	}
	
	var s AlignedStruct
	fmt.Printf("     Struct size: %d bytes\n", unsafe.Sizeof(s))
	fmt.Println("     Go automatically handles alignment")
}

func customMemoryManagement() {
	fmt.Println("   Custom Memory Management:")
	
	// Use sync.Pool for object reuse
	pool := &sync.Pool{
		New: func() interface{} {
			return &Person{}
		},
	}
	
	// Get from pool
	person := pool.Get().(*Person)
	person.Name = "John"
	person.Age = 30
	
	// Return to pool
	pool.Put(person)
	fmt.Println("     sync.Pool for object reuse")
}

func memoryMappedFiles() {
	fmt.Println("   Memory-Mapped Files:")
	fmt.Println("     Use mmap for large file operations")
	fmt.Println("     Reduces memory usage for large datasets")
	fmt.Println("     Example: github.com/edsrzf/mmap-go")
}

func zeroCopyTechniques() {
	fmt.Println("   Zero-Copy Techniques:")
	
	// Use []byte slices to avoid copying
	func() {
		data := []byte("Hello, World!")
		slice1 := data[0:5]   // "Hello"
		slice2 := data[7:12] // "World"
		
		fmt.Printf("     Slice1: %s\n", slice1)
		fmt.Printf("     Slice2: %s\n", slice2)
		fmt.Println("     No copying, just different views")
	}()
}

// Helper functions
// ===============
func processSmallStruct(s SmallStruct) int {
	return s.Value * 2
}

func processSmallStructPointer(s *SmallStruct) int {
	return s.Value * 2
}

func getFromPool(pool chan *Person) *Person {
	select {
	case person := <-pool:
		return person
	default:
		return &Person{}
	}
}

func returnToPool(pool chan *Person, person *Person) {
	select {
	case pool <- person:
		// Returned to pool
	default:
		// Pool full, let GC handle it
	}
}

type SmallStruct struct {
	Value int
}

type Person struct {
	Name string
	Age  int
}
