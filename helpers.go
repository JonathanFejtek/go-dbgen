package dbgen

import (
	"reflect"
)

func filterTags(tags []string, tagsToOmit []string) []string {
	var tagsFiltered []string

	for _, tag := range tags {
		isFiltered := false
		for _, omitTag := range tagsToOmit {
			if tag == omitTag {
				isFiltered = true
				break
			}
		}

		if !isFiltered {
			tagsFiltered = append(tagsFiltered, tag)
		}
	}

	return tagsFiltered

}

func getTagsByName(tagName string, i interface{}) []string {
	t := reflect.TypeOf(i)
	var tagValues []string

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get(tagName)
		if tag != "" {
			tagValues = append(tagValues, tag)
		}
	}
	return tagValues
}
