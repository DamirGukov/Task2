package main

import "fmt"

// Task1
type Game struct {
	TotalPlayers int
	Field1       string
	Field2       int
	Field3       float64
	Field4       bool
}

// Task2
type TableGame struct {
	Game
	Field5 string
}

func main() {
	//Task3
	object1 := TableGame{
		Game: Game{
			TotalPlayers: 11,
			Field1:       "Field1",
			Field2:       3,
			Field3:       33.654,
			Field4:       false,
		},
		Field5: "Field5",
	}

	//Task4, Task5
	slice := make([]TableGame, 0, 3)
	slice = append(slice, object1)
	slice = append(slice, TableGame{
		Game: Game{
			TotalPlayers: 8,
			Field1:       "NewField1",
			Field2:       6,
			Field3:       15.954,
			Field4:       true,
		},
		Field5: "NewField5",
	})
	slice = append(slice, TableGame{
		Game: Game{
			TotalPlayers: 5,
			Field1:       "DifferentField1",
			Field2:       16,
			Field3:       101.899,
			Field4:       false,
		},
		Field5: "DifferentField5",
	})
	fmt.Println(slice)

	//Task7
	newMap := make(map[int]TableGame)

	for _, data := range slice {
		newMap[data.TotalPlayers] = data
	}

	//Task8
	for _, value := range newMap {
		fmt.Println(value)
	}

	//Task9, Task10
	total := 0
	for _, value := range newMap {
		total += value.TotalPlayers
	}

	fmt.Println("TotalPlayers:", total)
}
