# Go Structs

This folder contains comprehensive examples of Go's struct types and their usage.

## 📁 Files

- **`go_structs.go`** - Complete guide to Go structs

## 🎯 What You'll Learn

### **Basic Struct Definition and Usage**
- Structs are composite data types that group related fields (variables) under a single name
- Define using `type Name struct { field1 type1; field2 type2 }` syntax
- Access and modify fields using dot notation: `person.name` or `person.age = 30`
- Think of structs as blueprints for creating objects with specific properties

### **Struct Initialization**
- Zero-value initialization: `var p Person` gives all fields their zero values (0, "", nil, false, etc.)
- Named field initialization: `Person{name: "John", age: 30}` - explicit and self-documenting (recommended)
- Positional initialization: `Person{"John", 30}` - order-dependent, breaks if field order changes (avoid)
- Partial initialization allowed: `Person{name: "John"}` - unspecified fields get zero values
- Using `new(Person)` allocates memory, zeros all fields, and returns `*Person` pointer
- Shorthand `&Person{name: "John"}` creates struct and returns pointer in one expression

### **Struct Fields and Access**
- Fields are accessed and modified directly using dot notation: `person.name = "Jane"`
- Pointer receivers automatically dereference: `ptr.name` works instead of verbose `(*ptr).name`
- Go's syntactic sugar makes pointer operations cleaner compared to C/C++
- Struct fields can be modified if: (1) struct variable is not const, or (2) you have a pointer to it
- Exported fields (capitalized) are accessible from other packages; unexported (lowercase) are private

### **Anonymous Structs**
- Declare and use structs on-the-fly without creating a named type definition
- Perfect for one-time use cases like temporary configurations or grouping test data
- Syntax: `var config struct { host string; port int }` or inline `struct{x, y int}{10, 20}`
- Commonly used in table-driven tests where each test case has unique fields
- Reduces boilerplate when you don't need to reuse the type elsewhere

### **Nested Structs**
- Structs can contain other structs as fields, creating hierarchical data structures
- Access nested fields by chaining dot operators: `person.address.city` or `user.profile.settings.theme`
- Ideal for modeling real-world relationships: User has Address, Order has Items, etc.
- Can nest to any depth, though excessive nesting (3-4+ levels) can hurt readability
- Both embedded types and named fields containing structs are supported

### **Struct Methods**
- Methods are functions attached to a type via a receiver parameter before the function name
- Value receiver `func (p Person) GetName()` operates on a copy - original struct is unchanged
- Pointer receiver `func (p *Person) SetName(name string)` operates on original - can modify struct state
- Choose pointer receivers when: (1) modifying the receiver, (2) struct is large (avoid copying), (3) maintaining consistency across all methods
- Method sets: value types can call value receiver methods; pointer types can call both value and pointer receivers
- Go automatically handles pointer/value conversion when calling methods

### **Struct Embedding (Composition)**
- Go favors composition over inheritance - embed types rather than extending classes
- Embed without field names: `type Employee struct { Person; salary int }` - Person's fields "promoted"
- Promoted fields/methods accessible directly: `emp.Name` instead of `emp.Person.Name`
- Override embedded methods by defining same method on outer struct - outer takes precedence
- Can embed multiple types, but method/field name conflicts require explicit qualification
- This pattern enables "has-a" relationships and interface satisfaction through embedding

### **Struct Tags**
- Tags are metadata strings enclosed in backticks that annotate struct fields
- Syntax: `Name string \`json:"name" db:"user_name" validate:"required"\``
- Accessed via reflection, primarily used by encoding/json, encoding/xml, ORMs, validators
- Common options: `omitempty` (skip zero values), `-` (ignore field), `string` (force string encoding)
- Multiple tags can be combined: `\`json:"email,omitempty" xml:"Email" db:"email_addr"\``
- Only exported (capitalized) fields are serialized - unexported fields are ignored by encoders

### **Struct Comparison**
- Structs support `==` and `!=` operators only when all fields are comparable types
- Comparable types: all primitives (int, string, bool, etc.), pointers, arrays, channels, interfaces
- Non-comparable types: slices, maps, functions - including these makes entire struct non-comparable
- Comparison is field-by-field: `p1 == p2` true if all corresponding fields are equal
- Pointer fields compare addresses, not values: `&Person{} != &Person{}` even with identical data
- For deep equality checking (comparing slice/map contents), use `reflect.DeepEqual()` (slower)

### **Struct Memory Layout**
- Compiler aligns fields to memory addresses that are multiples of field size for CPU efficiency
- Field ordering matters: poor order leads to padding bytes between fields, increasing struct size
- Example: `{bool, int64, bool}` uses 24 bytes vs `{int64, bool, bool}` uses 16 bytes
- Use `unsafe.Sizeof()` to get total size, `unsafe.Alignof()` for alignment, `unsafe.Offsetof()` for field positions
- Best practice: order fields from largest to smallest (int64, int32, int16, bool) to minimize wasted space
- Padding is automatic and invisible - compiler inserts bytes to maintain alignment requirements

## 🚀 How to Run

```bash
cd structs
go run go_structs.go
```

## 📚 Key Takeaways

- **Structs are value types** - passed by value (entire struct copied), use pointers to pass by reference and avoid copying overhead
- **Methods have two receiver types** - value receivers work on copies (read-only), pointer receivers work on originals (can mutate state)
- **Composition over inheritance** - embed types to "inherit" behavior without classical OOP inheritance hierarchies
- **Tags enable metadata** - critical for JSON/XML marshaling, database ORMs, validation frameworks, and custom tooling
- **Memory layout is optimizable** - field ordering dramatically affects struct size; order largest-to-smallest to minimize padding waste
- **Exported vs unexported fields** - capitalized fields are public (package-accessible), lowercase are private to the package

## 🔗 Related Topics

- **Primitive Types** - See `../primitives/` folder
- **Pointers** - See `../pointers/` folder
- **Advanced Concepts** - See `../advanced-concepts/` folder
