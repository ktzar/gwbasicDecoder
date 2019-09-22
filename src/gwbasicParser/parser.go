package gwbasicParser

import "fmt"
import b "binaryutils"
import "errors"

type Program []Line

func (p Program) String() string {
    code := ""
    for _, line := range(p) {
        code += line.code
    }
    return code
}

func (p Program) WithLines() string {
    code := ""
    for _, line := range(p) {
        code += fmt.Sprintf("%v\t%v", line.number, line.code);
    }
    return code
}

func ParseProgram(data []byte) (Program, error) {
	if int(data[0]) != 0xff {
		return []Line{}, errors.New("Wrong header")
	}

	pointer := 1
	finished := false
	var program Program

	for finished == false {
		nextLineOffset := b.BE16(data, pointer)
		//fmt.Println("Next line offset", nextLineOffset)
		if nextLineOffset == 0x00 {
			finished = true
		} else {
			pointer += 2
			line, newPointer := decodeLine(data, pointer)
			program = append(program, line)
			pointer = newPointer
		}
	}

	return program, nil
}

