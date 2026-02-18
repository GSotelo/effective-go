package interfaces

func ExamplePerformMove() {
	PerformMove(Thunderbolt{})
	PerformMove(QuickAttack{})
	// Output:
	// Pikachu used Thunderbolt!
	// Eevee used Quick Attack!
}

func ExampleThunderbolt_Use() {
	t := Thunderbolt{}
	t.Use()
	// Output: Pikachu used Thunderbolt!
}
