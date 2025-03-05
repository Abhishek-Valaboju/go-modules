# Object Oriented Programming (OOPs) Concept in Go
In Go (Golang), OOP (Object-Oriented Programming) is not implemented in the classical sense, as it lacks some traditional OOP features like classes and inheritance. However, Go still provides several concepts that allow you to apply object-oriented principles, such as encapsulation, polymorphism, and composition. Here's how OOP concepts map to Go:

### 1. **Encapsulation**
Go achieves encapsulation through package-level visibility. Identifiers (like variables, types, and functions) that start with an uppercase letter are exported and can be accessed from outside the package. Identifiers starting with a lowercase letter are unexported and private to the package.

Example:
```go
package mypackage

type Person struct {
    Name string  // Exported, accessible from outside the package
    age  int     // Unexported, only accessible within the package
}
```

### 2. **Structs and Methods**
Instead of classes, Go uses **structs** to define types. You can associate methods with structs to create behavior similar to that of a class.

Example:
```go
type Person struct {
    Name string
    Age  int
}

// Method associated with the Person struct
func (p Person) Greet() string {
    return "Hello, my name is " + p.Name
}
```

### 3. **Composition (instead of Inheritance)**
Go doesn’t support inheritance, but it emphasizes **composition** over inheritance. You can compose types by embedding structs within other structs, achieving a form of code reuse.

Example:
```go
type Employee struct {
    // Person from earlier examples
    Person          // Embeded struct
    JobTitle string
}

// Employee now has access to Person's fields and methods
func main() {
    e := Employee{
        Person: Person{
            Name: "Alice",
            Age: 30,
        },
        JobTitle: "Engineer",
    }
    fmt.Println(e.Greet()) // Output: Hello, my name is Alice
}
```

### 4. **Abstraction and Polymorphism using Interfaces**
Go uses **interfaces** to achieve polymorphism. An interface in Go is a set of method signatures. Any type that implements these methods satisfies the interface, enabling polymorphic behavior without explicit declarations.

Example:
```go
// Greeter hides the implementation of the Greet method by any type (Person or
// Dog), exposing only essential features, which demonstrates abstraction.
type Greeter interface {
    Greet() string
}

func SayHello(g Greeter) {
    fmt.Println(g.Greet())
}

type Dog struct {
    Name string
}

func (d Dog) Greet() string {
    return "Woof! My name is " + d.Name
}

func main() {
    // Person from earlier examples.
    // 
    // Both Person and Dog implement the Greeter interface, demonstrating polymorphism.
    p := Person{
        Name: "Bob",
        Age: 25,
    }
    d := Dog{
        Name: "Fido",
    }

    SayHello(p)  // Output: Hello, my name is Bob
    SayHello(d)  // Output: Woof! My name is Fido
}
```

### 5. **No Constructors/Destructors**
Go doesn’t have explicit constructors or destructors. However, you can define factory functions to initialize structs.

Example:
```go
func NewPerson(name string, age int) *Person {
    return &Person{Name: name, Age: age}
}
```

### 6. **Method Receivers: Value vs. Pointer**
Methods in Go can have either **value** or **pointer receivers**, which influences how the method operates on the struct.

- **Value receiver**: The method works with a copy of the struct.
- **Pointer receiver**: The method works with the actual struct, allowing you to modify it.

Example:
```go
func (p Person) GetAge() {
    return p.Age
}

func (p *Person) SetAge(age int) {
    p.Age = age // Modifies the actual struct
}
```

### 7. **Multiple Implementations via Interfaces**
A type can implement multiple interfaces by defining methods that match those interfaces.

Example:
```go
type Walker interface {
    Walk()
}

type Talker interface {
    Talk()
}

type WalkerTalker interface {
    Walker
    Talker
}

type Human struct {
    Name string
}

func (h Human) Walk() {
    fmt.Println(h.Name + " is walking")
}

func (h Human) Talk() {
    fmt.Println(h.Name + " is talking")
}

func doWalkTalk(human WalkerTalker) {
    h.Walk()
    h.Talk()
}

func main() {
    h := Human{
        Name: "John",
    }
    doWalkTalk(h)
}
```

### 8. **Lack of Inheritance**
Instead of relying on inheritance, Go encourages you to focus on interfaces and composition to share functionality between types. This keeps the code simple and avoids the complexity often associated with deep inheritance hierarchies.

In summary, while Go doesn’t implement classical OOP with inheritance and classes, it allows you to achieve object-oriented design principles through encapsulation, polymorphism via interfaces, and composition over inheritance. This promotes simpler and more maintainable code.

## Assignment
### Assignment-1: Vehicles
#### Task
- Create models for different `Vehicles` types: `Cars`, `Bikes`, and `Trucks`.
- Demonstrate OOPs concepts in it:
    - `Encapsulation`
        - Create a `Vehicle` struct that encapsulates common attributes of all vehicles, such as `Make`, `Model`, and `Year`.
        - Provide getter and setter methods to access and modify these fields.
    - `Abstraction`
        - Define a `VehicleActions` interface that includes common behaviors for all vehicles, such as `Start()`, `Stop()`, and `Drive()`.
    - `Composition`
        - Create separate structs for `Car`, `Bike`, and `Truck`, embedding the `Vehicle` struct into them to reuse common attributes.
        - Each type of vehicle should have a specific attribute of its own. For example, `Car` could have `NumberOfDoors`, `Truck` could have `PayloadCapacity`, and `Bike` could have `HasSidecar`.
    - `Polymorphism`
        - Implement the methods from the `VehicleActions` interface for each vehicle type. Each vehicle type should have its own implementation of how it `Starts`, `Stops`, and `Drives`.
        - Write a function that accepts a `VehicleActions` interface and demonstrates polymorphism by starting, stopping, and driving different vehicles.
- Additional inputs:
    - Ensure that the fields in the `Vehicle` struct are encapsulated (unexported), and provide appropriate getter and setter methods for them.
    - Use interfaces to abstract the behaviors (`Start`, `Stop`, `Drive`) that are shared across different vehicle types.
    - Demonstrate polymorphism by creating instances of `Car`, `Bike`, and `Truck`, and passing them to a function that handles them polymorphically.
    - Write a `main()` function to test your system by creating instances of `Car`, `Bike`, and `Truck`, and calling the polymorphic functions to perform actions on each vehicle.

#### Output
Output should be in this format:
```yaml
Starting Car: Toyota Corolla (2022)
Driving Car: Toyota Corolla (2022)
Stopping Car: Toyota Corolla (2022)

Starting Bike: Harley Davidson Street 750 (2021)
Driving Bike: Harley Davidson Street 750 (2021)
Stopping Bike: Harley Davidson Street 750 (2021)

Starting Truck: Ford F-150 (2023)
Driving Truck: Ford F-150 (2023)
Stopping Truck: Ford F-150 (2023)
```

#### Solution
Please use it only if you are unable to solve the problem:

https://go.dev/play/p/LLlC3xT30MJ
