package summarize

import (
	"encoding/json"
	"path/filepath"
	"testing"

	"github.com/jdkato/prose/internal/util"
	"github.com/jdkato/syllables"
	"github.com/stretchr/testify/assert"
)

func TestSyllables(t *testing.T) {
	cases := util.ReadDataFile(filepath.Join(testdata, "syllables.json"))
	tests := make(map[string]int)
	util.CheckError(json.Unmarshal(cases, &tests))

	for word, count := range tests {
		assert.Equal(t, count, Syllables(word), word)
	}
}

func BenchmarkSyllables(b *testing.B) {
	cases := util.ReadDataFile(filepath.Join(testdata, "syllables.json"))
	tests := make(map[string]int)
	util.CheckError(json.Unmarshal(cases, &tests))

	for n := 0; n < b.N; n++ {
		for word := range tests {
			Syllables(word)
		}
	}
}

func BenchmarkSyllablesIn(b *testing.B) {
	cases := util.ReadDataFile(filepath.Join(testdata, "syllables.json"))
	tests := make(map[string]int)
	util.CheckError(json.Unmarshal(cases, &tests))

	for n := 0; n < b.N; n++ {
		for word := range tests {
			syllables.In(word)
		}
	}
}

type testCase struct {
	Text       string
	Sentences  float64
	Words      float64
	PolyWords  float64
	Characters float64
}

func TestSummarize(t *testing.T) {
	tests := make([]testCase, 0)
	cases := util.ReadDataFile(filepath.Join(testdata, "summarize.json"))

	util.CheckError(json.Unmarshal(cases, &tests))
	for _, test := range tests {
		d := NewDocument(test.Text)
		assert.Equal(t, test.Sentences, d.NumSentences)
		assert.Equal(t, test.Words, d.NumWords)
		assert.Equal(t, test.Characters, d.NumCharacters)
	}
}
