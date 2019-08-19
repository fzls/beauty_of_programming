package main

import . "github.com/fzls/beauty_of_programming"

func main() {
	c := &CakeSorting{}
	c.Init([]int{3,2,1,6,5,4,9,8,7,0})
	c.Run()
}
