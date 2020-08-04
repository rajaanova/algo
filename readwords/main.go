package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func main() {
	//Below needs to be taken from file and put into list
	dictionaryWords := []string{"axpaj", "apxaj", "dnrbt", "pjxdn", "abd"}        //getListOfDictionary()
	stringsList := []string{"aapxjdnrbtvldptfzbbdbbzxtndrvjblnzjfpvhdhhpxjdnrbt"} //getListStrings()
	var count int
	for j, stringsLine := range stringsList {
		count = 0
		for _, value := range dictionaryWords {
			// i am ignoring error here
			isMatching, _ := regexp.MatchString(value, stringsLine)
			if isMatching {
				count++
				continue
			}

			dictWordStartIndex := value[0]
			dictWordLenIndex := len(value) - 1
			dictWordEndIndex := value[dictWordLenIndex]
			dictionarysorted := strings.Split(value, "")
			sort.Strings(dictionarysorted)
			for i := 0; i < len(stringsLine)-dictWordLenIndex; i++ {
				if stringsLine[i] == dictWordStartIndex && stringsLine[i+dictWordLenIndex] == dictWordEndIndex {
					lenOfString := strings.Split(stringsLine[i:i+dictWordLenIndex+1], "")
					sort.Strings(lenOfString)
					if strings.Join(lenOfString, "") == strings.Join(dictionarysorted, "") {
						count++
						break
					}

				}

			}
		}
		fmt.Println("case: ", j+1, count)
	}

}
