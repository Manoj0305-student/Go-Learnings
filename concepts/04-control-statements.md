# 🎛️ Control Statements in Go

## What are Control Statements?
Control statements let you **make decisions** and **repeat code** based on conditions.

Go has three main control statements:
1. If/Else
2. Switch/Case
3. For Loop

---

## If/Else Statement
Execute code **if a condition is true**, otherwise execute different code.

**Syntax:**
```go
if condition {
    // Do something if true
} else if anotherCondition {
    // Do something else
} else {
    // Do this if all above are false
}
```

**Example:**
```go
package main

import "fmt"

func main() {
    age := 18
    if age >= 18 {
        fmt.Println("You are an adult")
    } else if age >= 13 {
        fmt.Println("You are a teenager")
    } else {
        fmt.Println("You are a child")
    }
}
```

**Output:**
You are an adult

**Common Comparisons:**
- `==` Equal to
- `!=` Not equal to
- `>` Greater than
- `<` Less than
- `>=` Greater than or equal
- `<=` Less than or equal
- `&&` AND (both true)
- `||` OR (one or both true)

---

## Switch/Case Statement

Choose **one option from many** based on a value.

**Syntax:**
```go
switch value {
case option1:
    // Do something
case option2:
    // Do something else
case option3, option4:
    // Multiple cases with same action
default:
    // Do this if no match
}
```

**Example:**
```go
package main

import "fmt"

func main() {
    day := "Monday"
    switch day {
    case "Monday":
        fmt.Println("Start of week")
    case "Friday":
        fmt.Println("Almost weekend")
    case "Saturday", "Sunday":
        fmt.Println("Weekend!")
    default:
        fmt.Println("Some other day")
    }
}
```

**Output:**
Start of week

**Key Points:**
- No `break` needed (Go adds it automatically)
- Multiple cases can use same action (case "Saturday", "Sunday")
- `default` is optional

---

## For Loop

Repeat code **multiple times** in different ways.

### **Type 1: Basic For Loop**

```go
package main

import "fmt"

func main() {
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
}
```

**Output:**
0
1
2
3
4 (Every Output will be in new line)

**Syntax Breakdown:**
- `i := 0` - Initialize variable
- `i < 5` - Condition (repeat while true)
- `i++` - Increment after each iteration

---

### **Type 2: For Loop with Condition Only**

```go
package main

import "fmt"

func main() {
    i := 0
    for i < 5 {
        fmt.Println(i)
        i++
    }
}
```

**Output:**
0
1
2
3
4

---

## Range Function

The `range` function is used to **iterate over collections** like arrays, slices, and maps.

**What does range do?**
- Loops through each element in a collection
- Returns the **index** and **value** of each element
- Works with arrays, slices, maps, and strings

**Syntax:**
```go
for index, value := range collection {
    // Use index and value
}
```

**What is returned:**
- **index** - Position of element (starts from 0)
- **value** - The actual element/data
- You can use `_` if you only need one

---

### **Type 3: For Range (Iterate over Collections)**

```go
package main

import "fmt"

func main() {
    numbers := []int{10, 20, 30}
    for index, value := range numbers {
        fmt.Println(index, value)
    }
}
```

**Output:**
0 10
1 20
2 30

---

## Break and Continue

Control the flow inside loops.

### **Break: Stop Loop Immediately**

```go
package main

import "fmt"

func main() {
    for i := 0; i < 10; i++ {
        if i == 5 {
            break  // Stop loop
        }
        fmt.Println(i)
    }
}
```

**Output:**
0
1
2
3
4

(Loop stops when i is 5)

---

### **Continue: Skip Current Iteration**

```go
package main

import "fmt"

func main() {
    for i := 0; i < 5; i++ {
        if i == 2 {
            continue  // Skip this iteration
        }
        fmt.Println(i)
    }
}
```

**Output:**
0
1
3
4

(Number 2 is skipped)

---

## ⚠️ Important Rule: Curly Brace Placement

