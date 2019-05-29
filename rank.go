package main

import (
	"fmt"
	"sort"
)

func main() {
	// Declare a slice of type float64 for times to failure
	// need to input the data later - using test data to code the math now

	ttf := []float64{7000, 1500, 4300, 2250, 4000}
	ttfs := []float64{1750, 5000}
	ttfc := ttf

	// print ttf before sorting
	fmt.Println("Original failure input ", ttf)
	fmt.Println("Original suspended input ", ttfs)
	//combine the two slices into 1
	l1 := int(len(ttfs))
	for i1 := 0; i1 < l1; i1++ {
		ttfc = append(ttfc, ttfs[i1])
	}

	fmt.Println("Combined inputs unsorted ", ttfc)

	//sort in order
	sort.Float64s(ttfc)
	sort.Float64s(ttf) //needs to be sorted for Bernard ranking

	//print a blank line then results
	fmt.Println()
	fmt.Println("Sorted input ", ttfc)

	//Rank, Reverse Rank
	rankOrder := []int{}
	revRankOrder := []int{}
	ttfcLen := int(len(ttfc))
	for i2 := 0; i2 < ttfcLen; i2++ {
		rankOrder = append(rankOrder, i2+1)
		revRankOrder = append(revRankOrder, ttfcLen-i2)
	}
	fmt.Println("Rank Order ", rankOrder)
	fmt.Println("Reverse Rank Order ", revRankOrder)

	//Calculate i's for Bernard's Median Ranking
	// start with 3 float64 arrays  ttf (exists and is sorted), ranks of ttf, reverse ranks of ttf
	// output will be two arrays with values of ranks and reverse ranks associate with failures (in order)
	ttfLen := len(ttf)
	ttfrnk := []float64{}
	ttfrr := []float64{}
	for i := 0; i < ttfcLen; i++ {
		for j := 0; j < ttfLen; j++ {
			if ttf[j] == ttfc[i] {
				ttfrnk = append(ttfrnk, float64(rankOrder[i]))
				ttfrr = append(ttfrr, float64(revRankOrder[i]))
			}
		}
	}
	fmt.Println()

	fmt.Println("Times to Failures ", ttf)
	fmt.Println("Ranking of ttf    ", ttfrnk)
	fmt.Println("Reverse rankning  ", ttfrr)
	fmt.Println()
	//calculate adjusted ranks (RR x Prior Adj Rank) + (N+1)/(RR + 1)
	adjrnkN := float64(ttfcLen)
	rank1 := (adjrnkN + 1) / (ttfrr[0] + 1)
	adjRank := []float64{rank1}
	fmt.Println("rank1, AdjRank ", rank1, adjRank)

	for i := 1; i < ttfLen; i++ { //start indexing at 1 since first value already calculated
		z := i - 1
		w := float64(ttfcLen)
		q := float64(ttfrr[i])
		adjr := ((q * adjRank[z]) + (w + 1)) / (q + 1)
		adjRank = append(adjRank, adjr)
	}
	fmt.Println("Adjusted Ranks for Bernard Calcs ", adjRank)

	// Bernard's Median Rank (i - 0.3)/(N + 0.4) i = adjusted ranks and N is
	// declare the variable
	bmr := []float64{} //float64
	bf := float64(ttfcLen)
	for i := 0; i < ttfLen; i++ {
		bmr = append(bmr, ((adjRank[i] - 0.3) / (bf + 0.4)))
		fmt.Println(bmr[i])
	}

	fmt.Println()
	fmt.Println("The ranks are ", bmr)

}
