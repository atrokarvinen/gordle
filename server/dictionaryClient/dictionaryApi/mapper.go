package dictionaryApi

import "gordle/dictionaryClient/models"

func mapWordDetails(details []Response) models.DictionaryDetails {
	if len(details) == 0 {
		return models.DictionaryDetails{}
	}
	word := details[0].Word
	allDetails := models.DictionaryDetails{
		Word:        word,
		Definitions: mapDefinitions(details),
		Examples:    mapExamples(details),
	}
	return allDetails
}

func mapDefinitions(details []Response) []string {
	result := []string{}
	for _, detail := range details {
		for _, meaning := range detail.Meanings {
			for _, definition := range meaning.Definitions {
				result = append(result, definition.Definition)
			}
		}
	}
	return result
}

func mapExamples(details []Response) []string {
	result := []string{}
	for _, detail := range details {
		for _, meaning := range detail.Meanings {
			for _, definition := range meaning.Definitions {
				if definition.Example == "" {
					continue
				}
				result = append(result, definition.Example)
			}
		}
	}
	return result
}
