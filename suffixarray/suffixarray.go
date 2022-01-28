package suffixarray

import (
	"bytes"
	"sort"
)

type Index struct {
	data []byte
	sa   []int
}

func New(data []byte) *Index {
	ret := &Index{
		data: data,
		sa:   make([]int, 0, len(data)),
	}

	for i := 0; i < len(ret.data); i++ {
		ret.sa = append(ret.sa, i)
	}

	sort.Slice(ret.sa, func(i, j int) bool {
		return bytes.Compare(ret.at(i), ret.at(j)) < 0
	})

	return ret
}

func (x *Index) at(i int) []byte {
	return x.data[x.sa[i]:]
}

func (x *Index) Lookup(s []byte) []int {
	i := sort.Search(len(x.sa), func(i int) bool { return bytes.Compare(x.at(i), s) >= 0 })
	j := i + sort.Search(len(x.sa)-i, func(j int) bool { return !bytes.HasPrefix(x.at(j+i), s) })
	return x.sa[i:j]
}

func (x *Index) Data() []byte {
	return x.data
}

func CommonPrefix(left, right []byte) int {
	var i int

	for i = range left {
		if left[i] != right[i] {
			break
		}
	}

	return i
}

func (x *Index) LongestMatch(s []byte) []int {
	i := sort.Search(len(x.sa), func(i int) bool { return bytes.Compare(x.at(i), s) >= 0 })

	// If s is small it might match exactly with some part of the text.
	// In that case, return those occurances.
	if bytes.HasPrefix(x.at(i), s) {
		j := i + sort.Search(len(x.sa)-i, func(j int) bool { return !bytes.HasPrefix(x.at(j+i), s) })
		return x.sa[i:j]
	}

	// If s didn't match exactly, the suffix with the
	// longest match will be previous to i.
	n := CommonPrefix(s, x.at(i-1))
	if n == 0 {
		return nil
	}

	var j int

	// Inlude other longest prefixes
	for j = i - 2; j >= 0; j-- {
		if CommonPrefix(s, x.at(j)) < n {
			break
		}
	}

	return x.sa[j+1 : i]
}
