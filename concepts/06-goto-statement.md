# 🔀 Goto Statement in Go

## What is `goto` Statement?

`goto` is a statement that **jumps to a labeled location** in your code.

Think of it as a **teleporter** - it takes you directly to another part of the code.

**Syntax:**
```go
goto labelname

// Later in code...
labelname:
    // Code here
```

---

## 🔴 Marvel Analogy: Azazel from X-Men

Imagine **Azazel from X-Men: First Class**.

Azazel has the power to **teleport** - he disappears in a flash of red smoke from one location and **instantly reappears** somewhere else, even grabbing enemies and bringing them along!

**That's exactly what `goto` does:**
- Azazel is at location A → He disappears
- `goto` jumps from current location → to labeled location
- Azazel reappears at location B → Code execution continues there
- Everything in between is skipped (just like the red smoke!)

**The catch:** Just like Azazel's power is cool but dangerous if overused, `goto` should be used very carefully and rarely!

---

## Simple Example

```go
package main

import "fmt"

func main() {
    x := 5
    
    if x > 0 {
        goto positive  // Azazel disappears and reappears at "positive" label
    }
    
    fmt.Println("This will be skipped")  // Never executed (in the smoke!)
    
    positive:
        fmt.Println("Number is positive!")  // Azazel reappears here!
}
```

**Output:**
Number is positive!

## ❌ When NOT to Use `goto`

**Most of the time!**

`goto` is generally **considered bad practice** because:

1. **Makes code hard to follow** - Jumps around make logic confusing (too much teleporting!)
2. **Creates "spaghetti code"** - Hard to trace execution flow (where did Azazel go?)
3. **Difficult to debug** - Hard to understand where code goes
4. **Can be replaced** - Use `if/else`, `break`, `continue` instead

**Using `goto` everywhere = Azazel teleporting randomly = Chaos!** 🌪️

---

## 💡 Memory Points

1. **`goto`** = Jump to labeled location (like Azazel's teleportation)
2. **Syntax** = `goto labelname` and `labelname:`
3. **Use cases** = Error handling, breaking nested loops
4. **Alternatives** = Labeled `break`, `return`, `defer`, functions
5. **Best practice** = Avoid `goto` - use structured control flow instead
6. **Readability** = `goto` makes code hard to follow (teleportation chaos!)
7. **Go philosophy** = Prefer explicit control flow (organized team strategy)
8. **Remember** = Azazel's power is cool, but using `goto` everywhere = trouble!
9. **Suggestion** Even Azazel doesn't use teleportation for everything - sometimes walking is better! 😄