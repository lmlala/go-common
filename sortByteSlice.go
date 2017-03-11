// Date : 11/3/17 PM6:07
// Copyright: TradeShift.com
// Author = 'liming'
// Email = 'lim@cn.tradeshift.com'
package common

import (
	"sort"
	"bytes"
	"log"
)

// implement `Interface` in sort package.
// must implement
//type Interface interface {
//	// Len is the number of elements in the collection.
//	Len() int
//	// Less reports whether the element with
//	// index i should sort before the element with index j.
//	Less(i, j int) bool
//	// Swap swaps the elements with indexes i and j.
//	Swap(i, j int)
//}

type byteslice [][]byte

func (b byteslice) Len() int {
	return len(b)
}

func (b byteslice) Less(i, j int) bool {
	switch bytes.Compare(b[i], b[j]) {
	case -1:
		return true
	case 0, 1:
		return false
	default:
		log.Panic("bytes.compare error: ", b[i], ", ", b[j])
		return false
	}
}

func (b byteslice) Swap(i, j int) {
	b[j], b[i] = b[i], b[j]
}

func SortByteSlice(src [][]byte) [][]byte {
	sorted := byteslice(src)
	sort.Sort(sorted)
	return sorted
}