**In Go, the opening curly brace `{` MUST be on the SAME LINE as the statement.**

### **❌ WRONG (Curly brace on new line):**

```go
func PrintNumber(num int)
{
    fmt.Println(num)
}
```

This will cause a **compilation error**.

### **✅ CORRECT (Curly brace on same line):**

```go
func PrintNumber(num int) {
    fmt.Println(num)
}
```

This is the **correct Go style**.

### **Same Rule Applies to:**

**If statements:**
```go
if age >= 18 {  // ✅ Correct
    fmt.Println("Adult")
}

// ❌ Wrong:
if age >= 18
{
    fmt.Println("Adult")
}
```

**Switch statements:**
```go
switch day {  // ✅ Correct
case "Monday":
    fmt.Println("Monday")
}

// ❌ Wrong:
switch day
{
    case "Monday":
        fmt.Println("Monday")
}
```

**For loops:**
```go
for i := 0; i < 5; i++ {  // ✅ Correct
    fmt.Println(i)
}

// ❌ Wrong:
for i := 0; i < 5; i++
{
    fmt.Println(i)
}
```

**Why?** Go's automatic semicolon insertion rule: If a line ends with a keyword like `func`, `if`, `for`, or `switch`, Go inserts a semicolon, which breaks the code if the brace is on the next line.

---

## Complete Example: Using All Control Statements

```go
package main

import "fmt"

func main() {
    // If/Else
    score := 85
    if score >= 90 {
        fmt.Println("Grade: A")
    } else if score >= 80 {
        fmt.Println("Grade: B")
    } else {
        fmt.Println("Grade: C")
    }

    // Switch/Case
    day := "Monday"
    switch day {
    case "Monday":
        fmt.Println("Start of week")
    case "Friday":
        fmt.Println("Almost weekend")
    default:
        fmt.Println("Some other day")
    }

    // For Loop
    fmt.Println("Numbers:")
    for i := 1; i <= 5; i++ {
        fmt.Print(i," ")
    }

    // Break and Continue
    fmt.Println("Loop with break and continue:")
    for i := 1; i <= 10; i++ {
        if i == 5 {
            continue
        }
        if i == 8 {
            break
        }
        fmt.Print(i," ")
    }
}
```

**Output:**

Grade: B

Start of week

Numbers:
1
2
3
4
5

Loop with break and continue:
1
2
3
4
6
7

---

## 💡 Memory Points

1. **If/Else** = Make decisions based on conditions
2. **Switch/Case** = Choose one option from many
3. **For Loop** = Repeat code multiple times
4. **Break** = Stop loop immediately
5. **Continue** = Skip current iteration, go to next
6. **Range** = Loop through collections and get index + value
7. **Curly brace rule** = Opening `{` MUST be on same line as statement
8. **Default case** = Executes if no case matches in switch

---

## Practice Questions & Solutions

Below are 5 practice questions to test your understanding of control statements.

**How to use this section:**
1. **Try solving each question yourself first**
2. **If you get stuck**, check the solution in `daily-progress/control-statements/` folder
3. **Compare your code** with the solution and understand the differences

### **Question 1: If/Else (Easy)**
Check if a number is **positive, negative, or zero** and print the result.

### **Question 2: Switch/Case (Easy)**
Take a **grade** (A, B, C, D, F) and print its **description** (Excellent, Good, Average, Poor, Fail).

### **Question 3: For Loop (Easy)**
Print **numbers from 1 to N** where N is a variable.

### **Question 4: Break and Continue (Medium)**
Print numbers 1 to 10, but **skip 5** (use continue) and **stop at 8** (use break).

### **Question 5: For Loop with Sum (Medium)**
Calculate the **sum of numbers from 1 to N**.

**Example:** N = 5 → Sum = 1+2+3+4+5 = 15

---

**Solutions:** Find all solutions in `daily-progress/control-statements/` folder. Each question has a corresponding `.go` file with the complete solution.

---
