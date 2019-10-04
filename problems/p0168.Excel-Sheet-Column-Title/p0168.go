package p0168_Excel_Sheet_Column_Title


func convertToTitle(n int) string {
	var result string
	var r int
	for n != 0 {
		if n % 26 == 0 {
			return convertToTitle(n / 26 - 1) + "Z" + result
		}

		r = n % 26
		n = n / 26

		result =  convert(r) + result
	}

	return result
}

func convert(d int) string {
	return string('A' + d - 1)
}