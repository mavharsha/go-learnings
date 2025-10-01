package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Performance Implications of Stack vs Heap
// =========================================

func main() {
	fmt.Println("=== Performance Implications ===")
	
	// Memory allocation performance
	allocationPerformance()
	
	// Garbage collection impact
	garbageCollectionImpact()
	
	// Memory usage patterns
	memoryUsagePatterns()
	
	// Concurrency implications
	concurrencyImplications()
	
	// Best practices for performance
	performanceBestPractices()
}

// Memory Allocation Performance
// =============================
func allocationPerformance() {
	fmt.Println("\n1. ALLOCATION PERFORMANCE:")
	
	// Stack allocation benchmark
	stackBenchmark()
	
	// Heap allocation benchmark
	heapBenchmark()
	
	// Mixed allocation benchmark
	mixedBenchmark()
}

func stackBenchmark() {
	fmt.Println("   Stack Allocation (Fast):")
	
	iterations := 1000000
	start := time.Now()
	
	for i := 0; i < iterations; i++ {
		// Stack operations - very fast
		var a, b, c int = i, i+1, i+2
		var result int = a + b + c
		_ = result  // Prevent optimization
	}
	
	duration := time.Since(start)
	fmt.Printf("     %d iterations in %v\n", iterations, duration)
	fmt.Printf("     Average: %v per operation\n", duration/time.Duration(iterations))
}

func heapBenchmark() {
	fmt.Println("   Heap Allocation (Slower):")
	
	iterations := 1000000
	start := time.Now()
	
	for i := 0; i < iterations; i++ {
		// Heap operations - slower
		a := new(int)
		b := new(int)
		c := new(int)
		*a, *b, *c = i, i+1, i+2
		result := new(int)
		*result = *a + *b + *c
		_ = result  // Prevent optimization
	}
	
	duration := time.Since(start)
	fmt.Printf("     %d iterations in %v\n", iterations, duration)
	fmt.Printf("     Average: %v per operation\n", duration/time.Duration(iterations))
}

func mixedBenchmark() {
	fmt.Println("   Mixed Allocation:")
	
	iterations := 1000000
	start := time.Now()
	
	for i := 0; i < iterations; i++ {
		// Mix of stack and heap
		var stackVar int = i
		heapVar := new(int)
		*heapVar = i + 1
		result := stackVar + *heapVar
		_ = result
	}
	
	duration := time.Since(start)
	fmt.Printf("     %d iterations in %v\n", iterations, duration)
	fmt.Printf("     Average: %v per operation\n", duration/time.Duration(iterations))
}

// Garbage Collection Impact
// =========================
func garbageCollectionImpact() {
	fmt.Println("\n2. GARBAGE COLLECTION IMPACT:")
	
	// Show GC stats before
	var m1 runtime.MemStats
	runtime.ReadMemStats(&m1)
	fmt.Printf("   Before GC: Heap=%d KB, GC cycles=%d\n", m1.HeapAlloc/1024, m1.NumGC)
	
	// Create heap allocations
	createHeapAllocations()
	
	// Force GC
	runtime.GC()
	
	// Show GC stats after
	var m2 runtime.MemStats
	runtime.ReadMemStats(&m2)
	fmt.Printf("   After GC: Heap=%d KB, GC cycles=%d\n", m2.HeapAlloc/1024, m2.NumGC)
	fmt.Printf("   GC overhead: %d KB freed\n", (m1.HeapAlloc-m2.HeapAlloc)/1024)
}

func createHeapAllocations() {
	// Create many heap allocations
	for i := 0; i < 10000; i++ {
		// These will be garbage collected
		ptr := new(int)
		*ptr = i
		slice := make([]int, 100)
		_ = slice
		_ = ptr
	}
}

// Memory Usage Patterns
// ====================
func memoryUsagePatterns() {
	fmt.Println("\n3. MEMORY USAGE PATTERNS:")
	
	// Stack memory characteristics
	stackCharacteristics()
	
	// Heap memory characteristics
	heapCharacteristics()
	
	// Memory fragmentation
	memoryFragmentation()
}

func stackCharacteristics() {
	fmt.Println("   Stack Characteristics:")
	fmt.Println("     - Fixed size per goroutine (typically 1-8MB)")
	fmt.Println("     - Fast allocation/deallocation")
	fmt.Println("     - No fragmentation")
	fmt.Println("     - Automatic cleanup on function return")
	fmt.Println("     - Limited by stack size")
}

func heapCharacteristics() {
	fmt.Println("   Heap Characteristics:")
	fmt.Println("     - Dynamic size")
	fmt.Println("     - Slower allocation/deallocation")
	fmt.Println("     - Can fragment over time")
	fmt.Println("     - Managed by garbage collector")
	fmt.Println("     - Can grow to system limits")
}

