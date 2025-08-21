# Go Structs and Concepts Exercises

This directory contains comprehensive exercises to help you learn Go structs and other important Go concepts.

## Files

- `exercises.go` - The main exercise file with TODOs for you to complete
- `solutions.go` - Complete implementations of all exercises (check your work!)
- `EXERCISES_README.md` - This file with instructions and learning objectives

## Learning Objectives

By completing these exercises, you will learn:

### 1. Basic Structs
- How to define structs with different field types
- Creating constructor functions
- Adding methods to structs
- Understanding value vs pointer receivers

### 2. Pointers and Memory Management
- When to use pointer receivers vs value receivers
- How pointers affect method behavior
- Memory efficiency considerations

### 3. Struct Embedding (Composition)
- How Go achieves inheritance-like behavior through embedding
- Accessing embedded struct fields and methods
- Building complex types from simpler ones

### 4. Interfaces
- Defining and implementing interfaces
- Polymorphism in Go
- Interface satisfaction (implicit implementation)

### 5. Advanced Struct Usage
- Working with slices and maps in structs
- JSON tags for serialization
- Custom types and methods

### 6. Real-World Applications
- Building practical data structures
- Implementing business logic with structs
- Creating clean, maintainable code

## How to Use These Exercises

### Step 1: Start with the Exercises
```bash
cd Go-HelloMac
go run .  # This runs hello.go which includes all files
```

### Step 2: Complete Each Exercise
1. Open `exercises.go` in your editor
2. Find the TODO comments
3. Implement the required structs and methods
4. Uncomment the corresponding test functions in `exercises.go`
5. In `hello.go`, uncomment `runExercises()` to run your implementations

### Step 3: Check Your Solutions
- Compare your implementations with `solutions.go`
- In `hello.go`, uncomment `runSolutions()` to see expected output
- Or run all files together: `go run .`

### Step 4: Experiment and Extend
- Try modifying the exercises
- Add your own methods and fields
- Experiment with different approaches

## Exercise Breakdown

### Exercise 1: Basic Structs (Beginner)
- **Concept**: Struct definition, constructor functions, basic methods
- **Goal**: Create a `Person` struct with name, age, and email
- **Key Learning**: Understanding struct syntax and value receivers

### Exercise 2: Pointers and Methods (Beginner-Intermediate)
- **Concept**: Pointer receivers, modifying struct state
- **Goal**: Create a `BankAccount` with deposit/withdraw functionality
- **Key Learning**: When to use `*` for pointer receivers

### Exercise 3: Struct Embedding (Intermediate)
- **Concept**: Composition over inheritance
- **Goal**: Create an `Employee` struct that embeds `Person`
- **Key Learning**: How embedding provides access to embedded methods

### Exercise 4: Interfaces (Intermediate)
- **Concept**: Interface definition and implementation
- **Goal**: Create a `Shape` interface with `Rectangle` and `Circle` implementations
- **Key Learning**: Polymorphism and implicit interface satisfaction

### Exercise 5: Advanced Structs (Intermediate-Advanced)
- **Concept**: Complex data structures, slices in structs
- **Goal**: Build a `Library` system that manages books
- **Key Learning**: Working with collections and pointer returns

### Exercise 6: JSON Tags (Intermediate)
- **Concept**: Struct tags for serialization
- **Goal**: Define a `Product` struct with JSON tags
- **Key Learning**: How Go handles JSON serialization

### Exercise 7: Custom Types (Intermediate)
- **Concept**: Type definitions and custom methods
- **Goal**: Create a `Temperature` type with conversion methods
- **Key Learning**: Extending basic types with custom behavior

## Challenge Exercises

### Challenge 1: Generic Stack
- **Difficulty**: Advanced
- **Goal**: Implement a generic stack data structure
- **Concepts**: Generics, data structures

### Challenge 2: HTTP Server (Not Implemented)
- **Difficulty**: Advanced
- **Goal**: Create a struct-based HTTP router
- **Concepts**: Real-world application, method handling

### Challenge 3: TTL Cache (Not Implemented)
- **Difficulty**: Advanced
- **Goal**: Build a cache with automatic expiration
- **Concepts**: Goroutines, time handling, concurrent access

## Tips for Success

1. **Start Simple**: Begin with Exercise 1 and work your way up
2. **Read Error Messages**: Go's compiler provides helpful error messages
3. **Use `go fmt`**: Keep your code properly formatted
4. **Experiment**: Try different approaches and see what works
5. **Check Solutions**: Don't hesitate to peek at solutions if you're stuck
6. **Practice**: The more you write Go code, the more natural it becomes

## Common Gotchas

1. **Pointer vs Value Receivers**: Use pointer receivers when you need to modify the struct or when the struct is large
2. **Interface Satisfaction**: You don't need to explicitly declare that a type implements an interface
3. **Embedding vs Composition**: Embedded fields are promoted to the outer struct
4. **Zero Values**: Structs have sensible zero values - use this to your advantage
5. **Exported vs Unexported**: Capital letters make fields/methods public

## Next Steps

After completing these exercises, consider:
- Building a complete application using structs
- Learning about goroutines and channels
- Exploring Go's standard library packages
- Reading "The Go Programming Language" book
- Contributing to open-source Go projects

Happy coding! 🚀
