package piscine

func Atoi(s string) int {
    if len(s) == 0 { // 1️⃣ Empty string → 0
        return 0
    }

    num := 0
    sign := 1
    start := 0

    // 2️⃣ Handle sign
    if s[0] == '-' || s[0] == '+' {
        if len(s) > 1 && (s[1] == '-' || s[1] == '+') {
            // 2a. Multiple signs at start → invalid
            return 0
        }
        if s[0] == '-' {
            sign = -1
        }
        start = 1
    }

    // 3️⃣ Parse digits
    for i := start; i < len(s); i++ {
        r := s[i]
        if r >= '0' && r <= '9' {
            digit := int(r - '0')  // Convert rune to int
            num = num*10 + digit   // Build number
        } else {
            // 3a. First non-digit → stop parsing
            if num == 0 {
                return 0 // If no digits seen yet → invalid string
            }
            break
        }
    }

    return num * sign // 4️⃣ Apply sign
}
