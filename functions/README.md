# Go Functions

This folder contains comprehensive examples of Go's function features and usage patterns.

## 📁 Files

- **`go_functions.go`** - Complete guide to Go functions

## 🎯 What You'll Learn

### **Basic Function Declaration and Calling**
- Functions are first-class citizens in Go - they can be assigned to variables, passed as arguments, and returned from other functions
- Define using `func name(params) returnType { body }` syntax
- Parameters specify input types: `func add(a int, b int)` or shorthand `func add(a, b int)`
- Return type follows parameters: `func getName() string` or omit for no return
- Call functions using parentheses: `result := add(5, 3)` executes the function and captures return value

### **Multiple Parameters and Return Values**
- Go functions can return multiple values - a powerful feature for error handling and complex operations
- Common pattern: `func divide(a, b float64) (float64, error)` returns result and potential error
- Calling code checks error: `result, err := divide(10, 2); if err != nil { handle error }`
- Use blank identifier `_` to ignore unwanted return values: `result, _ := divide(10, 2)`
- Multiple returns eliminate need for out parameters or reference passing common in other languages
- Best practice: return error as last value in multi-return functions

### **Named Return Values**
- Declare return variable names in function signature: `func getCoords() (x, y int)`
- Named returns are automatically initialized to zero values at function start
- Use naked return statement: `return` without arguments returns all named values
- Makes code self-documenting - signature shows what values represent: `(width, height int)` vs `(int, int)`
- Can still use explicit returns: `return width, height` even with named returns
- Warning: naked returns in long functions can harm readability - use sparingly

### **Variadic Functions**
- Accept variable number of arguments using `...` syntax: `func sum(numbers ...int) int`
- Variadic parameter must be last in parameter list: `func printf(format string, args ...interface{})`
- Inside function, variadic parameter is a slice: iterate with `for _, num := range numbers`
- Pass slice to variadic function by unpacking: `sum(mySlice...)` spreads slice elements as arguments
- Common in Go standard library: `fmt.Println`, `append`, `fmt.Sprintf` all use variadic parameters
- Zero arguments is valid: calling `sum()` with no args passes empty slice

### **Functions as Values**
- Functions have types based on signature: `func(int, int) int` is the type of a function taking two ints, returning int
- Assign functions to variables: `var myFunc func(int) int = someFunction`
- Pass functions as arguments (callback pattern): `func apply(fn func(int) int, x int)`
- Return functions from functions (factory pattern): `func makeAdder(x int) func(int) int`
- Enables functional programming patterns: map, filter, reduce, decorators, middleware
- Type safety: compiler ensures function signatures match when assigned/passed

### **Anonymous Functions**
- Define functions inline without names: `func(x int) { fmt.Println(x) }`
- Immediately invoked: `func() { fmt.Println("Hello") }()` defines and calls in one expression
- Assign to variables: `square := func(x int) int { return x * x }`
- Useful for short, one-off operations that don't merit a separate named function
- Common in goroutines: `go func() { doWork() }()` spawns concurrent task
- Can capture variables from surrounding scope (becomes a closure)

### **Closures**
- Closure = function that references variables from outer (enclosing) scope
- Captured variables are shared, not copied: closure can modify original variables
- Each closure instance has independent captured state: calling factory twice creates separate closures
- Classic example: `makeCounter() func() int` returns function that increments private counter
- Enables data encapsulation without classes: captured variables act as private fields
- Warning: closures in loops capture loop variable reference, not value - create new variable per iteration

### **Recursion**
- Function calling itself to solve problems by breaking into smaller subproblems
- Requires base case (termination condition) to prevent infinite recursion: `if n <= 1 { return 1 }`
- Recursive case calls function with modified input: `return n * factorial(n-1)`
- Go doesn't optimize tail recursion - stack overflow possible with deep recursion
- Common uses: tree/graph traversal, divide-and-conquer algorithms, mathematical sequences
- Alternative: iterative solutions often more efficient in Go due to lack of tail call optimization

### **Defer Statements**
- `defer` schedules function call to execute when surrounding function returns (success or panic)
- Deferred calls execute in LIFO order (Last In, First Out): last defer runs first
- Arguments evaluated immediately: `defer fmt.Println(x)` captures x's current value, not value at return time
- Primary use: cleanup/resource management - close files, unlock mutexes, recover from panics
- Always executes even if function panics: ensures cleanup happens
- Multiple defers stack up: useful for paired operations (open/close, lock/unlock, begin/commit)

### **Higher-Order Functions**
- Functions that take other functions as arguments or return functions
- Map: transform each element - `mapInts(numbers, func(x int) int { return x * 2 })`
- Filter: select elements matching predicate - `filterInts(numbers, func(x int) bool { return x > 5 })`
- Reduce: combine elements into single value - `reduceInts(numbers, 0, func(acc, x int) int { return acc + x })`
- Enables functional programming style: chain operations, compose transformations
- Go doesn't have generics for these (until Go 1.18+) - implement per type or use interfaces
- Trade-off: more abstract/reusable code vs potential performance overhead from function calls

## 🚀 How to Run

```bash
cd functions
go run go_functions.go
```

## 📚 Key Takeaways

- **Functions are first-class** - can be assigned, passed, returned like any other value
- **Multiple return values** - idiomatic error handling via `(result, error)` pattern; no exceptions needed
- **Named returns improve clarity** - self-documenting signatures, but use sparingly (naked returns hurt readability in long functions)
- **Variadic functions are flexible** - variable arguments enable APIs like `fmt.Println`, use `...` to unpack slices
- **Closures capture scope** - powerful for callbacks, factories, encapsulation; watch out for loop variable capture bugs
- **Defer ensures cleanup** - always runs on function exit (return or panic); LIFO order; evaluated arguments immediately
- **Higher-order functions enable abstraction** - map/filter/reduce patterns for data transformation; functional programming style

## 🔗 Related Topics

- **Primitive Types** - See `../primitives/` folder
- **Structs** - See `../structs/` folder (methods vs functions)
- **Pointers** - See `../pointers/` folder (pointer receivers in methods)
- **Advanced Concepts** - See `../advanced-concepts/` folder (interfaces, goroutines)

