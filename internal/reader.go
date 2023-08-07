package internal

import (
	"encoding/json"
	"os"
)

type Sample struct {
	Sample struct {
		Nest string `json:"nest"`
	} `json:"sample"`
}

func CreateCSV() ([]*Follower, error) {
	files, err := os.ReadDir("./followers-json")
	if err != nil {
		return nil, err
	}

	outputFile, err := os.Create("./outputs/output.csv")
	if err != nil {
		return nil, err
	}
	defer outputFile.Close()

	for _, file := range files {
		var rawData RawFollowersJson
		byteValue, _ := os.ReadFile("./followers-json/" + file.Name())
		json.Unmarshal(byteValue, &rawData)

		for _, instruction := range rawData.Data.User.Result.Timeline.Timeline.Instructions {
			if instruction.Type != "TimelineAddEntries" {
				continue
			}

			for _, entry := range instruction.Entries {
				if entry.Content.EntryType != "TimelineTimelineItem" {
					continue
				}

				data := entry.Content.ItemContent.UserResults.Result.Legacy

				s := "@" + data.ScreenName + "," + data.Name + ",https://twitter.com/" + data.ScreenName + "\n"
				outputFile.WriteString(s)
			}
		}
	}

	return nil, nil
}
