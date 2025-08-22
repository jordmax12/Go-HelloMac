package main

import (
	"fmt"
	"math"
)

// =============================================================================
// EXERCISE 1: Basic Structs
// =============================================================================

type Person struct {
	Name  string
	Age   int
	Email string
}

func NewPerson(name string, age int, email string) Person {
	return Person{
		Name:  name,
		Age:   age,
		Email: email,
	}
}

func (p Person) Greet() string {
	return fmt.Sprintf("Hi, I'm %s and I'm %d years old", p.Name, p.Age)
}

// =============================================================================
// EXERCISE 2: Structs with Methods and Pointers
// =============================================================================

// TODO: Define a BankAccount struct with fields: AccountNumber (string), Balance (float64), Owner (string)
type BankAccount struct {
	AccountNumber string
	Balance       float64
	Owner         string
}

// TODO: Create a method to deposit money (should modify the balance)
func (ba *BankAccount) Deposit(amount float64) bool {
	if 0 >= amount {
		return false
	}

	ba.Balance += amount
	return true
}

// TODO: Create a method to withdraw money (should check if sufficient funds exist)
func (ba *BankAccount) Withdraw(amount float64) bool {
	fmt.Printf("Withdrawing %.2f from balance %.2f\n", amount, ba.Balance)
	if amount <= 0 || amount > ba.Balance {
		return false
	}

	ba.Balance -= amount
	return true
}

// TODO: Create a method to get the current balance
func (ba BankAccount) GetBalance() float64 {
	return ba.Balance
}

// =============================================================================
// EXERCISE 3: Struct Embedding (Composition)
// =============================================================================

// TODO: Define an Address struct
type Address struct {
	Street  string
	City    string
	State   string
	ZipCode string
}

// TODO: Define an Employee struct that embeds Person and has additional fields
type Employee struct {
	Person      // ✅ Embedded - no field name!
	EmployeeID  string
	Department  string
	Salary      float64
	WorkAddress Address
}

// TODO: Create a method for Employee that returns their full info
func (e Employee) GetEmployeeInfo() string {
	return fmt.Sprintf("Employee Name: %s, Employee Age: %d, Employee Email: %s, Employee ID: %s, Department: %s, Salary: %.2f, Address: %s, %s, %s, %s",
		e.Name,
		e.Age,
		e.Email,
		e.EmployeeID,
		e.Department,
		e.Salary,
		e.WorkAddress.Street,
		e.WorkAddress.City,
		e.WorkAddress.State,
		e.WorkAddress.ZipCode,
	)
}

// =============================================================================
// EXERCISE 4: Interfaces
// =============================================================================

// TODO: Define a Shape interface with methods Area() and Perimeter()
type Shape interface {
	Area() float64
	Perimeter() float64
}

// TODO: Define a Rectangle struct and implement the Shape interface
type Rectangle struct {
	Width  float64
	Height float64
}

// TODO: Implement Area method for Rectangle
func (r Rectangle) Area() float64 {
	fmt.Println("Area of Rectangle")
	return r.Width * r.Height
}

