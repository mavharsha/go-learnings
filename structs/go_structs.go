package main

import (
	"fmt"
	"unsafe"
)

// Go Structs - Complete Guide
// =========================
// This file demonstrates Go structs with comprehensive examples

// Struct definitions
// ==================
type Circle struct {
	Radius float64
}

type Animal struct {
	Name string
	Age  int
}

type Dog struct {
	Animal  // Embedded struct
	Breed   string
	IsGood  bool
}

func main() {
	fmt.Println("=== Go Structs ===")
	
	// 1. Basic struct definition and usage
	basicStructs()
	
	// 2. Struct initialization
	structInitialization()
	
	// 3. Struct fields and access
	structFields()
	
	// 4. Anonymous structs
	anonymousStructs()
	
	// 5. Nested structs
	nestedStructs()
	
	// 6. Struct methods
	structMethods()
	
	// 7. Struct embedding (composition)
	structEmbedding()
	
	// 8. Struct tags
	structTags()
	
	// 9. Struct comparison
	structComparison()
	
	// 10. Struct memory layout
	structMemoryLayout()
}

// 1. Basic Struct Definition and Usage
// ====================================
func basicStructs() {
	fmt.Println("\n1. BASIC STRUCTS:")
	
	// Define a struct
	type Person struct {
		Name string
		Age  int
		City string
	}
	
	// Create struct instances
	var person1 Person
	person1.Name = "Alice"
	person1.Age = 30
	person1.City = "New York"
	
	person2 := Person{
		Name: "Bob",
		Age:  25,
		City: "London",
	}
	
	fmt.Printf("   Person 1: %+v\n", person1)
	fmt.Printf("   Person 2: %+v\n", person2)
	fmt.Printf("   Person 1 name: %s\n", person1.Name)
	fmt.Printf("   Person 2 age: %d\n", person2.Age)
}

// 2. Struct Initialization
// =========================
func structInitialization() {
	fmt.Println("\n2. STRUCT INITIALIZATION:")
	
	type Point struct {
		X, Y int
	}
	
	// Different ways to initialize structs
	
	// 1. Zero value initialization
	var p1 Point
	fmt.Printf("   Zero value: %+v\n", p1)
	
	// 2. Field initialization
	var p2 Point = Point{X: 10, Y: 20}
	fmt.Printf("   Field init: %+v\n", p2)
	
	// 3. Short declaration with field names
	p3 := Point{X: 5, Y: 15}
	fmt.Printf("   Short decl: %+v\n", p3)
	
	// 4. Positional initialization (order matters)
	p4 := Point{100, 200}
	fmt.Printf("   Positional: %+v\n", p4)
	
	// 5. Partial initialization (remaining fields get zero values)
	p5 := Point{X: 50}
	fmt.Printf("   Partial init: %+v\n", p5)
	
	// 6. Using new() - returns pointer
	p6 := new(Point)
	p6.X = 75
	p6.Y = 85
	fmt.Printf("   New pointer: %+v\n", *p6)
	
	// 7. Address of struct literal
	p7 := &Point{X: 90, Y: 95}
	fmt.Printf("   Address of literal: %+v\n", *p7)
}

// 3. Struct Fields and Access
// ============================
func structFields() {
	fmt.Println("\n3. STRUCT FIELDS AND ACCESS:")
	
	type Rectangle struct {
		Width  float64
		Height float64
	}
	
	rect := Rectangle{Width: 10.5, Height: 5.5}
	
	// Accessing fields
	fmt.Printf("   Width: %f\n", rect.Width)
	fmt.Printf("   Height: %f\n", rect.Height)
	
	// Modifying fields
	rect.Width = 15.0
	rect.Height = 8.0
	fmt.Printf("   Modified: %+v\n", rect)
	
	// Field access with pointers
	rectPtr := &rect
	fmt.Printf("   Width via pointer: %f\n", rectPtr.Width)
	fmt.Printf("   Height via pointer: %f\n", rectPtr.Height)
	
	// Dereferencing pointer to access fields
	fmt.Printf("   Width via dereference: %f\n", (*rectPtr).Width)
	fmt.Printf("   Height via dereference: %f\n", (*rectPtr).Height)
}

// 4. Anonymous Structs
// ====================
func anonymousStructs() {
	fmt.Println("\n4. ANONYMOUS STRUCTS:")
	
	// Anonymous struct - no type name
	person := struct {
		Name string
		Age  int
	}{
		Name: "Charlie",
		Age:  35,
	}
	
	fmt.Printf("   Anonymous struct: %+v\n", person)
	
	// Anonymous struct with embedded fields
	config := struct {
		Host string
		Port int
		SSL  bool
	}{
		Host: "localhost",
		Port: 8080,
		SSL:  true,
	}
	
	fmt.Printf("   Config: %+v\n", config)
	
	// Anonymous struct as function parameter
	processAnonymousStruct(struct {
		ID   int
		Name string
	}{
		ID:   1,
		Name: "Test",
	})
}

// 5. Nested Structs
// =================
func nestedStructs() {
	fmt.Println("\n5. NESTED STRUCTS:")
	
	// Define nested structs
	type Address struct {
		Street string
		City   string
		Zip    string
	}
	
	type Employee struct {
		ID      int
		Name    string
		Address Address
		Salary  float64
	}
	
	// Create nested struct
	emp := Employee{
		ID:   1,
		Name: "John Doe",
		Address: Address{
			Street: "123 Main St",
			City:   "New York",
			Zip:    "10001",
		},
		Salary: 75000.0,
	}
	
	fmt.Printf("   Employee: %+v\n", emp)
	fmt.Printf("   Employee name: %s\n", emp.Name)
	fmt.Printf("   Employee city: %s\n", emp.Address.City)
	fmt.Printf("   Employee street: %s\n", emp.Address.Street)
	
	// Modify nested fields
	emp.Address.City = "Boston"
	emp.Salary = 80000.0
	fmt.Printf("   Updated employee: %+v\n", emp)
}

