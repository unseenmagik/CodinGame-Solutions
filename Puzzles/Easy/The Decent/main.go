package main

import "fmt"

/**
 * The while loop represents the game.
 * Each iteration represents a turn of the game
 * where you are given inputs (the heights of the mountains)
 * and where you have to print an output (the index of the mountain to fire on)
 * The inputs you are given are automatically updated according to your last actions.
 **/

func main() {
	for {
		var max = -1
		var indexMax int
		for i := 0; i < 8; i++ {
			var height int // represents the height of one mountain.
			fmt.Scan(&height)
			if height > max {
				max = height
				indexMax = i
			}
		}

		fmt.Println(indexMax) // The index of the mountain to fire on.
	}
}
