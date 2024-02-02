package wordsApi

import "gordle/dictionaryClient/models"

func mapWordDetails(details WordDetails) models.DictionaryDetails {
	return models.DictionaryDetails{
		Word:        details.Word,
		Definitions: mapDefinitions(details.Results),
		Examples:    mapExamples(details.Results),
	}
}

func mapDefinitions(definitions []WordResults) []string {
	result := []string{}
	for _, definition := range definitions {
		result = append(result, definition.Definition)
	}
	return result
}

func mapExamples(examples []WordResults) []string {
	result := []string{}
	for _, example := range examples {
		for _, example := range example.Examples {
			result = append(result, example)
		}
	}
	return result
}
