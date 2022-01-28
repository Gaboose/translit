package translit

import (
	"bufio"
	"io"
	"os"
	"sort"
	"strings"
)

type Translator struct {
	srcWordIndex *WordIndex
	tuples       [][2]string
	srcToTrg     map[string]string
}

func (t *Translator) ReadFile(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)

	t.tuples = t.tuples[:0]

	for {
		bts, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}

		var tpl StringTuple

		if err := tpl.UnmarshalJSON(bts); err != nil {
			return err
		}

		t.tuples = append(t.tuples, tpl)
	}

	sort.Slice(t.tuples, func(i, j int) bool {
		return t.tuples[i][0] < t.tuples[j][0]
	})

	sourceWords := make([]string, 0, len(t.tuples))
	for _, tpl := range t.tuples {
		sourceWords = append(sourceWords, tpl[0])
	}

	t.srcWordIndex = NewWordIndex(sourceWords)

	t.srcToTrg = map[string]string{}

	for _, tpl := range t.tuples {
		t.srcToTrg[tpl[0]] = tpl[1]
	}

	return nil
}

// Match returns a tuple that best matches the given string.
func (t *Translator) Match(s string) [2]string {
	res := t.srcWordIndex.Match(s)

	if len(res) == 0 {
		return [2]string{}
	}

	return [2]string{res[0], t.srcToTrg[res[0]]}
}

func (t *Translator) Suggestions(s string) []string {
	ret := t.srcWordIndex.Match(s)
	return ret
}

// Translate returns multiple ordered tuples that best match the given string.
func (t *Translator) Translate(s string) [][2]string {
	var ret = [][2]string{}
	for len(s) > 0 {
		s = strings.TrimLeft(s, " \n")
		ii := t.srcWordIndex.idx.LongestMatch([]byte(s))

		// if no matches found, remove the first char
		if len(ii) == 0 {
			s = s[1:]
			continue
		}

		w := t.srcWordIndex.word(ii[0])
		n := len(w)

		if n < len(s) {
			s = s[n:]
		} else {
			s = ""
		}

		ret = append(ret, [2]string{w, t.srcToTrg[w]})
	}

	return ret
}
