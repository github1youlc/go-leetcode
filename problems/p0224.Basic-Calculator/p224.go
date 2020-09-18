package p224

import "bytes"

func calculate(s string) int {
	var stack []string

	num := &bytes.Buffer{}
	for _, b := range s {
		if b == '+' || b == '-' || b == '(' {
			if num.Len() != 0 {
				stack = append(stack, num.String())
				num = &bytes.Buffer{}
			}
			stack = append(stack, string(b))
		} else if b == ')' {

		} else {
			num.WriteByte(byte(b))
		}
	}

	if num.Len() != 0 {
		stack = append(stack, num.String())
	}



}
