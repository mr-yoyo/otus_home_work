package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

const regPattern = `([^\pL]\p{Pd}[^\pL])|[^\pL\p{Pd}]`

var regExp = regexp.MustCompile(regPattern)

type topWord struct {
	Word      string
	Frequency int
}

func Top10(text string) (top []string) {
	words := fetchWordsLowered(text)
	freqs := countFrequencies(words)
	topWords := makeTopWords(freqs)
	sortTopWords(topWords)

	length := len(topWords)
	if length > 10 {
		length = 10
	}

	for i := 0; i < length; i++ {
		top = append(top, topWords[i].Word)
	}

	return top
}

func fetchWordsLowered(text string) []string {
	text = strings.ToLower(regExp.ReplaceAllString(text, " "))

	return strings.Fields(text)
}

func countFrequencies(words []string) map[string]int {
	freqs := make(map[string]int)

	for _, word := range words {
		freqs[word]++
	}

	return freqs
}

func makeTopWords(freqs map[string]int) []topWord {
	topWords := make([]topWord, len(freqs))

	i := 0
	for word, frequency := range freqs {
		topWords[i] = topWord{
			Word:      word,
			Frequency: frequency,
		}
		i++
	}

	return topWords
}

func sortTopWords(topWords []topWord) {
	sort.Slice(topWords, func(i, j int) bool {
		if topWords[i].Frequency == topWords[j].Frequency {
			return topWords[i].Word < topWords[j].Word
		}

		return topWords[i].Frequency > topWords[j].Frequency
	})
}
