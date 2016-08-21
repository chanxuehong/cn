package id

var __weightArray = []byte{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}

// Checksum 计算并返回 中华人民共和国公民身份号码 的校验码.
func Checksum(number string) byte {
	var sum int
	for i := 0; i < 17; i++ {
		sum += int(number[i]-'0') * int(__weightArray[i])
	}
	r := (12 - (sum % 11)) % 11
	if r == 10 {
		return 'X'
	}
	return byte(r) + '0'
}
