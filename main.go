package main

import (
	"Bg2Json/types"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	bible := make(map[string]map[int]map[int]types.Verse)

	/// Loop on the './bible' folder
	books, err := os.ReadDir("./bible")
	if err != nil {
		panic(err)
	}

	for _, book := range books {
		// Loop on the './bible/{book}' folder
		bible[book.Name()] = make(map[int]map[int]types.Verse)

		chapters, _ := os.ReadDir("./bible/" + book.Name())

		for _, chapter := range chapters {
			// Loop on the './bible/{book}/{chapter}' folder
			chapterNumber, err := getChapterNumber(chapter.Name())
			if err != nil { // Skip the file if it's not a chapter
				continue
			}
			bible[book.Name()][chapterNumber] = make(map[int]types.Verse)

			text, err := os.ReadFile("./bible/" + book.Name() + "/" + chapter.Name())
			if err != nil {
				panic(err)
			}

			lines := bytes.Split(text, []byte("\n"))

			isVerse := false
			verseCount := 1
			for i, line := range lines {
				if strings.HasPrefix(string(line), "######") {
					isVerse = true
					continue
				}
				if isVerse {
					// Parse the verse
					// Add the verse to the Bible struct
					verse := types.Verse{
						Text: string(line),
						Uri:  book.Name() + "/" + chapter.Name(),
						Line: i,
					}
					bible[book.Name()][chapterNumber][verseCount] = verse
					verseCount++
					isVerse = false
				}
			}

		}
	}

	// Marshal the Bible struct to JSON
	// Write the JSON to a file
	b, err := json.Marshal(bible)
	if err != nil {
		fmt.Println(err)
		return
	}
	//export to bible.json
	err = os.WriteFile("bible.json", b, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// Extract the number of the chapter from the chapter name
// Ex 1 roi 16 => 16
// Ex Jean 3 => 3
func getChapterNumber(chapter string) (int, error) {
	//remove .md
	chapter = strings.Replace(chapter, ".md", "", -1)
	spaces := strings.Split(chapter, " ")
	return strconv.Atoi(spaces[len(spaces)-1])
}
