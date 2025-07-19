# Arrays in Go
Arrays in Go are fixed-size, ordered sequences of elements of the same type. They're a fundamental data structure, though often less common in everyday Go programming than slices due to their fixed size.

## Declaring Arrays
In Go, you can declare arrays in several ways:

```go
// Declare an array of 5 integers (zero-initialized)
var arr [5]int

// Declare and initialize with values
var arr2 = [5]int{1, 2, 3, 4, 5}

// Short declaration syntax
arr3 := [5]int{1, 2, 3, 4, 5}
```

## Go Arrays: What Can Go Wrong
Here are the key pitfalls to watch out for when working with arrays in Go:

### Fixed Size 
Arrays have a fixed size defined at compile time. If you need a dynamic size, consider using slices instead.

```go
package main
import "fmt"

// Function that accepts an array of exactly 3 Pokémon
func printPokemons(pokemons [2]string){
    for i, pokemon := range pokemons {
        fmt.Printf("Pokemon %d: %s\n",i+1,pokemon)
    }
}

func main() {
    // This works as the array size matches
    pokemons := [2]string{"Pikachu", "Bulbasaur"}
    printPokemons(pokemons)
    
    // This will cause a compile-time error
    morePokemons := [3]string{"Pikachu", "Bulbasaur", "Charmander"}
    printPokemons(morePokemons)
}
```
