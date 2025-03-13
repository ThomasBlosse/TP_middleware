package helpers

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

func ParseICalEvents(rawData []byte) ([]map[string]string, error) {
	scanner := bufio.NewScanner(bytes.NewReader(rawData))

	var eventArray []map[string]string
	currentEvent := map[string]string{}
	currentKey := ""
	currentValue := ""

	inEvent := false

	for scanner.Scan() {
		line := scanner.Text()

		if !inEvent && line != "BEGIN:VEVENT" {
			continue
		}

		if line == "BEGIN:VEVENT" {
			inEvent = true
			currentEvent = map[string]string{}
			continue
		}

		// If end of event, store and reset
		if line == "END:VEVENT" {
			inEvent = false
			eventArray = append(eventArray, currentEvent)
			continue
		}

		// If multi-line data, append to the current key
		if strings.HasPrefix(line, " ") {
			currentEvent[currentKey] += strings.TrimSpace(line)
			continue
		}

		// Split scan
		fmt.Println(scanner.Text())
		splitted := strings.SplitN(scanner.Text(), ":", 2)
		currentKey = splitted[0]
		currentValue = splitted[1]

		// Store current event attribute
		currentEvent[currentKey] = currentValue
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading event data: %w", err)
	}

	return eventArray, nil
}
