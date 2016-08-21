package id

// __provinceLevelAddressCodeTable 保存了 省级行政单位 区划代码(区划代码的前两位数字)的数据.
var __provinceLevelAddressCodeTable = [100]bool{
	false, false, false, false, false, false, false, false, false, false,
	false, true, true, true, true, true, false, false, false, false,
	false, true, true, true, false, false, false, false, false, false,
	false, true, true, true, true, true, true, true, false, false,
	false, true, true, true, true, true, true, false, false, false,
	true, true, true, true, true, false, false, false, false, false,
	false, true, true, true, true, true, false, false, false, false,
	false, true, false, false, false, false, false, false, false, false,
	false, true, true, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false, false, false,
}

// isAddressCodeValid 返回 true 如果 code 是有效的, 否则返回 false.
func isAddressCodeValid(code string) bool {
	// TODO 目前只检查前两位, 以后检查整个6位?
	index := int(code[0]-'0')*10 + int(code[1]-'0')
	if index < 0 || index >= 100 {
		return false
	}
	return __provinceLevelAddressCodeTable[index]
}
