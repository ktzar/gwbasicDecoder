package gwbasicParser

import "fmt"
import "errors"
import b "binaryutils"

func ParseProgram(data []byte) (string, error) {
	fmt.Println(b.BE16(data, 0))
	if int(data[0]) == 0xff {
		fmt.Println("all ok")
	} else {
		return "", errors.New("Wrong header")
	}

	pointer := 1
	finished := false
	programBuffer := ""

	for finished == false {
		nextLineOffset := b.BE16(data, pointer)
		fmt.Println("Next line offset", nextLineOffset)
		if nextLineOffset == 0x00 {
			finished = true
		} else {
			pointer += 2
			line, newPointer := decodeLine(data, pointer)
			programBuffer += line
			pointer = newPointer
		}
	}

	return programBuffer, nil
}

func decodeLine(data []byte, pointer int) (string, int) {
	//fmt.Printf("Decode line at pointer 0x%x \n", pointer)
	lineNumber := b.LE16(data, pointer)
	//fmt.Println("Line number ", lineNumber)
	lineBuffer := fmt.Sprintf("%v\t", lineNumber)
	pointer += 2
	char := data[pointer]
	for char != 0x00 {
		var token string
		if char == 0x1c {
			/* Two byte constant */
			token = fmt.Sprintf("%v", b.LE16(data, pointer+1))
			pointer += 2
		} else if char == 0x0f {
			/* One byte constant */
			token = fmt.Sprintf("%v", (data[pointer+1]))
			pointer += 1
		} else if char >= 0x20 && char <= 0x7e {
			token = string(char)
		} else if oneByte[char] != "" {
			token = oneByte[char]
		} else if twoBytes[b.LE16(data, pointer)] != "" {
			token = twoBytes[b.LE16(data, pointer)]
			pointer += 1
		} else {
			fmt.Printf("Unrecognised token %x\n", char)
		}
		lineBuffer += token
		pointer++
		char = data[pointer]
	}
	pointer++
	lineBuffer += "\n"
	return lineBuffer, pointer
}
