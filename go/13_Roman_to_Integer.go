func romanToInt(s string) int {
	Integer := 0

	for i := 0; i < len(s); i++ {
		switch s[i:(i + 1)] {
		case "I":
			Integer += 1
		case "V":
			Integer += 5
			if i > 0 && s[(i-1):i] == "I" {
				Integer -= 2
			}
		case "X":
			Integer += 10
			if i > 0 && s[(i-1):i] == "I" {
				Integer -= 2
			}
		case "L":
			Integer += 50
			if i > 0 && s[(i-1):i] == "X" {
				Integer -= 20
			}
		case "C":
			Integer += 100
			if i > 0 && s[(i-1):i] == "X" {
				Integer -= 20
			}
		case "D":
			Integer += 500
			if i > 0 && s[(i-1):i] == "C" {
				Integer -= 200
			}
		case "M":
			Integer += 1000
			if i > 0 && s[(i-1):i] == "C" {
				Integer -= 200
			}
		}
	}
	return Integer
}