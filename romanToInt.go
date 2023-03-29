package main

func romanToInt(s string) int {
	romanMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	result := 0
	for i := range s {
		if i < len(s)-1 && romanMap[s[i:i+1]] < romanMap[s[i+1:i+2]] {
			result -= romanMap[s[i:i+1]]
		} else {
			result += romanMap[s[i:i+1]]
		}
	}

	return result
}
