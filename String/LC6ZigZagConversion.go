package main

func convert(s string, numRows int) string {
	// edge cases
	if len(s) == 0 {
		return ""
	}
	if numRows == 1 {
		return s
	}

	lines := make([]rune, len(s))
	skip := numRows - 2 // 观察得出 列之间的skip大小为numRows-2

	tmp := 0 // 输出rune的游标
	for row := 0; row < numRows; row++ { // 一行行打印
		ptr := row // 观察得出规律 每新一行第一个字符在原字符串的下标为row
		for ptr < len(s) {
			lines[tmp] = rune(s[ptr])
			tmp++

			ptr += numRows + skip // 找规律:ptr在原字符串的下一个位置在ptr+numsRows+skip(每次跳过numRows+skip个位置)

			// 中间行 可能有zigzag diagonal
			if row > 0 && row < numRows - 1 {
				diagIdx := ptr - row * 2 // 找规律:每个diag元素下标等于:ptr新位置-当前row*2
				if diagIdx < len(s) {
					lines[tmp] = rune(s[diagIdx])
					tmp++
				}
			}
		}
	}

	return string(lines)
}