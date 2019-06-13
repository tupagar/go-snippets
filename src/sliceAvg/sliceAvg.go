// for a given random sequence of numbers, find the slice (consecutive set of numbers in a sequence) 
// which when removed from the sequence, gives maximum average of remaining numbers.
// Solution: find the moving window slice for example: a,b,c[d,e],f,g....  which when removed from the 
// actual sequence gives the higher average of remaining seq

package main

import (
	"fmt"
)

func sum(seq []int) int{
	var total = 0
	for _, value:= range seq {
		total += value
	}
	return total
}

func main() {
	seq := []int{15,4,7,22,3,8,9,66,99,1,3,6}
	seqSize := len(seq)
	seqSum := sum(seq)
	avg := float64(seqSum) / float64(seqSize)
	fmt.Printf("Beginning avg = %g \n", avg)
	maxAvg := avg

	for sliceSize := 1; sliceSize < seqSize; sliceSize++ {
		for itr :=0; itr < (seqSize-sliceSize); itr++ {
			slice := seq[itr:(itr+sliceSize)]
			newAvg := float64(seqSum - sum(slice)) / float64(seqSize - len(slice))
			if newAvg > maxAvg {
				maxAvg = newAvg
				fmt.Printf("Found new bigger avg %g for slice: %v \n", newAvg, slice)
			}
		}
	}
}