// TODO: Implement Perimeter method for Rectangle
func (r Rectangle) Perimeter() float64 {
	fmt.Println("Perimeter of Rectangle")
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Test() float64 {
	fmt.Println("Test of Rectangle")
	return 2 * (r.Width + r.Height)
}

// TODO: Define a Circle struct and implement the Shape interface
type Circle struct {
	Radius float64
}

// TODO: Implement Area method for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// TODO: Implement Perimeter method for Circle
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// TODO: Create a function that takes a Shape interface and prints its area and perimeter
func PrintShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

// =============================================================================
// EXERCISE 5: Advanced Structs with Slices and Maps
// =============================================================================

// TODO: Define a Library struct that manages books
// type Book struct {
// 	Title    string
// 	Author   string
// 	ISBN     string
// 	Year     int
// 	Available bool
// }

// type Library struct {
// 	Name  string
// 	Books []Book
// 	// You can also use a map for faster lookups: Books map[string]Book
// }

// TODO: Create a method to add a book to the library
// func (l *Library) AddBook(book Book) {
// 	// Your code here
// }

// TODO: Create a method to find a book by title
// func (l Library) FindBookByTitle(title string) (*Book, bool) {
// 	// Your code here - return pointer to book and whether it was found
// }

// TODO: Create a method to check out a book (mark as unavailable)
// func (l *Library) CheckoutBook(title string) bool {
// 	// Your code here - return true if successful, false if book not found or already checked out
// }

// =============================================================================
// EXERCISE 6: Struct with JSON Tags (Bonus)
// =============================================================================

// TODO: Define a Product struct with JSON tags for API responses
// type Product struct {
// 	ID          int     `json:"id"`
// 	Name        string  `json:"name"`
// 	Price       float64 `json:"price"`
// 	Description string  `json:"description"`
// 	InStock     bool    `json:"in_stock"`
// }

// =============================================================================
// EXERCISE 7: Custom Types and Methods
// =============================================================================

// TODO: Define a custom type Temperature based on float64
// type Temperature float64

// TODO: Create methods to convert between Celsius and Fahrenheit
// func (t Temperature) ToFahrenheit() Temperature {
// 	// Your code here - assume t is in Celsius
// }

// func (t Temperature) ToCelsius() Temperature {
// 	// Your code here - assume t is in Fahrenheit
// }

// func (t Temperature) String() string {
// 	// Your code here - return a formatted string representation
// }

// =============================================================================
// TEST FUNCTIONS - Uncomment and run these to test your implementations
// =============================================================================

func testPersonStruct() {
	fmt.Println("=== Testing Person Struct ===")
	// TODO: Uncomment and test your Person implementation
	person := NewPerson("Alice", 30, "alice@example.com")
	fmt.Println(person.Greet())
}

func testBankAccount() {
	fmt.Println("\n\n=== Testing Bank Account ===")

	account := BankAccount{
		AccountNumber: "12345",
		Balance:       1000.0,
		Owner:         "John Doe",
	}

	fmt.Printf("Initial balance: $%.2f\n", account.GetBalance())
	account.Deposit(500.0)
	fmt.Printf("After deposit: $%.2f\n", account.GetBalance())

	success := account.Withdraw(200.0)
	fmt.Printf("Withdrawal successful: %t, Balance: $%.2f\n", success, account.GetBalance())

	success = account.Withdraw(2000.0)
	fmt.Printf("Large withdrawal successful: %t, Balance: $%.2f\n", success, account.GetBalance())
}

func testEmployeeStruct() {
	fmt.Println("\n\n=== Testing Employee Struct ===")
	// TODO: Uncomment and test your Employee implementation
	emp := Employee{
		Person: Person{
			Name:  "Jane Smith",
			Age:   28,
			Email: "jane@company.com",
		},
		EmployeeID: "EMP001",
		Department: "Engineering",
		Salary:     75000.0,
		WorkAddress: Address{
			Street:  "123 Tech St",
			City:    "San Francisco",
			State:   "CA",
			ZipCode: "94105",
		},
	}
	//
	fmt.Println(emp.GetEmployeeInfo())
	fmt.Println(emp.Greet()) // This should work due to embedding!
}

func testShapeInterface() {
	fmt.Println("\n\n=== Testing Shape Interface ===")
	// TODO: Uncomment and test your Shape implementations
	rect := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 7}

	PrintShapeInfo(rect)
	PrintShapeInfo(circle)
}

func testLibrary() {
	fmt.Println("=== Testing Library ===")
	// TODO: Uncomment and test your Library implementation
	// lib := Library{Name: "City Library"}
	//
	// book1 := Book{
	// 	Title:     "The Go Programming Language",
	// 	Author:    "Alan Donovan",
	// 	ISBN:      "978-0134190440",
	// 	Year:      2015,
	// 	Available: true,
	// }
	//
	// book2 := Book{
	// 	Title:     "Clean Code",
	// 	Author:    "Robert Martin",
	// 	ISBN:      "978-0132350884",
	// 	Year:      2008,
	// 	Available: true,
	// }
	//
	// lib.AddBook(book1)
	// lib.AddBook(book2)
	//
	// if book, found := lib.FindBookByTitle("Clean Code"); found {
	// 	fmt.Printf("Found book: %s by %s\n", book.Title, book.Author)
	// }
	//
	// success := lib.CheckoutBook("Clean Code")
	// fmt.Printf("Checkout successful: %t\n", success)
	//
	// success = lib.CheckoutBook("Clean Code")
	// fmt.Printf("Second checkout successful: %t\n", success)
}

func testTemperature() {
	fmt.Println("=== Testing Temperature ===")
	// TODO: Uncomment and test your Temperature implementation
	// celsius := Temperature(25.0)
	// fmt.Printf("25°C = %s\n", celsius.ToFahrenheit())
	//
	// fahrenheit := Temperature(77.0)
	// fmt.Printf("77°F = %s\n", fahrenheit.ToCelsius())
}

// =============================================================================
// CHALLENGE EXERCISES
// =============================================================================

// CHALLENGE 1: Create a generic Stack data structure using structs
// Implement Push, Pop, Peek, IsEmpty, and Size methods

// CHALLENGE 2: Create a simple HTTP server struct that can register routes
// and handle different HTTP methods

// CHALLENGE 3: Implement a simple cache with TTL (Time To Live) using structs
// Include methods for Set, Get, Delete, and automatic cleanup of expired items

func runExercises() {
	fmt.Println("Go Structs and Concepts Exercises")
	fmt.Println("=================================")
	fmt.Println()
	fmt.Println("Instructions:")
	fmt.Println("1. Implement the TODOs above")
	fmt.Println("2. Uncomment the test functions below to test your implementations")
	fmt.Println("3. Call runExercises() from main.go or create a separate main")
	fmt.Println()

	// Uncomment these as you complete each exercise
	testPersonStruct()
	testBankAccount()
	testEmployeeStruct()
	testShapeInterface()
	// testLibrary()
	// testTemperature()

	fmt.Println("\n\nComplete the exercises above and uncomment the test functions!")
}
