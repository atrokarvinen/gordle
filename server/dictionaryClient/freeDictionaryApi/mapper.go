package freeDictionaryApi

import "gordle/dictionaryClient/models"

func mapWordDetails(details Response) models.DictionaryDetails {
	word := details.Word
	allDetails := models.DictionaryDetails{
		Word:        word,
		Definitions: mapDefinitions(details),
		Examples:    mapExamples(details),
	}
	return allDetails
}

func mapDefinitions(details Response) []string {
	result := []string{}
	for _, entry := range details.Entries {
		for _, sense := range entry.Senses {
			result = append(result, sense.Definition)
		}
	}
	return result
}

func mapExamples(details Response) []string {
	result := []string{}
	for _, entry := range details.Entries {
		for _, sense := range entry.Senses {
			if len(sense.Examples) == 0 {
				continue
			}
			result = append(result, sense.Examples...)
		}
	}
	return result
}
