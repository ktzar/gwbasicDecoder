package gwbasicParser;

import "fmt"
import b "binaryutils"

type Line struct {
    number int
    code string
}

func decodeLine(data []byte, pointer int) (Line, int) {
	lineNumber := b.LE16(data, pointer)
	lineBuffer := ""
	pointer += 2
	char := data[pointer]
	for char != 0x00 {
		var token string
		if char == 0x1c {
			/* Two byte constant */
			token = fmt.Sprintf("%v", b.LE16(data, pointer+1))
			pointer += 2
		} else if char == 0x1d {
			/* four byte floating-point */
			token = fmt.Sprintf("%v", b.LE32(data, pointer+1))
			pointer += 4
		} else if char == 0x0f {
			/* One byte constant */
			token = fmt.Sprintf("%v", (data[pointer+1]))
			pointer += 1
		} else if char == 0x0e {
			/* Two byte line number */
			token = fmt.Sprintf("%v", b.LE16(data, pointer+1))
			pointer += 2
		} else if char >= 0x11 && char <= 0x1b {
			// 0x11 to 0x1b are numbers from 0 to 10
			token = fmt.Sprintf("%v", (char - 0x11))
		} else if char >= 0x20 && char <= 0x7e {
			token = string(char)
		} else if oneByte[char] != "" {
			token = oneByte[char]
		} else if char >= 0xfd {
			token = twoBytes[b.BE16(data, pointer)]
			pointer += 1
		} else {
			fmt.Printf("Unrecognised token 0x%x\n", char)
		}
		lineBuffer += token
		pointer++
		char = data[pointer]
	}
	pointer++
	lineBuffer += "\n"
	return Line{lineNumber, lineBuffer}, pointer
}
