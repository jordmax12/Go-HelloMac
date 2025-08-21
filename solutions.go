package main

import "fmt"

// import "fmt"

// import (
// 	"encoding/json"
// 	"fmt"
// 	"math"
// 	"strings"
// )

// // =============================================================================
// // SOLUTIONS - Complete implementations of all exercises
// // =============================================================================

// // EXERCISE 1: Basic Structs
// type Person struct {
// 	Name  string
// 	Age   int
// 	Email string
// }

// func NewPerson(name string, age int, email string) Person {
// 	return Person{
// 		Name:  name,
// 		Age:   age,
// 		Email: email,
// 	}
// }

// func (p Person) Greet() string {
// 	return fmt.Sprintf("Hi, I'm %s and I'm %d years old", p.Name, p.Age)
// }

// // EXERCISE 2: Structs with Methods and Pointers
// type BankAccount struct {
// 	AccountNumber string
// 	Balance       float64
// 	Owner         string
// }

// func (ba *BankAccount) Deposit(amount float64) {
// 	if amount > 0 {
// 		ba.Balance += amount
// 	}
// }

// func (ba *BankAccount) Withdraw(amount float64) bool {
// 	if amount > 0 && ba.Balance >= amount {
// 		ba.Balance -= amount
// 		return true
// 	}
// 	return false
// }

// func (ba BankAccount) GetBalance() float64 {
// 	return ba.Balance
// }

// // EXERCISE 3: Struct Embedding
// type Address struct {
// 	Street  string
// 	City    string
// 	State   string
// 	ZipCode string
// }

// type Employee struct {
// 	Person
// 	EmployeeID  string
// 	Department  string
// 	Salary      float64
// 	WorkAddress Address
// }

// func (e Employee) GetEmployeeInfo() string {
// 	return fmt.Sprintf("Employee: %s (ID: %s)\nDepartment: %s\nSalary: $%.2f\nWork Address: %s, %s, %s %s\nContact: %s",
// 		e.Name, e.EmployeeID, e.Department, e.Salary,
// 		e.WorkAddress.Street, e.WorkAddress.City, e.WorkAddress.State, e.WorkAddress.ZipCode,
// 		e.Email)
// }

// // EXERCISE 4: Interfaces
// type Shape interface {
// 	Area() float64
// 	Perimeter() float64
// }

// type Rectangle struct {
// 	Width  float64
// 	Height float64
// }

// func (r Rectangle) Area() float64 {
// 	return r.Width * r.Height
// }

// func (r Rectangle) Perimeter() float64 {
// 	return 2 * (r.Width + r.Height)
// }

// type Circle struct {
// 	Radius float64
// }

// func (c Circle) Area() float64 {
// 	return math.Pi * c.Radius * c.Radius
// }

// func (c Circle) Perimeter() float64 {
// 	return 2 * math.Pi * c.Radius
// }

// func PrintShapeInfo(s Shape) {
// 	fmt.Printf("Shape - Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
// }

// // EXERCISE 5: Advanced Structs with Slices
// type Book struct {
// 	Title     string
// 	Author    string
// 	ISBN      string
// 	Year      int
// 	Available bool
// }

// type Library struct {
// 	Name  string
// 	Books []Book
// }

// func (l *Library) AddBook(book Book) {
// 	l.Books = append(l.Books, book)
// }

// func (l Library) FindBookByTitle(title string) (*Book, bool) {
// 	for i := range l.Books {
// 		if strings.EqualFold(l.Books[i].Title, title) {
// 			return &l.Books[i], true
// 		}
// 	}
// 	return nil, false
// }

// func (l *Library) CheckoutBook(title string) bool {
// 	book, found := l.FindBookByTitle(title)
// 	if found && book.Available {
// 		book.Available = false
// 		return true
// 	}
// 	return false
// }

// // EXERCISE 6: JSON Tags
// type Product struct {
// 	ID          int     `json:"id"`
// 	Name        string  `json:"name"`
// 	Price       float64 `json:"price"`
// 	Description string  `json:"description"`
// 	InStock     bool    `json:"in_stock"`
// }

// // EXERCISE 7: Custom Types
// type Temperature float64

// func (t Temperature) ToFahrenheit() Temperature {
// 	return Temperature(float64(t)*9/5 + 32)
// }

// func (t Temperature) ToCelsius() Temperature {
// 	return Temperature((float64(t) - 32) * 5 / 9)
// }

// func (t Temperature) String() string {
// 	return fmt.Sprintf("%.2f°", float64(t))
// }

// // =============================================================================
// // CHALLENGE SOLUTIONS
// // =============================================================================

// // CHALLENGE 1: Generic Stack
// type Stack[T any] struct {
// 	items []T
// }

// func (s *Stack[T]) Push(item T) {
// 	s.items = append(s.items, item)
// }

// func (s *Stack[T]) Pop() (T, bool) {
// 	var zero T
// 	if len(s.items) == 0 {
// 		return zero, false
// 	}
// 	item := s.items[len(s.items)-1]
// 	s.items = s.items[:len(s.items)-1]
// 	return item, true
// }

// func (s *Stack[T]) Peek() (T, bool) {
// 	var zero T
// 	if len(s.items) == 0 {
// 		return zero, false
// 	}
// 	return s.items[len(s.items)-1], true
// }

// func (s *Stack[T]) IsEmpty() bool {
// 	return len(s.items) == 0
// }

// func (s *Stack[T]) Size() int {
// 	return len(s.items)
// }

// // =============================================================================
// // DEMONSTRATION FUNCTIONS
// // =============================================================================