// 6. Struct Methods
// =================
func structMethods() {
	fmt.Println("\n6. STRUCT METHODS:")
	
	// Use Circle struct with methods
	
	// Create circle and use methods
	circle := Circle{Radius: 5.0}
	fmt.Printf("   Circle radius: %f\n", circle.Radius)
	fmt.Printf("   Circle area: %f\n", circle.Area())
	
	// Use pointer receiver method
	circle.SetRadius(10.0)
	fmt.Printf("   Updated radius: %f\n", circle.Radius)
	fmt.Printf("   Updated area: %f\n", circle.Area())
	
	// Use value receiver method
	newCircle := circle.DoubleRadius()
	fmt.Printf("   New circle radius: %f\n", newCircle.Radius)
	fmt.Printf("   Original circle radius: %f\n", circle.Radius)
}

// 7. Struct Embedding (Composition)
// ==================================
func structEmbedding() {
	fmt.Println("\n7. STRUCT EMBEDDING (COMPOSITION):")
	
	// Use Animal and Dog structs
	
	// Create dog instance
	dog := Dog{
		Animal: Animal{
			Name: "Buddy",
			Age:  3,
		},
		Breed:  "Golden Retriever",
		IsGood: true,
	}
	
	fmt.Printf("   Dog: %+v\n", dog)
	fmt.Printf("   Dog name: %s\n", dog.Name)        // Access embedded field
	fmt.Printf("   Dog age: %d\n", dog.Age)          // Access embedded field
	fmt.Printf("   Dog breed: %s\n", dog.Breed)      // Access own field
	fmt.Printf("   Dog speaks: %s\n", dog.Speak())   // Method overriding
	
	// Access embedded struct directly
	fmt.Printf("   Embedded animal: %+v\n", dog.Animal)
}

// 8. Struct Tags
// ==============
func structTags() {
	fmt.Println("\n8. STRUCT TAGS:")
	
	// Struct with tags (used for JSON, XML, etc.)
	type User struct {
		ID       int    `json:"id" xml:"id"`
		Name     string `json:"name" xml:"name"`
		Email    string `json:"email" xml:"email"`
		Password string `json:"-" xml:"-"` // Hidden in JSON/XML
		Age      int    `json:"age,omitempty" xml:"age,omitempty"`
	}
	
	user := User{
		ID:       1,
		Name:     "Alice",
		Email:    "alice@example.com",
		Password: "secret123",
		Age:      30,
	}
	
	fmt.Printf("   User: %+v\n", user)
	fmt.Printf("   User name: %s\n", user.Name)
	fmt.Printf("   User email: %s\n", user.Email)
	
	// Tags are used by packages like encoding/json
	// json.Marshal(user) would use the json tags
	// xml.Marshal(user) would use the xml tags
}

// 9. Struct Comparison
// ====================
func structComparison() {
	fmt.Println("\n9. STRUCT COMPARISON:")
	
	type Point struct {
		X, Y int
	}
	
	// Structs can be compared if all fields are comparable
	p1 := Point{X: 10, Y: 20}
	p2 := Point{X: 10, Y: 20}
	p3 := Point{X: 5, Y: 15}
	
	fmt.Printf("   Point 1: %+v\n", p1)
	fmt.Printf("   Point 2: %+v\n", p2)
	fmt.Printf("   Point 3: %+v\n", p3)
	
	fmt.Printf("   p1 == p2: %t\n", p1 == p2)
	fmt.Printf("   p1 == p3: %t\n", p1 == p3)
	fmt.Printf("   p1 != p3: %t\n", p1 != p3)
	
	// Structs with slices/maps cannot be compared
	type Person struct {
		Name string
		// Hobbies []string  // This would make Person non-comparable
	}
	
	person1 := Person{Name: "Alice"}
	person2 := Person{Name: "Alice"}
	
	fmt.Printf("   Person 1 == Person 2: %t\n", person1 == person2)
}

// 10. Struct Memory Layout
// ========================
func structMemoryLayout() {
	fmt.Println("\n10. STRUCT MEMORY LAYOUT:")
	
	type ExampleStruct struct {
		A bool    // 1 byte
		B int32   // 4 bytes
		C int64   // 8 bytes
		D bool    // 1 byte
	}
	
	var s ExampleStruct
	fmt.Printf("   Struct size: %d bytes\n", unsafe.Sizeof(s))
	fmt.Printf("   Field A offset: %d\n", unsafe.Offsetof(s.A))
	fmt.Printf("   Field B offset: %d\n", unsafe.Offsetof(s.B))
	fmt.Printf("   Field C offset: %d\n", unsafe.Offsetof(s.C))
	fmt.Printf("   Field D offset: %d\n", unsafe.Offsetof(s.D))
	
	// Go automatically handles memory alignment
	// Fields are aligned to their natural boundaries
}

// Helper function for anonymous structs
// ====================================
func processAnonymousStruct(data struct {
	ID   int
	Name string
}) {
	fmt.Printf("   Processing anonymous struct: ID=%d, Name=%s\n", data.ID, data.Name)
}

// Methods for Circle struct
// =========================
func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func (c *Circle) SetRadius(r float64) {
	c.Radius = r
}

func (c Circle) DoubleRadius() Circle {
	return Circle{Radius: c.Radius * 2}
}

// Methods for Animal struct
// =========================
func (a Animal) Speak() string {
	return "Some generic animal sound"
}

// Methods for Dog struct
// ======================
func (d Dog) Speak() string {
	return "Woof! Woof!"
}
