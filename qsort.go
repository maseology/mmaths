package maths

import (
	"math/rand"
	"sort"
)

// QsortInterface : interface to sort.Sort
type QsortInterface interface {
	sort.Interface
	// Partition returns slice[:i] and slice[i+1:]
	// These should references the original memory
	// since this does an in-place sort
	Partition(i int) (left QsortInterface, right QsortInterface)
}

// IntSlice : alias to index array being sorted
type IntSlice []int

func (is IntSlice) Less(i, j int) bool {
	return is[i] < is[j]
}

func (is IntSlice) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

func (is IntSlice) Len() int {
	return len(is)
}

// Partition : slits index array around pivot
func (is IntSlice) Partition(i int) (left QsortInterface, right QsortInterface) {
	return IntSlice(is[:i]), IntSlice(is[i+1:])
}

// Qsort : a (quick) sorting algorithm that is apparently faster then Go's native sort.Sort
// from: https://stackoverflow.com/questions/23276417/golang-custom-sort-is-faster-than-native-sort#23278451
func Qsort(a QsortInterface, prng *rand.Rand) QsortInterface {
	if a.Len() < 2 {
		return a
	}

	left, right := 0, a.Len()-1

	// Pick a pivot
	pivotIndex := prng.Int() % a.Len()
	// Move the pivot to the right
	a.Swap(pivotIndex, right)

	// Pile elements smaller than the pivot on the left
	for i := 0; i < a.Len(); i++ {
		if a.Less(i, right) {
			a.Swap(i, left)
			left++
		}
	}

	// Place the pivot after the last smaller element
	a.Swap(left, right)

	// Go down the rabbit hole
	leftSide, rightSide := a.Partition(left)
	Qsort(leftSide, prng)
	Qsort(rightSide, prng)

	return a
}

// QsortIndx : same as above, but preserves original slice index
// modified from: https://stackoverflow.com/questions/23276417/golang-custom-sort-is-faster-than-native-sort#23278451
func QsortIndx(a QsortIndxInterface, prng *rand.Rand) []int {
	if a.Len() < 2 {
		return []int{0}
	}

	left, right := 0, a.Len()-1

	// Pick a pivot
	pivotIndex := prng.Int() % a.Len()
	// Move the pivot to the right
	a.Swap(pivotIndex, right)

	// Pile elements smaller than the pivot on the left
	for i := 0; i < a.Len(); i++ {
		if a.Less(i, right) {
			a.Swap(i, left)
			left++
		}
	}

	// Place the pivot after the last smaller element
	a.Swap(left, right)

	// Go down the rabbit hole
	leftSide, rightSide := a.Partition(left)
	QsortIndx(leftSide, prng)
	QsortIndx(rightSide, prng)

	return a.Indices()
}

// QsortIndxInterface : interface to sort.Sort
type QsortIndxInterface interface {
	sort.Interface
	// Partition returns slice[:i] and slice[i+1:]
	// These should references the original memory
	// since this does an in-place sort
	Partition(i int) (left QsortIndxInterface, right QsortIndxInterface)
	Indices() []int
}

// IndexedSlice : alias to float array being sorted
type IndexedSlice struct {
	indx []int
	val  []float64
}

// Indices : returns the index property
func (is IndexedSlice) Indices() []int {
	return is.indx
}

func (is IndexedSlice) Len() int {
	return len(is.indx)
}

func (is IndexedSlice) Less(i, j int) bool {
	return is.val[i] < is.val[j]
}

func (is IndexedSlice) Swap(i, j int) {
	is.indx[i], is.indx[j] = is.indx[j], is.indx[i]
	is.val[i], is.val[j] = is.val[j], is.val[i]
}

// Partition : splits index array around pivot
func (is IndexedSlice) Partition(i int) (left QsortIndxInterface, right QsortIndxInterface) {
	left = IndexedSlice{
		indx: is.indx[:i],
		val:  is.val[:i],
	}
	right = IndexedSlice{
		indx: is.indx[i+1:],
		val:  is.val[i+1:],
	}
	return
}