// func demonstratePersonStruct() {
// 	fmt.Println("=== Person Struct Demo ===")
// 	person := NewPerson("Alice", 30, "alice@example.com")
// 	fmt.Println(person.Greet())
// 	fmt.Printf("Person details: %+v\n\n", person)
// }

// func demonstrateBankAccount() {
// 	fmt.Println("=== Bank Account Demo ===")
// 	account := BankAccount{
// 		AccountNumber: "12345",
// 		Balance:       1000.0,
// 		Owner:         "John Doe",
// 	}

// 	fmt.Printf("Initial balance: $%.2f\n", account.GetBalance())
// 	account.Deposit(500.0)
// 	fmt.Printf("After deposit: $%.2f\n", account.GetBalance())

// 	success := account.Withdraw(200.0)
// 	fmt.Printf("Withdrawal successful: %t, Balance: $%.2f\n", success, account.GetBalance())

// 	success = account.Withdraw(2000.0)
// 	fmt.Printf("Large withdrawal successful: %t, Balance: $%.2f\n\n", success, account.GetBalance())
// }

// func demonstrateEmployeeStruct() {
// 	fmt.Println("=== Employee Struct Demo ===")
// 	emp := Employee{
// 		Person: Person{
// 			Name:  "Jane Smith",
// 			Age:   28,
// 			Email: "jane@company.com",
// 		},
// 		EmployeeID: "EMP001",
// 		Department: "Engineering",
// 		Salary:     75000.0,
// 		WorkAddress: Address{
// 			Street:  "123 Tech St",
// 			City:    "San Francisco",
// 			State:   "CA",
// 			ZipCode: "94105",
// 		},
// 	}

// 	fmt.Println(emp.GetEmployeeInfo())
// 	fmt.Println("Embedded method:", emp.Greet())
// 	fmt.Println()
// }

// func demonstrateShapeInterface() {
// 	fmt.Println("=== Shape Interface Demo ===")
// 	rect := Rectangle{Width: 10, Height: 5}
// 	circle := Circle{Radius: 7}

// 	fmt.Print("Rectangle: ")
// 	PrintShapeInfo(rect)
// 	fmt.Print("Circle: ")
// 	PrintShapeInfo(circle)
// 	fmt.Println()
// }

// func demonstrateLibrary() {
// 	fmt.Println("=== Library Demo ===")
// 	lib := Library{Name: "City Library"}

// 	book1 := Book{
// 		Title:     "The Go Programming Language",
// 		Author:    "Alan Donovan",
// 		ISBN:      "978-0134190440",
// 		Year:      2015,
// 		Available: true,
// 	}

// 	book2 := Book{
// 		Title:     "Clean Code",
// 		Author:    "Robert Martin",
// 		ISBN:      "978-0132350884",
// 		Year:      2008,
// 		Available: true,
// 	}

// 	lib.AddBook(book1)
// 	lib.AddBook(book2)

// 	if book, found := lib.FindBookByTitle("Clean Code"); found {
// 		fmt.Printf("Found book: %s by %s\n", book.Title, book.Author)
// 	}

// 	success := lib.CheckoutBook("Clean Code")
// 	fmt.Printf("Checkout successful: %t\n", success)

// 	success = lib.CheckoutBook("Clean Code")
// 	fmt.Printf("Second checkout successful: %t\n\n", success)
// }

// func demonstrateTemperature() {
// 	fmt.Println("=== Temperature Demo ===")
// 	celsius := Temperature(25.0)
// 	fmt.Printf("25°C = %s\n", celsius.ToFahrenheit())

// 	fahrenheit := Temperature(77.0)
// 	fmt.Printf("77°F = %s\n\n", fahrenheit.ToCelsius())
// }

// func demonstrateJSON() {
// 	fmt.Println("=== JSON Demo ===")
// 	product := Product{
// 		ID:          1,
// 		Name:        "Laptop",
// 		Price:       999.99,
// 		Description: "High-performance laptop",
// 		InStock:     true,
// 	}

// 	jsonData, err := json.MarshalIndent(product, "", "  ")
// 	if err != nil {
// 		fmt.Printf("Error marshaling JSON: %v\n", err)
// 		return
// 	}

// 	fmt.Printf("Product as JSON:\n%s\n\n", string(jsonData))
// }

// func demonstrateStack() {
// 	fmt.Println("=== Generic Stack Demo ===")
// 	stack := Stack[int]{}

// 	stack.Push(10)
// 	stack.Push(20)
// 	stack.Push(30)

// 	fmt.Printf("Stack size: %d\n", stack.Size())

// 	if item, ok := stack.Peek(); ok {
// 		fmt.Printf("Top item: %d\n", item)
// 	}

// 	for !stack.IsEmpty() {
// 		if item, ok := stack.Pop(); ok {
// 			fmt.Printf("Popped: %d\n", item)
// 		}
// 	}
// 	fmt.Println()
// }

// func runSolutions() {
// 	fmt.Println("Go Structs and Concepts - SOLUTIONS")
// 	fmt.Println("===================================")
// 	fmt.Println()

// 	demonstratePersonStruct()
// 	demonstrateBankAccount()
// 	demonstrateEmployeeStruct()
// 	demonstrateShapeInterface()
// 	demonstrateLibrary()
// 	demonstrateTemperature()
// 	demonstrateJSON()
// 	demonstrateStack()

// 	fmt.Println("All demonstrations completed!")
// }

func test() {
	fmt.Println("=== Testing Person Struct ===")
}
