var roman = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
    if s == "" {
        return 0
    }

    num, lastint, total := 0, 0, 0
    for _, r := range reverse(s) {
        num = roman[string(r)]
        if num < lastint {
            total = total - num
        } else {
            total = total + num
        }
        lastint = num
    }
    return total
}

func reverse(s string) string {
    n := len(s)
    runes := make([]rune, n)
    for _, rune := range s {
        n--
        runes[n] = rune
    }
    return string(runes)
}