# 📝 Variables in Go

## What is a Variable?

A variable is a **container that stores a value**.

Think of it as a **labeled box** where you put data.

**Example:**

Box name: age
Value stored: 25

---

## How to Declare a Variable?

Go has **3 ways** to declare variables:

```go
var name string
name = "Alice"
```

Or declare and assign together:
```go
var age int = 25
```

### **Method 2: Using `:=` shorthand (most common)**

```go
name := "Bob"
age := 30
```
This **automatically detects the type**.

### **Method 3: Multiple variables in single line**

```go
a, b, c := 10, 20, 30
name, age, city := "Alice", 25, "NewYork"
```

All variables declared and assigned in one statement.

---

## Basic Data Types in Go

| Type | What it stores | Example |
|------|---|---|
| **int** | Whole numbers | `age := 25` |
| **float64** | Decimal numbers | `price := 19.99` |
| **string** | Text | `name := "Alice"` |
| **bool** | True or False | `isActive := true` |

---

## Simple Example

```go
package main

import "fmt"

func main() {
    // Declare variables
    name := "Alice"
    age := 25
    height := 5.6
    isStudent := true

    // Print variables
    fmt.Println(name)       // Output: Alice
    fmt.Println(age)        // Output: 25
    fmt.Println(height)     // Output: 5.6
    fmt.Println(isStudent)  // Output: true
}
```

---

## Multiple Variables in Single Line

```go
package main

import "fmt"

func main() {
    // Declare multiple variables in one line
    a, b, c := 10, 20, 30
    
    fmt.Println(a)  // Output: 10
    fmt.Println(b)  // Output: 20
    fmt.Println(c)  // Output: 30

    // Different types in single line
    name, age, city := "Bob", 30, "London"
    
    fmt.Println(name)  // Output: Bob
    fmt.Println(age)   // Output: 30
    fmt.Println(city)  // Output: London
}
```

---

## Rules for Variable Names

1. **Start with letter or underscore** - `name`, `_age` ✅
2. **Can contain letters, numbers, underscores** - `var1`, `user_name` ✅
3. **Cannot start with number** - `1name` ❌
4. **Case sensitive** - `Name` and `name` are different
5. **Cannot use spaces** - `my name` ❌

---

## Variable Initialization

**Must initialize before using:**

```go
package main

import "fmt"

func main() {
    var x int      // Declares but doesn't initialize
    // x = 0 (default value for int)
    
    fmt.Println(x)  // Output: 0

    y := 10         // Declares and initializes with 10
    fmt.Println(y)  // Output: 10
}
```

**Default Values:**
- `int` → 0
- `float64` → 0.0
- `string` → "" (empty)
- `bool` → false

---

## Changing Variable Values

Once declared, you can change the value:

```go
package main

import "fmt"

func main() {
    age := 25
    fmt.Println(age)  // Output: 25
    
    age = 26
    fmt.Println(age)  // Output: 26
}
```

---

## 💡 Memory Points

1. **Variable** = Container that stores a value
2. **`var`** = Declare with explicit type
3. **`:=`** = Shorthand (auto-detect type)
4. **Multiple variables** = `a, b, c := 1, 2, 3`
5. **Types** = int, float64, string, bool
6. **Rules** = Start with letter/underscore, no spaces, case-sensitive
7. **Default values** = 0 for int, "" for string, false for bool

---

