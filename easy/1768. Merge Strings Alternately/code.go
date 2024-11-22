package main

func main() {

}

func mergeAlternately(word1 string, word2 string) string {
	var mergedString = ""
	lengthWord1 := len(word1)
	lengthWord2 := len(word2)

	biggestLen := lengthWord1
	if lengthWord2 > biggestLen {
		biggestLen = lengthWord2
	}

	i := 0
	for i < biggestLen {

		if lengthWord1 > i {
			mergedString += string(word1[i])
		}

		if lengthWord2 > i {
			mergedString += string(word2[i])
		}

		i++
	}

	return mergedString
}
