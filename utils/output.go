package utils

import (
	"bytes"
	"fmt"
	"strings"
)

func ParseCmdOutput(output []byte, expectedCountOfFields int) ([][]string, error) {
	lines := bytes.FieldsFunc(output, func(c rune) bool { return c == '\n' || c == '\r' })

	var recordLines [][]string
	for i, line := range lines {
		recordLine := splitBySeparator(":", string(line))
		if len(recordLine) != expectedCountOfFields {
			return nil, fmt.Errorf(
				"line %d contains %d fields but should %d",
				i, len(recordLine), expectedCountOfFields,
			)
		}

		recordLines = append(recordLines, recordLine)
	}

	return recordLines, nil
}

func splitBySeparator(separator, line string) []string {
	escape := `/`
	tempEscapedSeparator := "\x00"

	replacedEscape := strings.ReplaceAll(line, escape+separator, tempEscapedSeparator)
	records := strings.Split(replacedEscape, separator)

	for i, record := range records {
		records[i] = strings.ReplaceAll(record, tempEscapedSeparator, separator)
	}

	return records
}
