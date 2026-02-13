# How Go Uses Implicit Interfaces

In Go, types automatically satisfy interfaces if they implement the required methods—no explicit declaration needed. This means you can define interfaces at the point of consumption rather than at the point of implementation:

In many languages like Java or C#, you must explicitly declare that a type implements an interface. In Go, there's no such declaration. If a type has the methods an interface requires, it automatically satisfies that interface:

```go
package main

import "fmt"

// Move is an interface — any type with a Use() method satisfies it
type Move interface {
    Use()
}

// Thunderbolt automatically implements Move
type Thunderbolt struct{}

func (t Thunderbolt) Use() {
    fmt.Println("⚡ Pikachu used Thunderbolt!")
}

// QuickAttack automatically implements Move
type QuickAttack struct{}

func (q QuickAttack) Use() {
    fmt.Println("💨 Eevee used Quick Attack!")
}

// performMove accepts anything that satisfies the Move interface
func performMove(m Move) {
    m.Use()
}

func main() {
    t := Thunderbolt{}
    q := QuickAttack{}

    performMove(t) // ⚡ Pikachu used Thunderbolt!
    performMove(q) // 💨 Eevee used Quick Attack!
}
```

## Zero Declaration Overhead
Thunderbolt automatically satisfies `Move` just by having a `Use()` method. No boilerplate, no explicit declarations—it just works.

```go
type Thunderbolt struct{}
func (t Thunderbolt) Use() {
    fmt.Println("⚡ Pikachu used Thunderbolt!")
}
```
## Interface Defined at Point of Need
The interface serves the consumer's needs. If `performMove` only needs `Use()`, that's all the interface requires—even if moves could have other methods like `GetPower()` or `GetType()`.

```go
// The Move interface is defined where it's consumed (in performMove)
// not where it's implemented (Thunderbolt/QuickAttack)
func performMove(m Move) {
    m.Use()
}
```

## Effortless Extensibility

You can add new move types from anywhere in your codebase (or even external packages) and they automatically work with existing functions.

```go
// Add new move types without touching existing code
type Flamethrower struct{}
func (f Flamethrower) Use() {
    fmt.Println("🔥 Charizard used Flamethrower!")
}

type Splash struct{}
func (s Splash) Use() {
    fmt.Println("💧 Magikarp used Splash! But nothing happened...")
}

func main() {
    // These work immediately with performMove - no code changes needed!
    performMove(Flamethrower{})
    performMove(Splash{})
}
```
## Key Points

`Thunderbolt` and `QuickAttack` have no idea they implement `Move`—they just happen to have a `Use()` method. This makes the code:

- Flexible: Easy to add new moves
- Testable: Simple to mock
- Maintainable: Low coupling between components
- Natural: Follows how we think about behavior

This is why Go's implicit interfaces make dependency injection feel effortless—it's just good design falling into place naturally.