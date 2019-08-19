package beauty_of_programming

import "fmt"

type CakeSorting struct {
	Cakes []int

	ReverseOps []int

	MinStepToSorted int
	MinStepToSortedReverseOps []int

	SearchCount int
}

func (c *CakeSorting) Init(cakes[] int){
	c.Cakes = append(c.Cakes, cakes...)

	c.ReverseOps = nil

	c.MinStepToSorted = c.upperBound()
	c.MinStepToSortedReverseOps = nil
}

func (c *CakeSorting) isSorted() bool {
	for i := 1; i < len(c.Cakes); i++ {
		if c.Cakes[i] < c.Cakes[i-1]{
			return false
		}
	}

	return true
}

func (c *CakeSorting) upperBound() int{
	return 2*len(c.Cakes)
}

func (c *CakeSorting) lowerBound() int{
	minStep := 0
	for i:=1; i<len(c.Cakes); i++{
		diff := c.Cakes[i] - c.Cakes[i-1]
		if !( diff== 1 || diff == -1){
			minStep++
		}
	}

	return minStep
}
// reverse cakes[0 ... idx] => cakes[idx ... 0]
func (c *CakeSorting) _reverse(j int){
	i:=0
	for i<j {
		temp := c.Cakes[i]
		c.Cakes[i] = c.Cakes[j]
		c.Cakes[j] = temp

		i++
		j--
	}
}

func (c *CakeSorting) reverse(j int){
	c._reverse(j)
	c.ReverseOps = append(c.ReverseOps, j)
}

func (c *CakeSorting) unreverse(j int){
	c._reverse(j)
	c.ReverseOps = c.ReverseOps[:len(c.ReverseOps)-1]
}

func (c *CakeSorting) search(step int) {
	c.SearchCount++

	optimisticStepToSorted := c.lowerBound()
	if step + optimisticStepToSorted >= c.MinStepToSorted {
		return
	}

	if c.isSorted() && step < c.MinStepToSorted {
		c.MinStepToSorted = step
		c.MinStepToSortedReverseOps = append([]int{}, c.ReverseOps...)
		fmt.Printf("Debug:step=%v, ops=%v\n", step, c.ReverseOps)
		return
	}

	for i:=1;i<len(c.Cakes);i++{
		c.reverse(i)
		c.search(step+1)
		c.unreverse(i)
	}
}

func (c *CakeSorting) Run(){
	c.search(0)
	fmt.Printf("Debug: res: step=%v, searchCount=%v, ops=%v\n", c.MinStepToSorted, c.SearchCount, c.MinStepToSortedReverseOps)
}