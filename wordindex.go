package translit

import (
	"sort"

	"github.com/Gaboose/translit/suffixarray"
)

const sep byte = 0

type WordIndex struct {
	idx *suffixarray.Index
}

func NewWordIndex(words []string) *WordIndex {
	data := []byte{sep}
	for _, s := range words {
		data = append(data, []byte(s)...)
		data = append(data, sep)
	}

	return &WordIndex{
		idx: suffixarray.New([]byte(data)),
	}
}

func (sa *WordIndex) word(c int) string {
	var start, end int

	data := sa.idx.Data()

	for i := c - 1; i >= 0; i-- {
		if data[i] == sep {
			start = i
			break
		}
	}

	for i := c + 1; i < len(data); i++ {
		if data[i] == sep {
			end = i
			break
		}
	}

	return string(data[start+1 : end])
}

func (sa *WordIndex) Match(s string) []string {
	res := sa.idx.Lookup([]byte(s))

	set := map[string]struct{}{}
	ret := make([]string, 0, len(res))

	for _, i := range res {
		w := sa.word(i)

		if _, ok := set[w]; ok {
			continue
		}

		set[w] = struct{}{}
		ret = append(ret, w)
	}

	sort.Slice(ret, func(i, j int) bool {
		if len(ret[i]) != len(ret[j]) {
			return len(ret[i]) < len(ret[j])
		}

		return ret[i] < ret[j]
	})

	return ret
}