func memoryFragmentation() {
	fmt.Println("   Memory Fragmentation:")
	
	// Demonstrate fragmentation
	var pointers []*int
	
	// Allocate many small objects
	for i := 0; i < 1000; i++ {
		ptr := new(int)
		*ptr = i
		pointers = append(pointers, ptr)
	}
	
	// Free every other object (simulate fragmentation)
	for i := 0; i < len(pointers); i += 2 {
		pointers[i] = nil
	}
	
	fmt.Printf("     Allocated %d objects\n", len(pointers))
	fmt.Println("     Fragmentation occurs when objects are freed non-contiguously")
}

// Concurrency Implications
// =======================
func concurrencyImplications() {
	fmt.Println("\n4. CONCURRENCY IMPLICATIONS:")
	
	// Stack per goroutine
	stackPerGoroutine()
	
	// Heap sharing
	heapSharing()
	
	// Memory contention
	memoryContention()
}

func stackPerGoroutine() {
	fmt.Println("   Stack per Goroutine:")
	
	// Each goroutine has its own stack
	var wg sync.WaitGroup
	
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// Each goroutine has its own stack
			var localVar int = id * 100
			fmt.Printf("     Goroutine %d: localVar=%d\n", id, localVar)
		}(i)
	}
	
	wg.Wait()
	fmt.Println("     Each goroutine has independent stack memory")
}

func heapSharing() {
	fmt.Println("   Heap Sharing:")
	
	// Shared heap memory
	sharedData := make([]int, 10)
	var wg sync.WaitGroup
	var mu sync.Mutex
	
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			mu.Lock()
			sharedData[id] = id * 10
			fmt.Printf("     Goroutine %d: sharedData[%d]=%d\n", id, id, sharedData[id])
			mu.Unlock()
		}(i)
	}
	
	wg.Wait()
	fmt.Println("     Heap memory can be shared between goroutines")
}

func memoryContention() {
	fmt.Println("   Memory Contention:")
	
	// Demonstrate memory contention
	var wg sync.WaitGroup
	start := time.Now()
	
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Heavy heap allocation
			for j := 0; j < 1000; j++ {
				ptr := new(int)
				*ptr = j
				_ = ptr
			}
		}()
	}
	
	wg.Wait()
	duration := time.Since(start)
	fmt.Printf("     Concurrent heap allocation took: %v\n", duration)
	fmt.Println("     Contention occurs when multiple goroutines allocate heap memory")
}

// Performance Best Practices
// ==========================
func performanceBestPractices() {
	fmt.Println("\n5. PERFORMANCE BEST PRACTICES:")
	
	// Prefer stack allocation
	preferStackAllocation()
	
	// Minimize heap allocations
	minimizeHeapAllocations()
	
	// Use object pools
	useObjectPools()
	
	// Profile memory usage
	profileMemoryUsage()
}

func preferStackAllocation() {
	fmt.Println("   Prefer Stack Allocation:")
	
	// GOOD: Stack allocation
	func() {
		var data [1000]int
		for i := range data {
			data[i] = i
		}
		fmt.Println("     Stack allocation: Fast, automatic cleanup")
	}()
	
	// BAD: Unnecessary heap allocation
	func() {
		data := make([]int, 1000)
		for i := range data {
			data[i] = i
		}
		fmt.Println("     Heap allocation: Slower, requires GC")
	}()
}

func minimizeHeapAllocations() {
	fmt.Println("   Minimize Heap Allocations:")
	
	// BAD: Many small allocations
	// for i := 0; i < 1000; i++ {
	//     ptr := new(int)
	//     *ptr = i
	// }
	
	// GOOD: Fewer, larger allocations
	slice := make([]int, 1000)
	for i := range slice {
		slice[i] = i
	}
	fmt.Println("     Use slices instead of many small allocations")
}

func useObjectPools() {
	fmt.Println("   Use Object Pools:")
	
	// Object pool for frequently allocated objects
	pool := make(chan *Person, 10)
	
	// Get from pool
	person := getFromPool(pool)
	person.Name = "John"
	person.Age = 30
	
	// Return to pool
	returnToPool(pool, person)
	fmt.Println("     Object pools reduce allocation overhead")
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

func profileMemoryUsage() {
	fmt.Println("   Profile Memory Usage:")
	
	// Show current memory stats
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	
	fmt.Printf("     Heap size: %d KB\n", m.HeapAlloc/1024)
	fmt.Printf("     Stack size: %d KB\n", m.StackInuse/1024)
	fmt.Printf("     GC cycles: %d\n", m.NumGC)
	fmt.Printf("     GC time: %v\n", time.Duration(m.PauseTotalNs))
	
	fmt.Println("     Use 'go tool pprof' for detailed profiling")
}

type Person struct {
	Name string
	Age  int
}
