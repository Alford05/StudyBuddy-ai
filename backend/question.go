package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
)

func buildSingleWordQuestion(wordBank []string, index int, apiKey string, rng *rand.Rand) (Question, error) {
	if index >= len(wordBank) {
		index = index % len(wordBank)
	}

	correctWord := wordBank[index]
	sentence, err := generateSentence(correctWord, apiKey)
	if err != nil {
		return Question{Sentence: "Error generating sentence."}, err
	}

	blankSentence := replaceWord(sentence, correctWord, "_______")
	options := buildMultipleChoiceOptions(correctWord, wordBank, rng)

	correctIndex := -1
	for i, opt := range options {
		if opt == correctWord {
			correctIndex = i
			break
		}
	}

	return Question{
		Sentence:     blankSentence,
		Options:      options,
		CorrectIndex: correctIndex,
		AnswerWords:  []string{correctWord},
	}, nil
}

func buildTwoWordQuestion(wordBank []string, apiKey string, rng *rand.Rand) (Question, error) {
	i1 := rng.Intn(len(wordBank))
	i2 := rng.Intn(len(wordBank))
	for i2 == i1 {
		i2 = rng.Intn(len(wordBank))
	}

	w1 := wordBank[i1]
	w2 := wordBank[i2]

	sentence, err := generateSentenceTwoWords(w1, w2, apiKey)
	if err != nil {
		return Question{Sentence: "Error generating sentence."}, err
	}

	const placeholder1 = "<<WORD1>>"
	const placeholder2 = "<<WORD2>>"

	sentenceWithBlanks := replaceWord(sentence, w1, placeholder1)
	sentenceWithBlanks = replaceWord(sentenceWithBlanks, w2, placeholder2)

	sentenceWithBlanks = strings.Replace(sentenceWithBlanks, placeholder1, "_______(1)", 1)
	sentenceWithBlanks = strings.Replace(sentenceWithBlanks, placeholder2, "_______(2)", 1)

	index1 := strings.Index(sentenceWithBlanks, "_______(1)")
	index2 := strings.Index(sentenceWithBlanks, "_______(2)")

	if index1 > index2 {
		sentenceWithBlanks = strings.Replace(sentenceWithBlanks, "_______(2)", "<<TEMP>>", 1)
		sentenceWithBlanks = strings.Replace(sentenceWithBlanks, "_______(1)", "_______(2)", 1)
		sentenceWithBlanks = strings.Replace(sentenceWithBlanks, "<<TEMP>>", "_______(1)", 1)

		w1, w2 = w2, w1
	}
	correctPair := fmt.Sprintf("%s, %s", w1, w2)

	options := buildTwoWordOptions(w1, w2, wordBank, rng)

	correctIndex := -1
	for i, opt := range options {
		if normalize(opt) == normalize(correctPair) {
			correctIndex = i
			break
		}
	}

	return Question{
		Sentence:     sentenceWithBlanks,
		Options:      options,
		CorrectIndex: correctIndex,
		AnswerWords:  []string{w1, w2},
	}, nil
}

func generateSentence(word, apiKey string) (string, error) {
	prompt := fmt.Sprintf("Write an English sentence based on a 7th grade reading level that uses the word \"%s\". Do not quote the word. Output only the sentence.", word)
	return generateSentenceFromPrompt(prompt, apiKey)
}

func generateSentenceTwoWords(w1, w2, apiKey string) (string, error) {
	prompt := fmt.Sprintf(
		"Write an English sentence at a 7th-grade reading level that uses BOTH words \"%s\" and \"%s\". "+
			"Ensure both words appear naturally in the same sentence. Output only the sentence.",
		w1, w2,
	)
	return generateSentenceFromPrompt(prompt, apiKey)
}

func replaceWord(sentence, word, placeholder string) string {
	re := regexp.MustCompile(`(?i)\b` + regexp.QuoteMeta(word) + `\b`)
	return re.ReplaceAllString(sentence, placeholder)
}

func buildTwoWordOptions(w1, w2 string, bank []string, rng *rand.Rand) []string {
	correctPair := fmt.Sprintf("%s, %s", w1, w2)
	options := make([]string, 0, 4)
	used := map[string]bool{correctPair: true}

	correctIndex := rng.Intn(4)
	for i := 0; i < 4; i++ {
		if i == correctIndex {
			options = append(options, correctPair)
		} else {
			var p1, p2 string
			for {
				p1 = bank[rng.Intn(len(bank))]
				p2 = bank[rng.Intn(len(bank))]
				pair := fmt.Sprintf("%s, %s", p1, p2)
				if p1 != p2 && !used[pair] {
					options = append(options, pair)
					used[pair] = true
					break
				}
			}
		}
	}
	return options
}

func buildMultipleChoiceOptions(correct string, bank []string, rng *rand.Rand) []string {
	options := make([]string, 0, 4)
	used := map[string]bool{correct: true}

	correctIndex := rng.Intn(4)
	for i := 0; i < 4; i++ {
		if i == correctIndex {
			options = append(options, correct)
		} else {
			var distractor string
			for {
				distractor = bank[rand.Intn(len(bank))]
				if !used[distractor] {
					break
				}
			}
			options = append(options, distractor)
			used[distractor] = true
		}
	}
	return options
}

func normalize(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}
