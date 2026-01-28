// Package interfaces demonstrates implicit interface satisfaction in Go.
// Types automatically satisfy interfaces if they implement the required methods.
package interfaces

import "fmt"

// Move is an interface - any type with a Use() method satisfies it.
type Move interface {
	Use()
}

// Thunderbolt automatically implements Move.
type Thunderbolt struct{}

func (t Thunderbolt) Use() {
	fmt.Println("Pikachu used Thunderbolt!")
}

// QuickAttack automatically implements Move.
type QuickAttack struct{}

func (q QuickAttack) Use() {
	fmt.Println("Eevee used Quick Attack!")
}

// Flamethrower automatically implements Move.
type Flamethrower struct{}

func (f Flamethrower) Use() {
	fmt.Println("Charizard used Flamethrower!")
}

// PerformMove accepts anything that satisfies the Move interface.
func PerformMove(m Move) {
	m.Use()
}

// Run executes all interface examples.
func Run() {
	fmt.Println("=== Implicit Interfaces ===")
	fmt.Println()

	t := Thunderbolt{}
	q := QuickAttack{}
	f := Flamethrower{}

	PerformMove(t)
	PerformMove(q)
	PerformMove(f)
}