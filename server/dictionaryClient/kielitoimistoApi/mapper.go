package kielitoimistoApi

import (
	"fmt"
	"gordle/dictionaryClient/models"
)

func mapWordDetails(details Response) models.DictionaryDetails {
	return models.DictionaryDetails{
		Word:        details.Word,
		Definitions: mapDefinitions(details.Senses),
		Examples:    mapExamples(details.Senses),
	}
}

func mapDefinitions(sense []Sense) []string {
	result := []string{}
	for _, sense := range sense {
		for _, definition := range sense.Definitions {
			sanitized := sanitizeDescription(definition.Definition)
			fmt.Println("definition", definition, "=> sanitized:", sanitized)
			if len(sense.Senses) == 0 {
				result = append(result, sanitized)
				continue
			}
			prefix := sanitized
			for _, innerSense := range sense.Senses {
				for _, innerDefinition := range innerSense.Definitions {
					sanitized := sanitizeDescription(innerDefinition.Definition)
					fmt.Println("inner definition", innerDefinition, "=> sanitized:", sanitized)
					sanitized = prefix + " " + sanitized
					result = append(result, sanitized)
				}
			}
		}
	}
	return result
}

// Words may contain html like <span class="kt-pos">adv. </span>, filter out the html
func sanitizeDescription(description string) string {
	var sanitized []rune
	openingTagBegin := false
	openingTagEnd := false
	closingTagStart := false
	closingTagEnd := false
	for _, r := range description {
		if r == '<' {
			closingTagStart = openingTagBegin
			openingTagBegin = true
		} else if r == '>' {
			closingTagEnd = openingTagEnd
			openingTagEnd = true
		} else if openingTagEnd && !closingTagStart {
			sanitized = append(sanitized, r)
		} else if !openingTagBegin && !openingTagEnd {
			sanitized = append(sanitized, r)
		} else if closingTagStart && closingTagEnd {
			sanitized = append(sanitized, r)
		}
		if openingTagBegin && openingTagEnd && closingTagStart && closingTagEnd {
			openingTagBegin = false
			openingTagEnd = false
			closingTagStart = false
			closingTagEnd = false
		}
	}
	return string(sanitized)
}

func mapExamples(senses []Sense) []string {
	result := []string{}
	for _, sense := range senses {
		for _, example := range sense.Examples {
			sanitized := sanitizeDescription(example.Text)
			result = append(result, sanitized)
		}
		for _, innerSense := range sense.Senses {
			for _, innerExample := range innerSense.Examples {
				sanitized := sanitizeDescription(innerExample.Text)
				result = append(result, sanitized)
			}
		}
	}
	return result
}
